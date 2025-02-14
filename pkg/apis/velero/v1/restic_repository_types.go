/*
Copyright 2018 the Velero contributors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResticRepositorySpec is the specification for a ResticRepository.
type ResticRepositorySpec struct {
	// VolumeNamespace is the namespace this restic repository contains
	// pod volume backups for.
	VolumeNamespace string `json:"volumeNamespace"`

	// BackupStorageLocation is the name of the BackupStorageLocation
	// that should contain this repository.
	BackupStorageLocation string `json:"backupStorageLocation"`

	// ResticIdentifier is the full restic-compatible string for identifying
	// this repository.
	ResticIdentifier string `json:"resticIdentifier"`

	// MaintenanceFrequency is how often maintenance should be run.
	MaintenanceFrequency metav1.Duration `json:"maintenanceFrequency"`
}

// ResticRepositoryPhase represents the lifecycle phase of a ResticRepository.
// +kubebuilder:validation:Enum=New;Ready;NotReady
type ResticRepositoryPhase string

const (
	ResticRepositoryPhaseNew      ResticRepositoryPhase = "New"
	ResticRepositoryPhaseReady    ResticRepositoryPhase = "Ready"
	ResticRepositoryPhaseNotReady ResticRepositoryPhase = "NotReady"
)

// ResticRepositoryStatus is the current status of a ResticRepository.
type ResticRepositoryStatus struct {
	// Phase is the current state of the ResticRepository.
	// +optional
	Phase ResticRepositoryPhase `json:"phase,omitempty"`

	// Message is a message about the current status of the ResticRepository.
	// +optional
	Message string `json:"message,omitempty"`

	// LastMaintenanceTime is the last time maintenance was run.
	// +optional
	// +nullable
	LastMaintenanceTime *metav1.Time `json:"lastMaintenanceTime,omitempty"`
}

// TODO(2.0) After converting all resources to use the runtime-controller client,
// the genclient and k8s:deepcopy markers will no longer be needed and should be removed.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

type ResticRepository struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec ResticRepositorySpec `json:"spec,omitempty"`

	// +optional
	Status ResticRepositoryStatus `json:"status,omitempty"`
}

// TODO(2.0) After converting all resources to use the runtime-controller client,
// the k8s:deepcopy marker will no longer be needed and should be removed.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:rbac:groups=velero.io,resources=resticrepositories,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=velero.io,resources=resticrepositories/status,verbs=get;update;patch

// ResticRepositoryList is a list of ResticRepositories.
type ResticRepositoryList struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ResticRepository `json:"items"`
}
