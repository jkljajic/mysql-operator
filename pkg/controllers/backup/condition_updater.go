package backup

import (
	klog "k8s.io/klog/v2"

	backuputil "github.com/jkljajic/mysql-operator/pkg/api/backup"
	"github.com/jkljajic/mysql-operator/pkg/apis/mysql/v1alpha1"
	clientset "github.com/jkljajic/mysql-operator/pkg/generated/clientset/versioned/typed/mysql/v1alpha1"
)

// ConditionUpdater enables updating Backup conditions.
type ConditionUpdater interface {
	Update(backup *v1alpha1.Backup, condition *v1alpha1.BackupCondition) error
}

type conditionUpdater struct {
	client clientset.BackupsGetter
}

func (p *conditionUpdater) Update(backup *v1alpha1.Backup, condition *v1alpha1.BackupCondition) error {
	klog.Infof("Updating Backup condition for %s/%s to (%s==%s)", backup.Namespace, backup.Name, condition.Type, condition.Status)
	if backuputil.UpdateBackupCondition(&backup.Status, condition) {
		_, err := p.client.Backups(backup.Namespace).Update(backup)
		return err
	}
	return nil
}
