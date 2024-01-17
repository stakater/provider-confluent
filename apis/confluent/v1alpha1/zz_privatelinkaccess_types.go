// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivateLinkAccessAwsInitParameters struct {

	// The AWS account ID to enable for the Private Link Access. You can find your AWS account ID [here] (https://console.aws.amazon.com/billing/home?#/account) under My Account in your AWS Management Console. Must be a 12 character string.
	// AWS Account ID to allow for PrivateLink access. Find here (https://console.aws.amazon.com/billing/home?#/account) under My Account in your AWS Management Console.
	Account *string `json:"account,omitempty" tf:"account,omitempty"`
}

type PrivateLinkAccessAwsObservation struct {

	// The AWS account ID to enable for the Private Link Access. You can find your AWS account ID [here] (https://console.aws.amazon.com/billing/home?#/account) under My Account in your AWS Management Console. Must be a 12 character string.
	// AWS Account ID to allow for PrivateLink access. Find here (https://console.aws.amazon.com/billing/home?#/account) under My Account in your AWS Management Console.
	Account *string `json:"account,omitempty" tf:"account,omitempty"`
}

type PrivateLinkAccessAwsParameters struct {

	// The AWS account ID to enable for the Private Link Access. You can find your AWS account ID [here] (https://console.aws.amazon.com/billing/home?#/account) under My Account in your AWS Management Console. Must be a 12 character string.
	// AWS Account ID to allow for PrivateLink access. Find here (https://console.aws.amazon.com/billing/home?#/account) under My Account in your AWS Management Console.
	// +kubebuilder:validation:Optional
	Account *string `json:"account" tf:"account,omitempty"`
}

type PrivateLinkAccessAzureInitParameters struct {

	// The Azure subscription ID to enable for the Private Link Access. You can find your Azure subscription ID in the subscription section of your [Microsoft Azure Portal] (https://portal.azure.com/#blade/Microsoft_Azure_Billing/SubscriptionsBlade). Must be a valid 32 character UUID string.
	// Azure subscription to allow for PrivateLink access.
	Subscription *string `json:"subscription,omitempty" tf:"subscription,omitempty"`
}

type PrivateLinkAccessAzureObservation struct {

	// The Azure subscription ID to enable for the Private Link Access. You can find your Azure subscription ID in the subscription section of your [Microsoft Azure Portal] (https://portal.azure.com/#blade/Microsoft_Azure_Billing/SubscriptionsBlade). Must be a valid 32 character UUID string.
	// Azure subscription to allow for PrivateLink access.
	Subscription *string `json:"subscription,omitempty" tf:"subscription,omitempty"`
}

type PrivateLinkAccessAzureParameters struct {

	// The Azure subscription ID to enable for the Private Link Access. You can find your Azure subscription ID in the subscription section of your [Microsoft Azure Portal] (https://portal.azure.com/#blade/Microsoft_Azure_Billing/SubscriptionsBlade). Must be a valid 32 character UUID string.
	// Azure subscription to allow for PrivateLink access.
	// +kubebuilder:validation:Optional
	Subscription *string `json:"subscription" tf:"subscription,omitempty"`
}

type PrivateLinkAccessEnvironmentInitParameters struct {
}

