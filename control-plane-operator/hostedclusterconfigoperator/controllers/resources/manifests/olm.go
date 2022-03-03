package manifests

import (
	appsv1 "k8s.io/api/apps/v1"
	//TODO: Switch to k8s.io/api/batch/v1 when all management clusters at 1.21+ OR 4.8_openshift+
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"

	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	prometheusoperatorv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

// Certified Operators Catalog

func CertifiedOperatorsDeployment(ns string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "certified-operators-catalog",
			Namespace: ns,
		},
	}
}

func CertifiedOperatorsService(ns string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "certified-operators",
			Namespace: ns,
		},
	}
}

func CertifiedOperatorsCronJob(ns string) *batchv1beta1.CronJob {
	return &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "certified-operators-catalog-rollout",
			Namespace: ns,
		},
	}
}

func CertifiedOperatorsCatalogSource() *operatorsv1alpha1.CatalogSource {
	return &operatorsv1alpha1.CatalogSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "certified-operators",
			Namespace: "openshift-marketplace",
		},
	}
}

// Community Operators Catalog

func CommunityOperatorsDeployment(ns string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "community-operators-catalog",
			Namespace: ns,
		},
	}
}

func CommunityOperatorsService(ns string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "community-operators",
			Namespace: ns,
		},
	}
}

func CommunityOperatorsCronJob(ns string) *batchv1beta1.CronJob {
	return &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "community-operators-catalog-rollout",
			Namespace: ns,
		},
	}
}

func CommunityOperatorsCatalogSource() *operatorsv1alpha1.CatalogSource {
	return &operatorsv1alpha1.CatalogSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "community-operators",
			Namespace: "openshift-marketplace",
		},
	}
}

// RedHatMarketplace Operators Catalog

func RedHatMarketplaceOperatorsDeployment(ns string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-marketplace-catalog",
			Namespace: ns,
		},
	}
}

func RedHatMarketplaceOperatorsService(ns string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-marketplace",
			Namespace: ns,
		},
	}
}

func RedHatMarketplaceOperatorsCronJob(ns string) *batchv1beta1.CronJob {
	return &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-marketplace-catalog-rollout",
			Namespace: ns,
		},
	}
}

func RedHatMarketplaceCatalogSource() *operatorsv1alpha1.CatalogSource {
	return &operatorsv1alpha1.CatalogSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-marketplace",
			Namespace: "openshift-marketplace",
		},
	}
}

// RedHat Operators Catalog

func RedHatOperatorsDeployment(ns string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-operators-catalog",
			Namespace: ns,
		},
	}
}

func RedHatOperatorsService(ns string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-operators",
			Namespace: ns,
		},
	}
}

func RedHatOperatorsCronJob(ns string) *batchv1beta1.CronJob {
	return &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-operators-catalog-rollout",
			Namespace: ns,
		},
	}
}

func RedHatOperatorsCatalogSource() *operatorsv1alpha1.CatalogSource {
	return &operatorsv1alpha1.CatalogSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "redhat-operators",
			Namespace: "openshift-marketplace",
		},
	}
}

// Catalog Rollout RBAC

func CatalogRolloutRole(ns string) *rbacv1.Role {
	return &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "catalog-rollout",
			Namespace: ns,
		},
	}
}

func CatalogRolloutRoleBinding(ns string) *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "catalog-rollout",
			Namespace: ns,
		},
	}
}

func CatalogRolloutServiceAccount(ns string) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "catalog-rollout",
			Namespace: ns,
		},
	}
}

// OLM packageserver services

func OLMPackageServerAPIService() *apiregistrationv1.APIService {
	return &apiregistrationv1.APIService{
		ObjectMeta: metav1.ObjectMeta{
			Name: "v1.packages.operators.coreos.com",
		},
	}
}

func OLMPackageServerService() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "packageserver",
			Namespace: "default",
		},
	}
}

func OLMPackageServerControlPlaneService(ns string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "packageserver",
			Namespace: ns,
		},
	}
}

func OLMPackageServerEndpoints() *corev1.Endpoints {
	return &corev1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "packageserver",
			Namespace: "default",
		},
	}
}

func OLMAlertRules() *prometheusoperatorv1.PrometheusRule {
	return &prometheusoperatorv1.PrometheusRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "olm-alert-rules",
			Namespace: "openshift-operator-lifecycle-manager",
		},
	}
}
