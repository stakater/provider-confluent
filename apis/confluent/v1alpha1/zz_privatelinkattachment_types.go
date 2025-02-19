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

type PrivateLinkAttachmentAwsInitParameters struct {
}

type PrivateLinkAttachmentAwsObservation struct {

	// AWS VPC Endpoint Service that can be used to establish connections for all zones, for example com.amazonaws.vpce.us-west-2.vpce-svc-0d3be37e21708ecd3.
	VPCEndpointServiceName *string `json:"vpcEndpointServiceName,omitempty" tf:"vpc_endpoint_service_name,omitempty"`
}

type PrivateLinkAttachmentAwsParameters struct {
}

type PrivateLinkAttachmentAzureInitParameters struct {
}

type PrivateLinkAttachmentAzureObservation struct {
	PrivateLinkServiceAlias *string `json:"privateLinkServiceAlias,omitempty" tf:"private_link_service_alias,omitempty"`

	// The ID of the Environment that the Private Link Attachment belongs to, for example env-xyz456.
	PrivateLinkServiceResourceID *string `json:"privateLinkServiceResourceId,omitempty" tf:"private_link_service_resource_id,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateLinkAttachmentAzureParameters struct {
}

type PrivateLinkAttachmentEnvironmentInitParameters struct {
}

type PrivateLinkAttachmentEnvironmentObservation struct {

	// The ID of the Environment that the Private Link Attachment belongs to, for example env-xyz456.
	// The unique identifier for the environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateLinkAttachmentEnvironmentParameters struct {

	// The ID of the Environment that the Private Link Attachment belongs to, for example env-xyz456.
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

type PrivateLinkAttachmentGCPInitParameters struct {
}

type PrivateLinkAttachmentGCPObservation struct {
	PrivateServiceConnectServiceAttachment *string `json:"privateServiceConnectServiceAttachment,omitempty" tf:"private_service_connect_service_attachment,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateLinkAttachmentGCPParameters struct {
}

type PrivateLinkAttachmentInitParameters struct {

	// The cloud service provider that hosts the resources to access with the Private Link Attachment.
	// The cloud service provider that hosts the resources to access with the PrivateLink attachment.
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// The name of the Private Link Attachment.
	// The name of the Private Link Attachment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// supports the following:
	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	Environment []PrivateLinkAttachmentEnvironmentInitParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// The cloud service provider region where the resources to be accessed using the Private Link Attachment are located.
	// The cloud service provider region where the resources to be accessed using the PrivateLink attachment are located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PrivateLinkAttachmentObservation struct {

	// supports the following:
	Aws []PrivateLinkAttachmentAwsObservation `json:"aws,omitempty" tf:"aws,omitempty"`

	Azure []PrivateLinkAttachmentAzureObservation `json:"azure,omitempty" tf:"azure,omitempty"`

	// The cloud service provider that hosts the resources to access with the Private Link Attachment.
	// The cloud service provider that hosts the resources to access with the PrivateLink attachment.
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// The root DNS domain for the Private Link Attachment, for example, `pr123a.us-east-2.aws.confluent.
	// The root DNS domain for the private link attachment.
	DNSDomain *string `json:"dnsDomain,omitempty" tf:"dns_domain,omitempty"`

	// The name of the Private Link Attachment.
	// The name of the Private Link Attachment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// supports the following:
	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	Environment []PrivateLinkAttachmentEnvironmentObservation `json:"environment,omitempty" tf:"environment,omitempty"`

	GCP []PrivateLinkAttachmentGCPObservation `json:"gcp,omitempty" tf:"gcp,omitempty"`

	// The ID of the Environment that the Private Link Attachment belongs to, for example env-xyz456.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The cloud service provider region where the resources to be accessed using the Private Link Attachment are located.
	// The cloud service provider region where the resources to be accessed using the PrivateLink attachment are located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Confluent Resource Name of the Private Link Attachment, for example crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-75gxp2/private-link-attachment=platt-1q0ky0.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`
}

type PrivateLinkAttachmentParameters struct {

	// The cloud service provider that hosts the resources to access with the Private Link Attachment.
	// The cloud service provider that hosts the resources to access with the PrivateLink attachment.
	// +kubebuilder:validation:Optional
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// The name of the Private Link Attachment.
	// The name of the Private Link Attachment.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// supports the following:
	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	// +kubebuilder:validation:Optional
	Environment []PrivateLinkAttachmentEnvironmentParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// The cloud service provider region where the resources to be accessed using the Private Link Attachment are located.
	// The cloud service provider region where the resources to be accessed using the PrivateLink attachment are located.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// PrivateLinkAttachmentSpec defines the desired state of PrivateLinkAttachment
type PrivateLinkAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLinkAttachmentParameters `json:"forProvider"`
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
	InitProvider PrivateLinkAttachmentInitParameters `json:"initProvider,omitempty"`
}

// PrivateLinkAttachmentStatus defines the observed state of PrivateLinkAttachment.
type PrivateLinkAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLinkAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkAttachment is the Schema for the PrivateLinkAttachments API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type PrivateLinkAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cloud) || (has(self.initProvider) && has(self.initProvider.cloud))",message="spec.forProvider.cloud is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || (has(self.initProvider) && has(self.initProvider.environment))",message="spec.forProvider.environment is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	Spec   PrivateLinkAttachmentSpec   `json:"spec"`
	Status PrivateLinkAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkAttachmentList contains a list of PrivateLinkAttachments
type PrivateLinkAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLinkAttachment `json:"items"`
}

// Repository type metadata.
var (
	PrivateLinkAttachment_Kind             = "PrivateLinkAttachment"
	PrivateLinkAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateLinkAttachment_Kind}.String()
	PrivateLinkAttachment_KindAPIVersion   = PrivateLinkAttachment_Kind + "." + CRDGroupVersion.String()
	PrivateLinkAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PrivateLinkAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateLinkAttachment{}, &PrivateLinkAttachmentList{})
}
