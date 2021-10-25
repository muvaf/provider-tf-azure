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

// GetCondition of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsFunctionJavascriptUdf.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsFunctionJavascriptUdf) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsFunctionJavascriptUdf.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsFunctionJavascriptUdf) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsFunctionJavascriptUdf.
func (mg *StreamAnalyticsFunctionJavascriptUdf) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsJob.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsJob) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsJob.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsJob) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsJob.
func (mg *StreamAnalyticsJob) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsOutputBlob.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputBlob) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsOutputBlob.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputBlob) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsOutputBlob.
func (mg *StreamAnalyticsOutputBlob) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsOutputEventhub.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputEventhub) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsOutputEventhub.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputEventhub) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsOutputEventhub.
func (mg *StreamAnalyticsOutputEventhub) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsOutputMssql.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputMssql) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsOutputMssql.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputMssql) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsOutputMssql.
func (mg *StreamAnalyticsOutputMssql) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsOutputServicebusQueue.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputServicebusQueue) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsOutputServicebusQueue.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputServicebusQueue) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsOutputServicebusQueue.
func (mg *StreamAnalyticsOutputServicebusQueue) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsOutputServicebusTopic.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputServicebusTopic) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsOutputServicebusTopic.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsOutputServicebusTopic) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsOutputServicebusTopic.
func (mg *StreamAnalyticsOutputServicebusTopic) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsReferenceInputBlob.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsReferenceInputBlob) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsReferenceInputBlob.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsReferenceInputBlob) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsReferenceInputBlob.
func (mg *StreamAnalyticsReferenceInputBlob) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsStreamInputBlob.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsStreamInputBlob) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsStreamInputBlob.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsStreamInputBlob) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsStreamInputBlob.
func (mg *StreamAnalyticsStreamInputBlob) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsStreamInputEventhub.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsStreamInputEventhub) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsStreamInputEventhub.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsStreamInputEventhub) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsStreamInputEventhub.
func (mg *StreamAnalyticsStreamInputEventhub) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this StreamAnalyticsStreamInputIothub.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *StreamAnalyticsStreamInputIothub) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this StreamAnalyticsStreamInputIothub.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *StreamAnalyticsStreamInputIothub) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this StreamAnalyticsStreamInputIothub.
func (mg *StreamAnalyticsStreamInputIothub) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}