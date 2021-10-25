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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CacheExpirationActionObservation struct {
}

type CacheExpirationActionParameters struct {

	// +kubebuilder:validation:Required
	Behavior *string `json:"behavior" tf:"behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type CacheKeyQueryStringActionObservation struct {
}

type CacheKeyQueryStringActionParameters struct {

	// +kubebuilder:validation:Required
	Behavior *string `json:"behavior" tf:"behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type CdnEndpointObservation struct {
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`
}

type CdnEndpointParameters struct {

	// +kubebuilder:validation:Optional
	ContentTypesToCompress []*string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress,omitempty"`

	// +kubebuilder:validation:Optional
	DeliveryRule []DeliveryRuleParameters `json:"deliveryRule,omitempty" tf:"delivery_rule,omitempty"`

	// +kubebuilder:validation:Optional
	GeoFilter []GeoFilterParameters `json:"geoFilter,omitempty" tf:"geo_filter,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalDeliveryRule []GlobalDeliveryRuleParameters `json:"globalDeliveryRule,omitempty" tf:"global_delivery_rule,omitempty"`

	// +kubebuilder:validation:Optional
	IsCompressionEnabled *bool `json:"isCompressionEnabled,omitempty" tf:"is_compression_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IsHTTPAllowed *bool `json:"isHttpAllowed,omitempty" tf:"is_http_allowed,omitempty"`

	// +kubebuilder:validation:Optional
	IsHTTPSAllowed *bool `json:"isHttpsAllowed,omitempty" tf:"is_https_allowed,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OptimizationType *string `json:"optimizationType,omitempty" tf:"optimization_type,omitempty"`

	// +kubebuilder:validation:Required
	Origin []OriginParameters `json:"origin" tf:"origin,omitempty"`

	// +kubebuilder:validation:Optional
	OriginHostHeader *string `json:"originHostHeader,omitempty" tf:"origin_host_header,omitempty"`

	// +kubebuilder:validation:Optional
	OriginPath *string `json:"originPath,omitempty" tf:"origin_path,omitempty"`

	// +kubebuilder:validation:Optional
	ProbePath *string `json:"probePath,omitempty" tf:"probe_path,omitempty"`

	// +kubebuilder:validation:Required
	ProfileName *string `json:"profileName" tf:"profile_name,omitempty"`

	// +kubebuilder:validation:Optional
	QuerystringCachingBehaviour *string `json:"querystringCachingBehaviour,omitempty" tf:"querystring_caching_behaviour,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CookiesConditionObservation struct {
}

type CookiesConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type DeliveryRuleObservation struct {
}

