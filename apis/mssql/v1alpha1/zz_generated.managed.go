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

// GetCondition of this MssqlDatabase.
func (mg *MssqlDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlDatabase.
func (mg *MssqlDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlDatabase.
func (mg *MssqlDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlDatabase.
func (mg *MssqlDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlDatabase.
func (mg *MssqlDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlDatabase.
func (mg *MssqlDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlDatabase.
func (mg *MssqlDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlDatabase.
func (mg *MssqlDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlDatabaseExtendedAuditingPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlDatabaseExtendedAuditingPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlDatabaseExtendedAuditingPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlDatabaseExtendedAuditingPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlDatabaseExtendedAuditingPolicy.
func (mg *MssqlDatabaseExtendedAuditingPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlDatabaseVulnerabilityAssessmentRuleBaseline.
func (mg *MssqlDatabaseVulnerabilityAssessmentRuleBaseline) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlElasticpool.
func (mg *MssqlElasticpool) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlElasticpool.
func (mg *MssqlElasticpool) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlElasticpool.
func (mg *MssqlElasticpool) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlElasticpool.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlElasticpool) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlElasticpool.
func (mg *MssqlElasticpool) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlElasticpool.
func (mg *MssqlElasticpool) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlElasticpool.
func (mg *MssqlElasticpool) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlElasticpool.
func (mg *MssqlElasticpool) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlElasticpool.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlElasticpool) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlElasticpool.
func (mg *MssqlElasticpool) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlFailoverGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlFailoverGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlFailoverGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlFailoverGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlFailoverGroup.
func (mg *MssqlFailoverGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlFirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlFirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlFirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlFirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlFirewallRule.
func (mg *MssqlFirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlJobAgent.
func (mg *MssqlJobAgent) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlJobAgent.
func (mg *MssqlJobAgent) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlJobAgent.
func (mg *MssqlJobAgent) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlJobAgent.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlJobAgent) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlJobAgent.
func (mg *MssqlJobAgent) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlJobAgent.
func (mg *MssqlJobAgent) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlJobAgent.
func (mg *MssqlJobAgent) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlJobAgent.
func (mg *MssqlJobAgent) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlJobAgent.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlJobAgent) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlJobAgent.
func (mg *MssqlJobAgent) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlJobCredential.
func (mg *MssqlJobCredential) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlJobCredential.
func (mg *MssqlJobCredential) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlJobCredential.
func (mg *MssqlJobCredential) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlJobCredential.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlJobCredential) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlJobCredential.
func (mg *MssqlJobCredential) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlJobCredential.
func (mg *MssqlJobCredential) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlJobCredential.
func (mg *MssqlJobCredential) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlJobCredential.
func (mg *MssqlJobCredential) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlJobCredential.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlJobCredential) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlJobCredential.
func (mg *MssqlJobCredential) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlServer.
func (mg *MssqlServer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlServer.
func (mg *MssqlServer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlServer.
func (mg *MssqlServer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlServer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlServer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlServer.
func (mg *MssqlServer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlServer.
func (mg *MssqlServer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlServer.
func (mg *MssqlServer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlServer.
func (mg *MssqlServer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlServer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlServer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlServer.
func (mg *MssqlServer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlServerSecurityAlertPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlServerSecurityAlertPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlServerSecurityAlertPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlServerSecurityAlertPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlServerSecurityAlertPolicy.
func (mg *MssqlServerSecurityAlertPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlServerTransparentDataEncryption.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlServerTransparentDataEncryption) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlServerTransparentDataEncryption.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlServerTransparentDataEncryption) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlServerTransparentDataEncryption.
func (mg *MssqlServerTransparentDataEncryption) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlServerVulnerabilityAssessment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlServerVulnerabilityAssessment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlServerVulnerabilityAssessment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlServerVulnerabilityAssessment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlServerVulnerabilityAssessment.
func (mg *MssqlServerVulnerabilityAssessment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlVirtualMachine.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlVirtualMachine) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlVirtualMachine.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlVirtualMachine) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlVirtualMachine.
func (mg *MssqlVirtualMachine) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MssqlVirtualNetworkRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MssqlVirtualNetworkRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MssqlVirtualNetworkRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MssqlVirtualNetworkRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MssqlVirtualNetworkRule.
func (mg *MssqlVirtualNetworkRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}