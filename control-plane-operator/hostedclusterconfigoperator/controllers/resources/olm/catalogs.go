package olm

import (
	"fmt"
	"math/big"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/hypershift/support/assets"
	"github.com/openshift/hypershift/support/config"

	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
)

var (
	certifiedCatalogService         = assets.MustService("assets/catalog-certified.service.yaml")
	communityCatalogService         = assets.MustService("assets/catalog-community.service.yaml")
	redHatMarketplaceCatalogService = assets.MustService("assets/catalog-redhat-marketplace.service.yaml")
	redHatOperatorsCatalogService   = assets.MustService("assets/catalog-redhat-operators.service.yaml")
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

var (
	certifiedCatalogDeployment         = assets.MustDeployment("assets/catalog-certified.deployment.yaml")
	communityCatalogDeployment         = assets.MustDeployment("assets/catalog-community.deployment.yaml")
	redHatMarketplaceCatalogDeployment = assets.MustDeployment("assets/catalog-redhat-marketplace.deployment.yaml")
	redHatOperatorsCatalogDeployment   = assets.MustDeployment("assets/catalog-redhat-operators.deployment.yaml")
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

var (
	certifiedCatalogRolloutCronJob         = assets.MustCronJob("assets/catalog-certified-rollout.cronjob.yaml")
	communityCatalogRolloutCronJob         = assets.MustCronJob("assets/catalog-community-rollout.cronjob.yaml")
	redHatMarketplaceCatalogRolloutCronJob = assets.MustCronJob("assets/catalog-redhat-marketplace-rollout.cronjob.yaml")
	redHatOperatorsCatalogRolloutCronJob   = assets.MustCronJob("assets/catalog-redhat-operators-rollout.cronjob.yaml")
)

func ReconcileCertifiedOperatorsCronJob(cronJob *batchv1beta1.CronJob, ownerRef config.OwnerRef, cliImage string) error {
	return reconcileCatalogCronJob(cronJob, ownerRef, cliImage, certifiedCatalogRolloutCronJob)
}
func ReconcileCommunityOperatorsCronJob(cronJob *batchv1beta1.CronJob, ownerRef config.OwnerRef, cliImage string) error {
	return reconcileCatalogCronJob(cronJob, ownerRef, cliImage, communityCatalogRolloutCronJob)
}
func ReconcileRedHatMarketplaceOperatorsCronJob(cronJob *batchv1beta1.CronJob, ownerRef config.OwnerRef, cliImage string) error {
	return reconcileCatalogCronJob(cronJob, ownerRef, cliImage, redHatMarketplaceCatalogRolloutCronJob)
}
func ReconcileRedHatOperatorsCronJob(cronJob *batchv1beta1.CronJob, ownerRef config.OwnerRef, cliImage string) error {
	return reconcileCatalogCronJob(cronJob, ownerRef, cliImage, redHatOperatorsCatalogRolloutCronJob)
}

func reconcileCatalogCronJob(cronJob *batchv1beta1.CronJob, ownerRef config.OwnerRef, cliImage string, sourceCronJob *batchv1beta1.CronJob) error {
	ownerRef.ApplyTo(cronJob)
	cronJob.Spec = sourceCronJob.DeepCopy().Spec
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image = cliImage
	cronJob.Spec.Schedule = generateModularDailyCronSchedule([]byte(cronJob.Namespace))
	return nil
}

// generateModularDailyCronSchedule returns a daily crontab schedule
// where, given a is input's integer representation, the minute is a % 60
// and hour is a % 24.
func generateModularDailyCronSchedule(input []byte) string {
	a := big.NewInt(0).SetBytes(input)
	var hi, mi big.Int
	m := mi.Mod(a, big.NewInt(60))
	h := hi.Mod(a, big.NewInt(24))
	return fmt.Sprintf("%d %d * * *", m.Int64(), h.Int64())
}

var (
	catalogRolloutRole        = assets.MustRole("assets/catalog-rollout.role.yaml")
	catalogRolloutRoleBinding = assets.MustRoleBinding("assets/catalog-rollout.rolebinding.yaml")
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

func ReconcileCertifiedOperatorsCatalogSource(cs *operatorsv1alpha1.CatalogSource) {
	reconcileCatalogSource(cs, "certified-operators:50051", "Certified Operators", -200)
}

func ReconcileCommunityOperatorsCatalogSource(cs *operatorsv1alpha1.CatalogSource) {
	reconcileCatalogSource(cs, "community-operators:50051", "Community Operators", -400)
}

func ReconcileRedHatMarketplaceCatalogSource(cs *operatorsv1alpha1.CatalogSource) {
	reconcileCatalogSource(cs, "redhat-marketplace:50051", "Red Hat Marketplace", -300)
}

func ReconcileRedHatOperatorsCatalogSource(cs *operatorsv1alpha1.CatalogSource) {
	reconcileCatalogSource(cs, "redhat-operators:50051", "Red Hat Operators", -100)
}

func reconcileCatalogSource(cs *operatorsv1alpha1.CatalogSource, address, displayName string, priority int) {
	if cs.Annotations == nil {
		cs.Annotations = map[string]string{}
	}
	cs.Annotations["target.workload.openshift.io/management"] = `{"effect": "PreferredDuringScheduling"}`
	cs.Spec = operatorsv1alpha1.CatalogSourceSpec{
		SourceType:  operatorsv1alpha1.SourceTypeGrpc,
		Address:     address,
		DisplayName: displayName,
		Publisher:   "Red Hat",
		Priority:    priority,
		UpdateStrategy: &operatorsv1alpha1.UpdateStrategy{
			RegistryPoll: &operatorsv1alpha1.RegistryPoll{
				Interval: &metav1.Duration{Duration: 10 * time.Minute},
			},
		},
	}
}
