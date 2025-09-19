/*
Copyright 2025.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MigratedbSpec defines the desired state of Migratedb.

type SourceDatabase struct {
	DbURL               string                  `json:"dbUrl"`
	DbPort              int                     `json:"dbPort"`
	DbName              string                  `json:"dbName"`
	DbCredentialsSecret *corev1.SecretReference `json:"dbCredentialsSecret"`
	Type                string                  `json:"type"`
}

// DestinationDatabase represents the destination database configuration
type DestinationDatabase struct {
	DbURL               string                  `json:"dbUrl"`
	DbPort              int                     `json:"dbPort"`
	DbName              string                  `json:"dbName"`
	DbCredentialsSecret *corev1.SecretReference `json:"dbCredentialsSecret"`
	Type                string                  `json:"type"`
}
type MigratedbSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Migratedb. Edit migratedb_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	SourceDb      SourceDatabase        `json:"sourceDb"`
	DestinationDb []DestinationDatabase `json:"destinationDb"`
	IncludeTable  []string              `json:"includeTable"`
	ExcludeTable  []string              `json:"excludeTable"`
}

// MigratedbStatus defines the observed state of Migratedb.
type MigratedbStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	MigrationStatus string `json:"migrationStatus,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Migratedb is the Schema for the migratedbs API.
type Migratedb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   MigratedbSpec   `json:"spec,omitempty"`
	Status MigratedbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MigratedbList contains a list of Migratedb.
type MigratedbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Migratedb `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Migratedb{}, &MigratedbList{})
}
