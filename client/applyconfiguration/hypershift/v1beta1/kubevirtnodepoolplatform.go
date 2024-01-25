/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	hypershiftv1beta1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
)

// KubevirtNodePoolPlatformApplyConfiguration represents an declarative configuration of the KubevirtNodePoolPlatform type for use
// with apply.
type KubevirtNodePoolPlatformApplyConfiguration struct {
	RootVolume                 *KubevirtRootVolumeApplyConfiguration `json:"rootVolume,omitempty"`
	Compute                    *KubevirtComputeApplyConfiguration    `json:"compute,omitempty"`
	NetworkInterfaceMultiQueue *hypershiftv1beta1.MultiQueueSetting  `json:"networkInterfaceMultiqueue,omitempty"`
	AdditionalNetworks         []KubevirtNetworkApplyConfiguration   `json:"additionalNetworks,omitempty"`
	AttachDefaultNetwork       *bool                                 `json:"attachDefaultNetwork,omitempty"`
}

// KubevirtNodePoolPlatformApplyConfiguration constructs an declarative configuration of the KubevirtNodePoolPlatform type for use with
// apply.
func KubevirtNodePoolPlatform() *KubevirtNodePoolPlatformApplyConfiguration {
	return &KubevirtNodePoolPlatformApplyConfiguration{}
}

// WithRootVolume sets the RootVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RootVolume field is set to the value of the last call.
func (b *KubevirtNodePoolPlatformApplyConfiguration) WithRootVolume(value *KubevirtRootVolumeApplyConfiguration) *KubevirtNodePoolPlatformApplyConfiguration {
	b.RootVolume = value
	return b
}

// WithCompute sets the Compute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Compute field is set to the value of the last call.
func (b *KubevirtNodePoolPlatformApplyConfiguration) WithCompute(value *KubevirtComputeApplyConfiguration) *KubevirtNodePoolPlatformApplyConfiguration {
	b.Compute = value
	return b
}

// WithNetworkInterfaceMultiQueue sets the NetworkInterfaceMultiQueue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkInterfaceMultiQueue field is set to the value of the last call.
func (b *KubevirtNodePoolPlatformApplyConfiguration) WithNetworkInterfaceMultiQueue(value hypershiftv1beta1.MultiQueueSetting) *KubevirtNodePoolPlatformApplyConfiguration {
	b.NetworkInterfaceMultiQueue = &value
	return b
}

// WithAdditionalNetworks adds the given value to the AdditionalNetworks field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalNetworks field.
func (b *KubevirtNodePoolPlatformApplyConfiguration) WithAdditionalNetworks(values ...*KubevirtNetworkApplyConfiguration) *KubevirtNodePoolPlatformApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdditionalNetworks")
		}
		b.AdditionalNetworks = append(b.AdditionalNetworks, *values[i])
	}
	return b
}

// WithAttachDefaultNetwork sets the AttachDefaultNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AttachDefaultNetwork field is set to the value of the last call.
func (b *KubevirtNodePoolPlatformApplyConfiguration) WithAttachDefaultNetwork(value bool) *KubevirtNodePoolPlatformApplyConfiguration {
	b.AttachDefaultNetwork = &value
	return b
}