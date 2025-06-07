/*
Copyright 2025 rickyalex88@gmail.com.

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

func init() {
	SchemeBuilder.Register(
		&MongoDBBackup{},
		&MongoDBBackupList{},
	)
}

// MongoDBBackupSpec defines the desired state
type MongoDBBackupSpec struct {
	DatabaseURI string `json:"databaseURI"`
	BackupPath  string `json:"backupPath"`
	Schedule    string `json:"schedule"`
}

// MongoDBBackupStatus defines the observed state
type MongoDBBackupStatus struct {
	LastBackupTime metav1.Time `json:"lastBackupTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MongoDBBackup is the Schema for the mongodbbackups API
type MongoDBBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBBackupSpec   `json:"spec,omitempty"`
	Status MongoDBBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDBBackupList contains a list of MongoDBBackup
type MongoDBBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBBackup `json:"items"`
}

