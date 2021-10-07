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

// GetCondition of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlActiveDirectoryAdministrator.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlActiveDirectoryAdministrator) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlActiveDirectoryAdministrator.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlActiveDirectoryAdministrator) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlActiveDirectoryAdministrator.
func (mg *PostgresqlActiveDirectoryAdministrator) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlConfiguration.
func (mg *PostgresqlConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlFirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlFirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlFirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlFirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlFirewallRule.
func (mg *PostgresqlFirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlFlexibleServer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlFlexibleServer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlFlexibleServer.
func (mg *PostgresqlFlexibleServer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlFlexibleServerConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServerConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlFlexibleServerConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServerConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlFlexibleServerConfiguration.
func (mg *PostgresqlFlexibleServerConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlFlexibleServerDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServerDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlFlexibleServerDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServerDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlFlexibleServerDatabase.
func (mg *PostgresqlFlexibleServerDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlFlexibleServerFirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServerFirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlFlexibleServerFirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlFlexibleServerFirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlFlexibleServerFirewallRule.
func (mg *PostgresqlFlexibleServerFirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlServer.
func (mg *PostgresqlServer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlServer.
func (mg *PostgresqlServer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlServer.
func (mg *PostgresqlServer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlServer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlServer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlServer.
func (mg *PostgresqlServer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlServer.
func (mg *PostgresqlServer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlServer.
func (mg *PostgresqlServer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlServer.
func (mg *PostgresqlServer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlServer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlServer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlServer.
func (mg *PostgresqlServer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlServerKey.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlServerKey) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlServerKey.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlServerKey) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlServerKey.
func (mg *PostgresqlServerKey) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PostgresqlVirtualNetworkRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PostgresqlVirtualNetworkRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PostgresqlVirtualNetworkRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PostgresqlVirtualNetworkRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this PostgresqlVirtualNetworkRule.
func (mg *PostgresqlVirtualNetworkRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