type DeliveryRuleParameters struct {

	// +kubebuilder:validation:Optional
	CacheExpirationAction []CacheExpirationActionParameters `json:"cacheExpirationAction,omitempty" tf:"cache_expiration_action,omitempty"`

	// +kubebuilder:validation:Optional
	CacheKeyQueryStringAction []CacheKeyQueryStringActionParameters `json:"cacheKeyQueryStringAction,omitempty" tf:"cache_key_query_string_action,omitempty"`

	// +kubebuilder:validation:Optional
	CookiesCondition []CookiesConditionParameters `json:"cookiesCondition,omitempty" tf:"cookies_condition,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceCondition []DeviceConditionParameters `json:"deviceCondition,omitempty" tf:"device_condition,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPVersionCondition []HTTPVersionConditionParameters `json:"httpVersionCondition,omitempty" tf:"http_version_condition,omitempty"`

	// +kubebuilder:validation:Optional
	ModifyRequestHeaderAction []ModifyRequestHeaderActionParameters `json:"modifyRequestHeaderAction,omitempty" tf:"modify_request_header_action,omitempty"`

	// +kubebuilder:validation:Optional
	ModifyResponseHeaderAction []ModifyResponseHeaderActionParameters `json:"modifyResponseHeaderAction,omitempty" tf:"modify_response_header_action,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Order *int64 `json:"order" tf:"order,omitempty"`

	// +kubebuilder:validation:Optional
	PostArgCondition []PostArgConditionParameters `json:"postArgCondition,omitempty" tf:"post_arg_condition,omitempty"`

	// +kubebuilder:validation:Optional
	QueryStringCondition []QueryStringConditionParameters `json:"queryStringCondition,omitempty" tf:"query_string_condition,omitempty"`

	// +kubebuilder:validation:Optional
	RemoteAddressCondition []RemoteAddressConditionParameters `json:"remoteAddressCondition,omitempty" tf:"remote_address_condition,omitempty"`

	// +kubebuilder:validation:Optional
	RequestBodyCondition []RequestBodyConditionParameters `json:"requestBodyCondition,omitempty" tf:"request_body_condition,omitempty"`

	// +kubebuilder:validation:Optional
	RequestHeaderCondition []RequestHeaderConditionParameters `json:"requestHeaderCondition,omitempty" tf:"request_header_condition,omitempty"`

	// +kubebuilder:validation:Optional
	RequestMethodCondition []RequestMethodConditionParameters `json:"requestMethodCondition,omitempty" tf:"request_method_condition,omitempty"`

	// +kubebuilder:validation:Optional
	RequestSchemeCondition []RequestSchemeConditionParameters `json:"requestSchemeCondition,omitempty" tf:"request_scheme_condition,omitempty"`

	// +kubebuilder:validation:Optional
	RequestURICondition []RequestURIConditionParameters `json:"requestUriCondition,omitempty" tf:"request_uri_condition,omitempty"`

	// +kubebuilder:validation:Optional
	URLFileExtensionCondition []URLFileExtensionConditionParameters `json:"urlFileExtensionCondition,omitempty" tf:"url_file_extension_condition,omitempty"`

	// +kubebuilder:validation:Optional
	URLFileNameCondition []URLFileNameConditionParameters `json:"urlFileNameCondition,omitempty" tf:"url_file_name_condition,omitempty"`

	// +kubebuilder:validation:Optional
	URLPathCondition []URLPathConditionParameters `json:"urlPathCondition,omitempty" tf:"url_path_condition,omitempty"`

	// +kubebuilder:validation:Optional
	URLRedirectAction []URLRedirectActionParameters `json:"urlRedirectAction,omitempty" tf:"url_redirect_action,omitempty"`

	// +kubebuilder:validation:Optional
	URLRewriteAction []URLRewriteActionParameters `json:"urlRewriteAction,omitempty" tf:"url_rewrite_action,omitempty"`
}

type DeviceConditionObservation struct {
}

type DeviceConditionParameters struct {

	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type GeoFilterObservation struct {
}

type GeoFilterParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	CountryCodes []*string `json:"countryCodes" tf:"country_codes,omitempty"`

	// +kubebuilder:validation:Required
	RelativePath *string `json:"relativePath" tf:"relative_path,omitempty"`
}

type GlobalDeliveryRuleCacheExpirationActionObservation struct {
}

type GlobalDeliveryRuleCacheExpirationActionParameters struct {

	// +kubebuilder:validation:Required
	Behavior *string `json:"behavior" tf:"behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type GlobalDeliveryRuleCacheKeyQueryStringActionObservation struct {
}

