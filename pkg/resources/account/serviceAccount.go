package account

import (
	"github.com/jkljajic/mysql-operator/pkg/apis/mysql/v1alpha1"
	"github.com/jkljajic/mysql-operator/pkg/constants"
	"k8s.io/apimachinery/pkg/runtime/schema"

	corev1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createServiceAccount(cluster *v1alpha1.Cluster) *rbac.ClusterRoleBinding {

	//ss, err := m.statefulSetLister.StatefulSets(cluster.Namespace).Get(cluster.Name)

	ac := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{constants.ClusterLabel: cluster.Name},
			Name:      cluster.Spec.ServiceAccountName,
			Namespace: cluster.ObjectMeta.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    v1alpha1.ClusterCRDResourceKind,
				}),
			},
		},
	}
	cr := &rbac.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{constants.ClusterLabel: cluster.Name},
			Name:      cluster.Spec.ServiceAccountName,
			Namespace: cluster.ObjectMeta.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    v1alpha1.ClusterCRDResourceKind,
				}),
			},
		},
		Rules: []rbac.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"pods"},
				Verbs:     []string{"get", "list", "patch", "update", "watch"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"secrets"},
				Verbs:     []string{"get"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"events"},
				Verbs:     []string{"create", "update", "patch"},
			},
			{
				APIGroups: []string{"mysql.oracle.com"},
				Resources: []string{"mysqlbackups", "mysqlbackupschedules", "mysqlclusters", "mysqlclusters/finalizers", "mysqlrestores"},
				Verbs:     []string{"get", "list", "patch", "update", "watch"},
			},
		},
	}

	bind := &rbac.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: rbac.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{constants.ClusterLabel: cluster.Name},
			Name:      cluster.Spec.ServiceAccountName,
			Namespace: cluster.ObjectMeta.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cluster, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    v1alpha1.ClusterCRDResourceKind,
				}),
			},
		},
		Subjects: []rbac.Subject{
			{
				Kind:      ac.Kind,
				APIGroup:  ac.APIVersion,
				Name:      ac.Name,
				Namespace: ac.Namespace,
			},
		},
		RoleRef: rbac.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     cr.Kind,
			Name:     cr.Name,
		},
	}

	return bind
}
