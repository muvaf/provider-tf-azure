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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DatabricksWorkspace.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DatabricksWorkspace) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DatabricksWorkspace.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DatabricksWorkspace) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DatabricksWorkspace.
func (mg *DatabricksWorkspace) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DatabricksWorkspaceCustomerManagedKey.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DatabricksWorkspaceCustomerManagedKey) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DatabricksWorkspaceCustomerManagedKey.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DatabricksWorkspaceCustomerManagedKey) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DatabricksWorkspaceCustomerManagedKey.
func (mg *DatabricksWorkspaceCustomerManagedKey) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}