type GlobalDeliveryRuleCacheKeyQueryStringActionParameters struct {

	// +kubebuilder:validation:Required
	Behavior *string `json:"behavior" tf:"behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type GlobalDeliveryRuleModifyRequestHeaderActionObservation struct {
}

type GlobalDeliveryRuleModifyRequestHeaderActionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GlobalDeliveryRuleModifyResponseHeaderActionObservation struct {
}

type GlobalDeliveryRuleModifyResponseHeaderActionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GlobalDeliveryRuleObservation struct {
}

type GlobalDeliveryRuleParameters struct {

	// +kubebuilder:validation:Optional
	CacheExpirationAction []GlobalDeliveryRuleCacheExpirationActionParameters `json:"cacheExpirationAction,omitempty" tf:"cache_expiration_action,omitempty"`

	// +kubebuilder:validation:Optional
	CacheKeyQueryStringAction []GlobalDeliveryRuleCacheKeyQueryStringActionParameters `json:"cacheKeyQueryStringAction,omitempty" tf:"cache_key_query_string_action,omitempty"`

	// +kubebuilder:validation:Optional
	ModifyRequestHeaderAction []GlobalDeliveryRuleModifyRequestHeaderActionParameters `json:"modifyRequestHeaderAction,omitempty" tf:"modify_request_header_action,omitempty"`

	// +kubebuilder:validation:Optional
	ModifyResponseHeaderAction []GlobalDeliveryRuleModifyResponseHeaderActionParameters `json:"modifyResponseHeaderAction,omitempty" tf:"modify_response_header_action,omitempty"`

	// +kubebuilder:validation:Optional
	URLRedirectAction []GlobalDeliveryRuleURLRedirectActionParameters `json:"urlRedirectAction,omitempty" tf:"url_redirect_action,omitempty"`

	// +kubebuilder:validation:Optional
	URLRewriteAction []GlobalDeliveryRuleURLRewriteActionParameters `json:"urlRewriteAction,omitempty" tf:"url_rewrite_action,omitempty"`
}

type GlobalDeliveryRuleURLRedirectActionObservation struct {
}

type GlobalDeliveryRuleURLRedirectActionParameters struct {

	// +kubebuilder:validation:Optional
	Fragment *string `json:"fragment,omitempty" tf:"fragment,omitempty"`

	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	QueryString *string `json:"queryString,omitempty" tf:"query_string,omitempty"`

	// +kubebuilder:validation:Required
	RedirectType *string `json:"redirectType" tf:"redirect_type,omitempty"`
}

type GlobalDeliveryRuleURLRewriteActionObservation struct {
}

type GlobalDeliveryRuleURLRewriteActionParameters struct {

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	PreserveUnmatchedPath *bool `json:"preserveUnmatchedPath,omitempty" tf:"preserve_unmatched_path,omitempty"`

	// +kubebuilder:validation:Required
	SourcePattern *string `json:"sourcePattern" tf:"source_pattern,omitempty"`
}

type HTTPVersionConditionObservation struct {
}

type HTTPVersionConditionParameters struct {

	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type ModifyRequestHeaderActionObservation struct {
}

type ModifyRequestHeaderActionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ModifyResponseHeaderActionObservation struct {
}

type ModifyResponseHeaderActionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OriginObservation struct {
}

type OriginParameters struct {

	// +kubebuilder:validation:Optional
	HTTPPort *int64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPSPort *int64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type PostArgConditionObservation struct {
}

type PostArgConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type QueryStringConditionObservation struct {
}

type QueryStringConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type RemoteAddressConditionObservation struct {
}

type RemoteAddressConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`
}

type RequestBodyConditionObservation struct {
}

type RequestBodyConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type RequestHeaderConditionObservation struct {
}

type RequestHeaderConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type RequestMethodConditionObservation struct {
}

type RequestMethodConditionParameters struct {

	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type RequestSchemeConditionObservation struct {
}

type RequestSchemeConditionParameters struct {

	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type RequestURIConditionObservation struct {
}

type RequestURIConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type URLFileExtensionConditionObservation struct {
}

type URLFileExtensionConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type URLFileNameConditionObservation struct {
}

type URLFileNameConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type URLPathConditionObservation struct {
}

type URLPathConditionParameters struct {

	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type URLRedirectActionObservation struct {
}

type URLRedirectActionParameters struct {

	// +kubebuilder:validation:Optional
	Fragment *string `json:"fragment,omitempty" tf:"fragment,omitempty"`

	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	QueryString *string `json:"queryString,omitempty" tf:"query_string,omitempty"`

	// +kubebuilder:validation:Required
	RedirectType *string `json:"redirectType" tf:"redirect_type,omitempty"`
}

type URLRewriteActionObservation struct {
}

type URLRewriteActionParameters struct {

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	PreserveUnmatchedPath *bool `json:"preserveUnmatchedPath,omitempty" tf:"preserve_unmatched_path,omitempty"`

	// +kubebuilder:validation:Required
	SourcePattern *string `json:"sourcePattern" tf:"source_pattern,omitempty"`
}

// CdnEndpointSpec defines the desired state of CdnEndpoint
type CdnEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CdnEndpointParameters `json:"forProvider"`
}

// CdnEndpointStatus defines the observed state of CdnEndpoint.
type CdnEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CdnEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CdnEndpoint is the Schema for the CdnEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CdnEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnEndpointSpec   `json:"spec"`
	Status            CdnEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CdnEndpointList contains a list of CdnEndpoints
type CdnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CdnEndpoint `json:"items"`
}

// Repository type metadata.
var (
	CdnEndpointKind             = "CdnEndpoint"
	CdnEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: CdnEndpointKind}.String()
	CdnEndpointKindAPIVersion   = CdnEndpointKind + "." + GroupVersion.String()
	CdnEndpointGroupVersionKind = GroupVersion.WithKind(CdnEndpointKind)
)

func init() {
	SchemeBuilder.Register(&CdnEndpoint{}, &CdnEndpointList{})
}