package olmcatalogs

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/openshift/hypershift/support/assets"
	"github.com/openshift/hypershift/support/config"
)

type ServiceReconciler struct {
	Manifest  func(string) *corev1.Service
	Reconcile func(*corev1.Service, config.OwnerRef) error
}

func GetServiceReconcilers() []ServiceReconciler {
	return []ServiceReconciler{
		{Manifest: CertifiedOperatorsService, Reconcile: ReconcileCertifiedOperatorsService},
		{Manifest: CommunityOperatorsService, Reconcile: ReconcileCommunityOperatorsService},
		{Manifest: RedHatMarketplaceOperatorsService, Reconcile: ReconcileRedHatMarketplaceOperatorsService},
		{Manifest: RedHatOperatorsService, Reconcile: ReconcileRedHatOperatorsService},
	}
}

var (
	certifiedCatalogService         = assets.MustService("olm/catalog-certified.service.yaml")
	communityCatalogService         = assets.MustService("olm/catalog-community.service.yaml")
	redHatMarketplaceCatalogService = assets.MustService("olm/catalog-redhat-marketplace.service.yaml")
	redHatOperatorsCatalogService   = assets.MustService("olm/catalog-redhat-operators.service.yaml")
)

func ReconcileCertifiedOperatorsService(svc *corev1.Service, ownerRef config.OwnerRef) error {
	return reconcileCatalogService(svc, ownerRef, certifiedCatalogService)
}

func ReconcileCommunityOperatorsService(svc *corev1.Service, ownerRef config.OwnerRef) error {
	return reconcileCatalogService(svc, ownerRef, communityCatalogService)
}

func ReconcileRedHatMarketplaceOperatorsService(svc *corev1.Service, ownerRef config.OwnerRef) error {
	return reconcileCatalogService(svc, ownerRef, redHatMarketplaceCatalogService)
}

func ReconcileRedHatOperatorsService(svc *corev1.Service, ownerRef config.OwnerRef) error {
	return reconcileCatalogService(svc, ownerRef, redHatOperatorsCatalogService)
}

func reconcileCatalogService(svc *corev1.Service, ownerRef config.OwnerRef, sourceService *corev1.Service) error {
	ownerRef.ApplyTo(svc)
	// The service is assigned a cluster IP when it is created.
	// This field is immutable as shown here: https://github.com/kubernetes/api/blob/62998e98c313b2ca15b1da278aa702bdd7b84cb0/core/v1/types.go#L4114-L4130
	// As such, to avoid an error when updating the object, only update the fields OLM specifies.
	sourceServiceDeepCopy := sourceService.DeepCopy()
	svc.Spec.Ports = sourceServiceDeepCopy.Spec.Ports
	svc.Spec.Type = sourceServiceDeepCopy.Spec.Type
	svc.Spec.Selector = sourceServiceDeepCopy.Spec.Selector

	return nil
}
