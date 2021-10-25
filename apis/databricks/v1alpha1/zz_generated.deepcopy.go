// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomParametersObservation) DeepCopyInto(out *CustomParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomParametersObservation.
func (in *CustomParametersObservation) DeepCopy() *CustomParametersObservation {
	if in == nil {
		return nil
	}
	out := new(CustomParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomParametersParameters) DeepCopyInto(out *CustomParametersParameters) {
	*out = *in
	if in.MachineLearningWorkspaceID != nil {
		in, out := &in.MachineLearningWorkspaceID, &out.MachineLearningWorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.NatGatewayName != nil {
		in, out := &in.NatGatewayName, &out.NatGatewayName
		*out = new(string)
		**out = **in
	}
	if in.NoPublicIP != nil {
		in, out := &in.NoPublicIP, &out.NoPublicIP
		*out = new(bool)
		**out = **in
	}
	if in.PrivateSubnetName != nil {
		in, out := &in.PrivateSubnetName, &out.PrivateSubnetName
		*out = new(string)
		**out = **in
	}
	if in.PrivateSubnetNetworkSecurityGroupAssociationID != nil {
		in, out := &in.PrivateSubnetNetworkSecurityGroupAssociationID, &out.PrivateSubnetNetworkSecurityGroupAssociationID
		*out = new(string)
		**out = **in
	}
	if in.PublicIPName != nil {
		in, out := &in.PublicIPName, &out.PublicIPName
		*out = new(string)
		**out = **in
	}
	if in.PublicSubnetName != nil {
		in, out := &in.PublicSubnetName, &out.PublicSubnetName
		*out = new(string)
		**out = **in
	}
	if in.PublicSubnetNetworkSecurityGroupAssociationID != nil {
		in, out := &in.PublicSubnetNetworkSecurityGroupAssociationID, &out.PublicSubnetNetworkSecurityGroupAssociationID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountName != nil {
		in, out := &in.StorageAccountName, &out.StorageAccountName
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountSkuName != nil {
		in, out := &in.StorageAccountSkuName, &out.StorageAccountSkuName
		*out = new(string)
		**out = **in
	}
	if in.VirtualNetworkID != nil {
		in, out := &in.VirtualNetworkID, &out.VirtualNetworkID
		*out = new(string)
		**out = **in
	}
	if in.VnetAddressPrefix != nil {
		in, out := &in.VnetAddressPrefix, &out.VnetAddressPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomParametersParameters.
func (in *CustomParametersParameters) DeepCopy() *CustomParametersParameters {
	if in == nil {
		return nil
	}
	out := new(CustomParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspace) DeepCopyInto(out *DatabricksWorkspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspace.
func (in *DatabricksWorkspace) DeepCopy() *DatabricksWorkspace {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabricksWorkspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceCustomerManagedKey) DeepCopyInto(out *DatabricksWorkspaceCustomerManagedKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceCustomerManagedKey.
func (in *DatabricksWorkspaceCustomerManagedKey) DeepCopy() *DatabricksWorkspaceCustomerManagedKey {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceCustomerManagedKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabricksWorkspaceCustomerManagedKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceCustomerManagedKeyList) DeepCopyInto(out *DatabricksWorkspaceCustomerManagedKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabricksWorkspaceCustomerManagedKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceCustomerManagedKeyList.
func (in *DatabricksWorkspaceCustomerManagedKeyList) DeepCopy() *DatabricksWorkspaceCustomerManagedKeyList {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceCustomerManagedKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabricksWorkspaceCustomerManagedKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceCustomerManagedKeyObservation) DeepCopyInto(out *DatabricksWorkspaceCustomerManagedKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceCustomerManagedKeyObservation.
func (in *DatabricksWorkspaceCustomerManagedKeyObservation) DeepCopy() *DatabricksWorkspaceCustomerManagedKeyObservation {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceCustomerManagedKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceCustomerManagedKeyParameters) DeepCopyInto(out *DatabricksWorkspaceCustomerManagedKeyParameters) {
	*out = *in
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceCustomerManagedKeyParameters.
func (in *DatabricksWorkspaceCustomerManagedKeyParameters) DeepCopy() *DatabricksWorkspaceCustomerManagedKeyParameters {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceCustomerManagedKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceCustomerManagedKeySpec) DeepCopyInto(out *DatabricksWorkspaceCustomerManagedKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceCustomerManagedKeySpec.
func (in *DatabricksWorkspaceCustomerManagedKeySpec) DeepCopy() *DatabricksWorkspaceCustomerManagedKeySpec {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceCustomerManagedKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceCustomerManagedKeyStatus) DeepCopyInto(out *DatabricksWorkspaceCustomerManagedKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceCustomerManagedKeyStatus.
func (in *DatabricksWorkspaceCustomerManagedKeyStatus) DeepCopy() *DatabricksWorkspaceCustomerManagedKeyStatus {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceCustomerManagedKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceList) DeepCopyInto(out *DatabricksWorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabricksWorkspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceList.
func (in *DatabricksWorkspaceList) DeepCopy() *DatabricksWorkspaceList {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabricksWorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceObservation) DeepCopyInto(out *DatabricksWorkspaceObservation) {
	*out = *in
	if in.ManagedResourceGroupID != nil {
		in, out := &in.ManagedResourceGroupID, &out.ManagedResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountIdentity != nil {
		in, out := &in.StorageAccountIdentity, &out.StorageAccountIdentity
		*out = make([]StorageAccountIdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceURL != nil {
		in, out := &in.WorkspaceURL, &out.WorkspaceURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceObservation.
func (in *DatabricksWorkspaceObservation) DeepCopy() *DatabricksWorkspaceObservation {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceParameters) DeepCopyInto(out *DatabricksWorkspaceParameters) {
	*out = *in
	if in.CustomParameters != nil {
		in, out := &in.CustomParameters, &out.CustomParameters
		*out = make([]CustomParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomerManagedKeyEnabled != nil {
		in, out := &in.CustomerManagedKeyEnabled, &out.CustomerManagedKeyEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InfrastructureEncryptionEnabled != nil {
		in, out := &in.InfrastructureEncryptionEnabled, &out.InfrastructureEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LoadBalancerBackendAddressPoolID != nil {
		in, out := &in.LoadBalancerBackendAddressPoolID, &out.LoadBalancerBackendAddressPoolID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagedResourceGroupName != nil {
		in, out := &in.ManagedResourceGroupName, &out.ManagedResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ManagedServicesCmkKeyVaultKeyID != nil {
		in, out := &in.ManagedServicesCmkKeyVaultKeyID, &out.ManagedServicesCmkKeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkSecurityGroupRulesRequired != nil {
		in, out := &in.NetworkSecurityGroupRulesRequired, &out.NetworkSecurityGroupRulesRequired
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceParameters.
func (in *DatabricksWorkspaceParameters) DeepCopy() *DatabricksWorkspaceParameters {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceSpec) DeepCopyInto(out *DatabricksWorkspaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceSpec.
func (in *DatabricksWorkspaceSpec) DeepCopy() *DatabricksWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabricksWorkspaceStatus) DeepCopyInto(out *DatabricksWorkspaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabricksWorkspaceStatus.
func (in *DatabricksWorkspaceStatus) DeepCopy() *DatabricksWorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(DatabricksWorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAccountIdentityObservation) DeepCopyInto(out *StorageAccountIdentityObservation) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAccountIdentityObservation.
func (in *StorageAccountIdentityObservation) DeepCopy() *StorageAccountIdentityObservation {
	if in == nil {
		return nil
	}
	out := new(StorageAccountIdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAccountIdentityParameters) DeepCopyInto(out *StorageAccountIdentityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAccountIdentityParameters.
func (in *StorageAccountIdentityParameters) DeepCopy() *StorageAccountIdentityParameters {
	if in == nil {
		return nil
	}
	out := new(StorageAccountIdentityParameters)
	in.DeepCopyInto(out)
	return out
}