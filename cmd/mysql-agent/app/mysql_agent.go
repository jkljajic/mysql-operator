package app

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/heptiolabs/healthcheck"
	"github.com/pkg/errors"

	kubeinformers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"

	options "github.com/oracle/mysql-operator/cmd/mysql-agent/app/options"
	cluster "github.com/oracle/mysql-operator/pkg/cluster"
	backupcontroller "github.com/oracle/mysql-operator/pkg/controllers/backup"
	clustermgr "github.com/oracle/mysql-operator/pkg/controllers/cluster/manager"
	restorecontroller "github.com/oracle/mysql-operator/pkg/controllers/restore"
	mysqlop "github.com/oracle/mysql-operator/pkg/generated/clientset/versioned"
	informers "github.com/oracle/mysql-operator/pkg/generated/informers/externalversions"
	signals "github.com/oracle/mysql-operator/pkg/util/signals"
)

// resyncPeriod computes the time interval a shared informer waits before
// resyncing with the api server.
func resyncPeriod(opts *options.MySQLAgentOpts) func() time.Duration {
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(opts.MinResyncPeriod.Nanoseconds()) * factor)
	}
}

// Run runs the MySQL backup controller. It should never exit.
func Run(opts *options.MySQLAgentOpts) error {
	kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	// Set up signals so we handle the first shutdown signal gracefully.
	signals.SetupSignalHandler(cancelFunc)

	// Set up healthchecks (liveness and readiness).
	health := healthcheck.NewHandler()
	health.AddReadinessCheck("node-in-cluster",
		healthcheck.Async(
			healthcheck.Timeout(func() error { return cluster.CheckNodeInCluster(ctx) }, 5*time.Second),
			10*time.Second,
		))
	go func() {
		glog.Fatal(http.ListenAndServe(
			net.JoinHostPort(opts.Address, strconv.Itoa(int(opts.HealthcheckPort))),
			health,
		))
	}()

	kubeclient := kubernetes.NewForConfigOrDie(kubeconfig)
	mysqlopClient := mysqlop.NewForConfigOrDie(kubeconfig)

	sharedInformerFactory := informers.NewFilteredSharedInformerFactory(mysqlopClient, 0, opts.Namespace, nil)
	kubeInformerFactory := kubeinformers.NewFilteredSharedInformerFactory(kubeclient, resyncPeriod(opts)(), opts.Namespace, nil)

	var wg sync.WaitGroup

	manager, err := clustermgr.NewLocalClusterManger(kubeclient, kubeInformerFactory)
	if err != nil {
		return errors.Wrap(err, "failed to create new local MySQL InnoDB cluster manager")
	}
	// Block until local instance successfully initialised.
	for !manager.Sync(ctx) {
		time.Sleep(10 * time.Second)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		manager.Run(ctx)
	}()

	backupController := backupcontroller.NewAgentController(
		kubeclient,
		mysqlopClient.MysqlV1(),
		sharedInformerFactory.Mysql().V1().MySQLBackups(),
		sharedInformerFactory.Mysql().V1().MySQLClusters(),
		kubeInformerFactory.Core().V1().Pods(),
		opts.Hostname,
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		backupController.Run(ctx, 5)
	}()

	restoreController := restorecontroller.NewAgentController(
		kubeclient,
		mysqlopClient.MysqlV1(),
		sharedInformerFactory.Mysql().V1().MySQLRestores(),
		sharedInformerFactory.Mysql().V1().MySQLClusters(),
		sharedInformerFactory.Mysql().V1().MySQLBackups(),
		kubeInformerFactory.Core().V1().Pods(),
		opts.Hostname,
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		restoreController.Run(ctx, 5)
	}()

	// Shared informers have to be started after ALL controllers.
	go sharedInformerFactory.Start(ctx.Done())
	go kubeInformerFactory.Start(ctx.Done())

	<-ctx.Done()

	glog.Info("Waiting for all controllers to shut down gracefully")
	wg.Wait()

	return nil
}