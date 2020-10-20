package restore

import (
	klog "k8s.io/klog/v2"

	restoreutil "github.com/jkljajic/mysql-operator/pkg/api/restore"
	"github.com/jkljajic/mysql-operator/pkg/apis/mysql/v1alpha1"
	clientset "github.com/jkljajic/mysql-operator/pkg/generated/clientset/versioned/typed/mysql/v1alpha1"
)

// ConditionUpdater enables updating Restore conditions.
type ConditionUpdater interface {
	Update(restore *v1alpha1.Restore, condition *v1alpha1.RestoreCondition) error
}

type conditionUpdater struct {
	client clientset.RestoresGetter
}

func (p *conditionUpdater) Update(restore *v1alpha1.Restore, condition *v1alpha1.RestoreCondition) error {
	klog.Infof("Updating Restore condition for %s/%s to (%s==%s)", restore.Namespace, restore.Name, condition.Type, condition.Status)
	if restoreutil.UpdateRestoreCondition(&restore.Status, condition) {
		_, err := p.client.Restores(restore.Namespace).Update(restore)
		return err
	}
	return nil
}
