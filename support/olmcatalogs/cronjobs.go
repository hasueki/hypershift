package olmcatalogs

import (
	// TODO: Switch to k8s.io/api/batch/v1 when all management clusters at 1.21+ OR 4.8_openshift+
	batchv1beta1 "k8s.io/api/batch/v1beta1"

	"github.com/openshift/hypershift/support/assets"
	"github.com/openshift/hypershift/support/config"
	"github.com/openshift/hypershift/support/util"
)

type CronJobReconciler struct {
	Manifest  func(string) *batchv1beta1.CronJob
	Reconcile func(*batchv1beta1.CronJob, config.OwnerRef, string) error
}

func GetCronJobReconcilers() []CronJobReconciler {
	return []CronJobReconciler{
		{Manifest: CertifiedOperatorsCronJob, Reconcile: ReconcileCertifiedOperatorsCronJob},
		{Manifest: CommunityOperatorsCronJob, Reconcile: ReconcileCommunityOperatorsCronJob},
		{Manifest: RedHatMarketplaceOperatorsCronJob, Reconcile: ReconcileRedHatMarketplaceOperatorsCronJob},
		{Manifest: RedHatOperatorsCronJob, Reconcile: ReconcileRedHatOperatorsCronJob},
	}
}

var (
	certifiedCatalogRolloutCronJob         = assets.MustCronJob("olm/catalog-certified-rollout.cronjob.yaml")
	communityCatalogRolloutCronJob         = assets.MustCronJob("olm/catalog-community-rollout.cronjob.yaml")
	redHatMarketplaceCatalogRolloutCronJob = assets.MustCronJob("olm/catalog-redhat-marketplace-rollout.cronjob.yaml")
	redHatOperatorsCatalogRolloutCronJob   = assets.MustCronJob("olm/catalog-redhat-operators-rollout.cronjob.yaml")
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
	cronJob.Spec.Schedule = util.GenerateModularDailyCronSchedule([]byte(cronJob.Namespace))
	return nil
}