type PrivateLinkAccessEnvironmentObservation struct {

	// The ID of the Environment that the Private Link Access belongs to, for example, env-abc123.
	// The unique identifier for the environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateLinkAccessEnvironmentParameters struct {

	// The ID of the Environment that the Private Link Access belongs to, for example, env-abc123.
	// The unique identifier for the environment.
	// +crossplane:generate:reference:type=github.com/stakater/provider-confluent/apis/confluent/v1alpha1.Environment
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Environment in confluent to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Environment in confluent to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type PrivateLinkAccessGCPInitParameters struct {

	// The GCP project ID to allow for Private Service Connect access. You can find your Google Cloud Project ID under Project ID section of your Google Cloud Console dashboard.
	// The GCP project ID to allow for Private Service Connect access.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type PrivateLinkAccessGCPObservation struct {

	// The GCP project ID to allow for Private Service Connect access. You can find your Google Cloud Project ID under Project ID section of your Google Cloud Console dashboard.
	// The GCP project ID to allow for Private Service Connect access.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type PrivateLinkAccessGCPParameters struct {

	// The GCP project ID to allow for Private Service Connect access. You can find your Google Cloud Project ID under Project ID section of your Google Cloud Console dashboard.
	// The GCP project ID to allow for Private Service Connect access.
	// +kubebuilder:validation:Optional
	Project *string `json:"project" tf:"project,omitempty"`
}

type PrivateLinkAccessInitParameters struct {

	// The AWS-specific Private Link Access details if available. It supports the following:
	Aws []PrivateLinkAccessAwsInitParameters `json:"aws,omitempty" tf:"aws,omitempty"`

	// The Azure-specific Private Link Access details if available. It supports the following:
	Azure []PrivateLinkAccessAzureInitParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// The name of the Private Link Access.
	// The name of the PrivateLink access.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// supports the following:
	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	Environment []PrivateLinkAccessEnvironmentInitParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// The GCP-specific Private Service Connect details if available. It supports the following:
	GCP []PrivateLinkAccessGCPInitParameters `json:"gcp,omitempty" tf:"gcp,omitempty"`

	// supports the following:
	// Network represents a network (VPC) in Confluent Cloud. All Networks exist within Confluent-managed cloud provider accounts.
	Network []PrivateLinkAccessNetworkInitParameters `json:"network,omitempty" tf:"network,omitempty"`
}

type PrivateLinkAccessNetworkInitParameters struct {
}

type PrivateLinkAccessNetworkObservation struct {

	// The ID of the Environment that the Private Link Access belongs to, for example, env-abc123.
	// The unique identifier for the network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateLinkAccessNetworkParameters struct {

	// The ID of the Environment that the Private Link Access belongs to, for example, env-abc123.
	// The unique identifier for the network.
	// +crossplane:generate:reference:type=github.com/stakater/provider-confluent/apis/confluent/v1alpha1.Network
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Network in confluent to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Network in confluent to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type PrivateLinkAccessObservation struct {

	// The AWS-specific Private Link Access details if available. It supports the following:
	Aws []PrivateLinkAccessAwsObservation `json:"aws,omitempty" tf:"aws,omitempty"`

	// The Azure-specific Private Link Access details if available. It supports the following:
	Azure []PrivateLinkAccessAzureObservation `json:"azure,omitempty" tf:"azure,omitempty"`

	// The name of the Private Link Access.
	// The name of the PrivateLink access.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// supports the following:
	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	Environment []PrivateLinkAccessEnvironmentObservation `json:"environment,omitempty" tf:"environment,omitempty"`

	// The GCP-specific Private Service Connect details if available. It supports the following:
	GCP []PrivateLinkAccessGCPObservation `json:"gcp,omitempty" tf:"gcp,omitempty"`

	// The ID of the Environment that the Private Link Access belongs to, for example, env-abc123.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// supports the following:
	// Network represents a network (VPC) in Confluent Cloud. All Networks exist within Confluent-managed cloud provider accounts.
	Network []PrivateLinkAccessNetworkObservation `json:"network,omitempty" tf:"network,omitempty"`
}

type PrivateLinkAccessParameters struct {

	// The AWS-specific Private Link Access details if available. It supports the following:
	// +kubebuilder:validation:Optional
	Aws []PrivateLinkAccessAwsParameters `json:"aws,omitempty" tf:"aws,omitempty"`

	// The Azure-specific Private Link Access details if available. It supports the following:
	// +kubebuilder:validation:Optional
	Azure []PrivateLinkAccessAzureParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// The name of the Private Link Access.
	// The name of the PrivateLink access.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// supports the following:
	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	// +kubebuilder:validation:Optional
	Environment []PrivateLinkAccessEnvironmentParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// The GCP-specific Private Service Connect details if available. It supports the following:
	// +kubebuilder:validation:Optional
	GCP []PrivateLinkAccessGCPParameters `json:"gcp,omitempty" tf:"gcp,omitempty"`

	// supports the following:
	// Network represents a network (VPC) in Confluent Cloud. All Networks exist within Confluent-managed cloud provider accounts.
	// +kubebuilder:validation:Optional
	Network []PrivateLinkAccessNetworkParameters `json:"network,omitempty" tf:"network,omitempty"`
}

// PrivateLinkAccessSpec defines the desired state of PrivateLinkAccess
type PrivateLinkAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLinkAccessParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PrivateLinkAccessInitParameters `json:"initProvider,omitempty"`
}

// PrivateLinkAccessStatus defines the observed state of PrivateLinkAccess.
type PrivateLinkAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLinkAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkAccess is the Schema for the PrivateLinkAccesss API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type PrivateLinkAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || (has(self.initProvider) && has(self.initProvider.environment))",message="spec.forProvider.environment is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.network) || (has(self.initProvider) && has(self.initProvider.network))",message="spec.forProvider.network is a required parameter"
	Spec   PrivateLinkAccessSpec   `json:"spec"`
	Status PrivateLinkAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkAccessList contains a list of PrivateLinkAccesss
type PrivateLinkAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLinkAccess `json:"items"`
}

// Repository type metadata.
var (
	PrivateLinkAccess_Kind             = "PrivateLinkAccess"
	PrivateLinkAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateLinkAccess_Kind}.String()
	PrivateLinkAccess_KindAPIVersion   = PrivateLinkAccess_Kind + "." + CRDGroupVersion.String()
	PrivateLinkAccess_GroupVersionKind = CRDGroupVersion.WithKind(PrivateLinkAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateLinkAccess{}, &PrivateLinkAccessList{})
}
