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

// GetCondition of this HpcCache.
func (mg *HpcCache) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HpcCache.
func (mg *HpcCache) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HpcCache.
func (mg *HpcCache) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HpcCache.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HpcCache) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HpcCache.
func (mg *HpcCache) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HpcCache.
func (mg *HpcCache) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HpcCache.
func (mg *HpcCache) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HpcCache.
func (mg *HpcCache) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HpcCache.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HpcCache) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HpcCache.
func (mg *HpcCache) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HpcCacheAccessPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HpcCacheAccessPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HpcCacheAccessPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HpcCacheAccessPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HpcCacheAccessPolicy.
func (mg *HpcCacheAccessPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HpcCacheBlobNfsTarget.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HpcCacheBlobNfsTarget) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HpcCacheBlobNfsTarget.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HpcCacheBlobNfsTarget) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HpcCacheBlobNfsTarget.
func (mg *HpcCacheBlobNfsTarget) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HpcCacheBlobTarget.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HpcCacheBlobTarget) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HpcCacheBlobTarget.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HpcCacheBlobTarget) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HpcCacheBlobTarget.
func (mg *HpcCacheBlobTarget) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this HpcCacheNfsTarget.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *HpcCacheNfsTarget) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this HpcCacheNfsTarget.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *HpcCacheNfsTarget) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this HpcCacheNfsTarget.
func (mg *HpcCacheNfsTarget) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}