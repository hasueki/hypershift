package olmcatalogs

import (
	appsv1 "k8s.io/api/apps/v1"

	"github.com/openshift/hypershift/support/assets"
	"github.com/openshift/hypershift/support/config"
)

type DeploymentReconciler struct {
	Manifest  func(string) *appsv1.Deployment
	Reconcile func(*appsv1.Deployment, config.OwnerRef, config.DeploymentConfig) error
}

func GetDeploymentReconcilers() []DeploymentReconciler {
	return []DeploymentReconciler{
		{Manifest: CertifiedOperatorsDeployment, Reconcile: ReconcileCertifiedOperatorsDeployment},
		{Manifest: CommunityOperatorsDeployment, Reconcile: ReconcileCommunityOperatorsDeployment},
		{Manifest: RedHatMarketplaceOperatorsDeployment, Reconcile: ReconcileRedHatMarketplaceOperatorsDeployment},
		{Manifest: RedHatOperatorsDeployment, Reconcile: ReconcileRedHatOperatorsDeployment},
	}
}

var (
	certifiedCatalogDeployment         = assets.MustDeployment("olm/catalog-certified.deployment.yaml")
	communityCatalogDeployment         = assets.MustDeployment("olm/catalog-community.deployment.yaml")
	redHatMarketplaceCatalogDeployment = assets.MustDeployment("olm/catalog-redhat-marketplace.deployment.yaml")
	redHatOperatorsCatalogDeployment   = assets.MustDeployment("olm/catalog-redhat-operators.deployment.yaml")
)

func ReconcileCertifiedOperatorsDeployment(deployment *appsv1.Deployment, ownerRef config.OwnerRef, dc config.DeploymentConfig) error {
	return reconcileCatalogDeployment(deployment, ownerRef, dc, certifiedCatalogDeployment)
}

func ReconcileCommunityOperatorsDeployment(deployment *appsv1.Deployment, ownerRef config.OwnerRef, dc config.DeploymentConfig) error {
	return reconcileCatalogDeployment(deployment, ownerRef, dc, communityCatalogDeployment)
}

func ReconcileRedHatMarketplaceOperatorsDeployment(deployment *appsv1.Deployment, ownerRef config.OwnerRef, dc config.DeploymentConfig) error {
	return reconcileCatalogDeployment(deployment, ownerRef, dc, redHatMarketplaceCatalogDeployment)
}

func ReconcileRedHatOperatorsDeployment(deployment *appsv1.Deployment, ownerRef config.OwnerRef, dc config.DeploymentConfig) error {
	return reconcileCatalogDeployment(deployment, ownerRef, dc, redHatOperatorsCatalogDeployment)
}

func reconcileCatalogDeployment(deployment *appsv1.Deployment, ownerRef config.OwnerRef, dc config.DeploymentConfig, sourceDeployment *appsv1.Deployment) error {
	ownerRef.ApplyTo(deployment)
	deployment.Spec = sourceDeployment.DeepCopy().Spec
	dc.ApplyTo(deployment)
	return nil
}
