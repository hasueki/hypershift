package olm

import (
	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	"github.com/openshift/hypershift/support/config"
)

type OperatorLifecycleManagerParams struct {
	CLIImage         string
	DeploymentConfig config.DeploymentConfig
	config.OwnerRef
}

func NewOperatorLifecycleManagerParams(hcp *hyperv1.HostedControlPlane, images map[string]string) *OperatorLifecycleManagerParams {
	params := &OperatorLifecycleManagerParams{
		CLIImage: images["cli"],
		OwnerRef: config.OwnerRefFrom(hcp),
	}
	params.DeploymentConfig = config.DeploymentConfig{
		Replicas: 1,
		Scheduling: config.Scheduling{
			PriorityClass: config.DefaultPriorityClass,
		},
	}
	params.DeploymentConfig.SetColocation(hcp)
	params.DeploymentConfig.SetRestartAnnotation(hcp.ObjectMeta)
	params.DeploymentConfig.SetReleaseImageAnnotation(hcp.Spec.ReleaseImage)
	params.DeploymentConfig.SetControlPlaneIsolation(hcp)

	params.DeploymentConfig.SetDefaultSecurityContext = false

	return params
}
