package olmcatalogs

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/openshift/hypershift/support/assets"
	"github.com/openshift/hypershift/support/config"
)

type RBACReconcilers struct {
	ServiceAccount *ServiceAccountReconciler
	Role           *RoleReconciler
	RoleBinding    *RoleBindingReconciler
}

type ServiceAccountReconciler struct {
	Manifest  func(string) *corev1.ServiceAccount
	Reconcile func(*corev1.ServiceAccount, config.OwnerRef) error
}

type RoleReconciler struct {
	Manifest  func(string) *rbacv1.Role
	Reconcile func(*rbacv1.Role, config.OwnerRef) error
}

type RoleBindingReconciler struct {
	Manifest  func(string) *rbacv1.RoleBinding
	Reconcile func(*rbacv1.RoleBinding, config.OwnerRef) error
}

func GetRBACReconcilers() *RBACReconcilers {
	return &RBACReconcilers{
		ServiceAccount: &ServiceAccountReconciler{
			Manifest: CatalogRolloutServiceAccount, Reconcile: ReconcileCatalogRolloutServiceAccount,
		},
		Role: &RoleReconciler{
			Manifest: CatalogRolloutRole, Reconcile: ReconcileCatalogRolloutRole,
		},
		RoleBinding: &RoleBindingReconciler{
			Manifest: CatalogRolloutRoleBinding, Reconcile: ReconcileCatalogRolloutRoleBinding,
		},
	}
}

var (
	catalogRolloutRole        = assets.MustRole("olm/catalog-rollout.role.yaml")
	catalogRolloutRoleBinding = assets.MustRoleBinding("olm/catalog-rollout.rolebinding.yaml")
)

func ReconcileCatalogRolloutServiceAccount(sa *corev1.ServiceAccount, ownerRef config.OwnerRef) error {
	ownerRef.ApplyTo(sa)
	return nil
}

func ReconcileCatalogRolloutRole(role *rbacv1.Role, ownerRef config.OwnerRef) error {
	ownerRef.ApplyTo(role)
	role.Rules = catalogRolloutRole.DeepCopy().Rules
	return nil
}

func ReconcileCatalogRolloutRoleBinding(roleBinding *rbacv1.RoleBinding, ownerRef config.OwnerRef) error {
	ownerRef.ApplyTo(roleBinding)
	roleBinding.RoleRef = catalogRolloutRoleBinding.DeepCopy().RoleRef
	roleBinding.Subjects = catalogRolloutRoleBinding.DeepCopy().Subjects
	return nil
}
