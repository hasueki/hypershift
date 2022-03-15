package olm

import (
	"strings"

	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
)

type OperatorLifecycleManagerParams struct {
	CertifiedOperatorsImage string
	CommunityOperatorsImage string
	RedHatMarketplaceImage  string
	RedHatOperatorsImage    string
	OLMCatalogPlacement     string
}

func NewOperatorLifecycleManagerParams(hcp *hyperv1.HostedControlPlane, releaseVersion string) *OperatorLifecycleManagerParams {
	tag := strings.Join(strings.Split(releaseVersion, ".")[:2], ".")

	params := &OperatorLifecycleManagerParams{
		CertifiedOperatorsImage: "registry.redhat.io/redhat/certified-operator-index:v" + tag,
		CommunityOperatorsImage: "registry.redhat.io/redhat/community-operator-index:v" + tag,
		RedHatMarketplaceImage:  "registry.redhat.io/redhat/redhat-marketplace-index:v" + tag,
		RedHatOperatorsImage:    "registry.redhat.io/redhat/redhat-operator-index:v" + tag,
		OLMCatalogPlacement:     hcp.Spec.OLMCatalogPlacement,
	}

	return params
}
