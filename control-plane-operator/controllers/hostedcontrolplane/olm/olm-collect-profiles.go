package olm

import (
	"github.com/openshift/hypershift/support/assets"
	"github.com/openshift/hypershift/support/config"
	"github.com/openshift/hypershift/support/util"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

var (
	olmCollectProfilesConfigMap   = assets.MustConfigMap("olm/olm-collect-profiles.configmap.yaml")
	olmCollectProfilesCronJob     = assets.MustCronJob("olm/olm-collect-profiles.cronjob.yaml")
	olmCollectProfilesRole        = assets.MustRole("olm/olm-collect-profiles.role.yaml")
	olmCollectProfilesRoleBinding = assets.MustRoleBinding("olm/olm-collect-profiles.rolebinding.yaml")
	olmCollectProfilesSecret      = assets.MustSecret("olm/olm-collect-profiles.secret.yaml")
)

func ReconcileCollectProfilesCronJob(cronJob *batchv1beta1.CronJob, ownerRef config.OwnerRef, olmImage, namespace string) {
	ownerRef.ApplyTo(cronJob)
	cronJob.Spec = olmCollectProfilesCronJob.DeepCopy().Spec
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image = olmImage
	for i, arg := range cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Args {
		if arg == "OLM_NAMESPACE" {
			cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Args[i] = namespace
		}
	}
	cronJob.Spec.Schedule = util.GenerateModularDailyCronSchedule([]byte(cronJob.Namespace))
}

func ReconcileCollectProfilesConfigMap(configMap *corev1.ConfigMap, ownerRef config.OwnerRef) {
	ownerRef.ApplyTo(configMap)
	configMap.Data = olmCollectProfilesConfigMap.DeepCopy().Data
}

func ReconcileCollectProfilesRole(role *rbacv1.Role, ownerRef config.OwnerRef) {
	ownerRef.ApplyTo(role)
	role.Rules = olmCollectProfilesRole.DeepCopy().Rules
}

func ReconcileCollectProfilesRoleBinding(roleBinding *rbacv1.RoleBinding, ownerRef config.OwnerRef) {
	ownerRef.ApplyTo(roleBinding)
	roleBinding.RoleRef = olmCollectProfilesRoleBinding.DeepCopy().RoleRef
	roleBinding.Subjects = olmCollectProfilesRoleBinding.DeepCopy().Subjects
}

func ReconcileCollectProfilesSecret(secret *corev1.Secret, ownerRef config.OwnerRef) {
	ownerRef.ApplyTo(secret)
	secret.Type = olmCollectProfilesSecret.Type
	secret.Data = olmCollectProfilesSecret.DeepCopy().Data
}

func ReconcileCollectProfilesServiceAccount(serviceAccount *corev1.ServiceAccount, ownerRef config.OwnerRef) {
	ownerRef.ApplyTo(serviceAccount)
}
