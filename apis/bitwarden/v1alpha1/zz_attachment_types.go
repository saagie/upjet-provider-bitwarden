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

type AttachmentInitParameters struct {

	// (String) Path to the content of the attachment.
	// Path to the content of the attachment.
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// (String) Identifier of the item the attachment belongs to
	// Identifier of the item the attachment belongs to
	ItemID *string `json:"itemId,omitempty" tf:"item_id,omitempty"`
}

type AttachmentObservation struct {

	// (String) Path to the content of the attachment.
	// Path to the content of the attachment.
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// (String) File name
	// File name
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// (String) Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Identifier of the item the attachment belongs to
	// Identifier of the item the attachment belongs to
	ItemID *string `json:"itemId,omitempty" tf:"item_id,omitempty"`

	// (String) Size in bytes
	// Size in bytes
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// (String) Size as string
	// Size as string
	SizeName *string `json:"sizeName,omitempty" tf:"size_name,omitempty"`

	// (String) URL
	// URL
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type AttachmentParameters struct {

	// (String) Path to the content of the attachment.
	// Path to the content of the attachment.
	// +kubebuilder:validation:Optional
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// (String) Identifier of the item the attachment belongs to
	// Identifier of the item the attachment belongs to
	// +kubebuilder:validation:Optional
	ItemID *string `json:"itemId,omitempty" tf:"item_id,omitempty"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
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
	InitProvider AttachmentInitParameters `json:"initProvider,omitempty"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Attachment is the Schema for the Attachments API. (EXPERIMENTAL) Manages a Vault item's attachment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,bitwarden}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.file) || (has(self.initProvider) && has(self.initProvider.file))",message="spec.forProvider.file is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.itemId) || (has(self.initProvider) && has(self.initProvider.itemId))",message="spec.forProvider.itemId is a required parameter"
	Spec   AttachmentSpec   `json:"spec"`
	Status AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	Attachment_Kind             = "Attachment"
	Attachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attachment_Kind}.String()
	Attachment_KindAPIVersion   = Attachment_Kind + "." + CRDGroupVersion.String()
	Attachment_GroupVersionKind = CRDGroupVersion.WithKind(Attachment_Kind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}
