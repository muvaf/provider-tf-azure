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
func (in *RelayHybridConnection) DeepCopyInto(out *RelayHybridConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnection.
func (in *RelayHybridConnection) DeepCopy() *RelayHybridConnection {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayHybridConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionAuthorizationRule) DeepCopyInto(out *RelayHybridConnectionAuthorizationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionAuthorizationRule.
func (in *RelayHybridConnectionAuthorizationRule) DeepCopy() *RelayHybridConnectionAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayHybridConnectionAuthorizationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionAuthorizationRuleList) DeepCopyInto(out *RelayHybridConnectionAuthorizationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RelayHybridConnectionAuthorizationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionAuthorizationRuleList.
func (in *RelayHybridConnectionAuthorizationRuleList) DeepCopy() *RelayHybridConnectionAuthorizationRuleList {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionAuthorizationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayHybridConnectionAuthorizationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionAuthorizationRuleObservation) DeepCopyInto(out *RelayHybridConnectionAuthorizationRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionAuthorizationRuleObservation.
func (in *RelayHybridConnectionAuthorizationRuleObservation) DeepCopy() *RelayHybridConnectionAuthorizationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionAuthorizationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionAuthorizationRuleParameters) DeepCopyInto(out *RelayHybridConnectionAuthorizationRuleParameters) {
	*out = *in
	if in.HybridConnectionName != nil {
		in, out := &in.HybridConnectionName, &out.HybridConnectionName
		*out = new(string)
		**out = **in
	}
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionAuthorizationRuleParameters.
func (in *RelayHybridConnectionAuthorizationRuleParameters) DeepCopy() *RelayHybridConnectionAuthorizationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionAuthorizationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionAuthorizationRuleSpec) DeepCopyInto(out *RelayHybridConnectionAuthorizationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionAuthorizationRuleSpec.
func (in *RelayHybridConnectionAuthorizationRuleSpec) DeepCopy() *RelayHybridConnectionAuthorizationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionAuthorizationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionAuthorizationRuleStatus) DeepCopyInto(out *RelayHybridConnectionAuthorizationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionAuthorizationRuleStatus.
func (in *RelayHybridConnectionAuthorizationRuleStatus) DeepCopy() *RelayHybridConnectionAuthorizationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionAuthorizationRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionList) DeepCopyInto(out *RelayHybridConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RelayHybridConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionList.
func (in *RelayHybridConnectionList) DeepCopy() *RelayHybridConnectionList {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayHybridConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionObservation) DeepCopyInto(out *RelayHybridConnectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionObservation.
func (in *RelayHybridConnectionObservation) DeepCopy() *RelayHybridConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionParameters) DeepCopyInto(out *RelayHybridConnectionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceName != nil {
		in, out := &in.RelayNamespaceName, &out.RelayNamespaceName
		*out = new(string)
		**out = **in
	}
	if in.RequiresClientAuthorization != nil {
		in, out := &in.RequiresClientAuthorization, &out.RequiresClientAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.UserMetadata != nil {
		in, out := &in.UserMetadata, &out.UserMetadata
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionParameters.
func (in *RelayHybridConnectionParameters) DeepCopy() *RelayHybridConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionSpec) DeepCopyInto(out *RelayHybridConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionSpec.
func (in *RelayHybridConnectionSpec) DeepCopy() *RelayHybridConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayHybridConnectionStatus) DeepCopyInto(out *RelayHybridConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayHybridConnectionStatus.
func (in *RelayHybridConnectionStatus) DeepCopy() *RelayHybridConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(RelayHybridConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespace) DeepCopyInto(out *RelayNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespace.
func (in *RelayNamespace) DeepCopy() *RelayNamespace {
	if in == nil {
		return nil
	}
	out := new(RelayNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceAuthorizationRule) DeepCopyInto(out *RelayNamespaceAuthorizationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceAuthorizationRule.
func (in *RelayNamespaceAuthorizationRule) DeepCopy() *RelayNamespaceAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayNamespaceAuthorizationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceAuthorizationRuleList) DeepCopyInto(out *RelayNamespaceAuthorizationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RelayNamespaceAuthorizationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceAuthorizationRuleList.
func (in *RelayNamespaceAuthorizationRuleList) DeepCopy() *RelayNamespaceAuthorizationRuleList {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceAuthorizationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayNamespaceAuthorizationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceAuthorizationRuleObservation) DeepCopyInto(out *RelayNamespaceAuthorizationRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceAuthorizationRuleObservation.
func (in *RelayNamespaceAuthorizationRuleObservation) DeepCopy() *RelayNamespaceAuthorizationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceAuthorizationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceAuthorizationRuleParameters) DeepCopyInto(out *RelayNamespaceAuthorizationRuleParameters) {
	*out = *in
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceAuthorizationRuleParameters.
func (in *RelayNamespaceAuthorizationRuleParameters) DeepCopy() *RelayNamespaceAuthorizationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceAuthorizationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceAuthorizationRuleSpec) DeepCopyInto(out *RelayNamespaceAuthorizationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceAuthorizationRuleSpec.
func (in *RelayNamespaceAuthorizationRuleSpec) DeepCopy() *RelayNamespaceAuthorizationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceAuthorizationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceAuthorizationRuleStatus) DeepCopyInto(out *RelayNamespaceAuthorizationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceAuthorizationRuleStatus.
func (in *RelayNamespaceAuthorizationRuleStatus) DeepCopy() *RelayNamespaceAuthorizationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceAuthorizationRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceList) DeepCopyInto(out *RelayNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RelayNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceList.
func (in *RelayNamespaceList) DeepCopy() *RelayNamespaceList {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RelayNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceObservation) DeepCopyInto(out *RelayNamespaceObservation) {
	*out = *in
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceObservation.
func (in *RelayNamespaceObservation) DeepCopy() *RelayNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceParameters) DeepCopyInto(out *RelayNamespaceParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceParameters.
func (in *RelayNamespaceParameters) DeepCopy() *RelayNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceSpec) DeepCopyInto(out *RelayNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceSpec.
func (in *RelayNamespaceSpec) DeepCopy() *RelayNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelayNamespaceStatus) DeepCopyInto(out *RelayNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelayNamespaceStatus.
func (in *RelayNamespaceStatus) DeepCopy() *RelayNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(RelayNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}