/*

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

type ProvisionPhase string

const (
	ProvisionInit       ProvisionPhase = "ProvisionInit"
	ProvisionInProgress ProvisionPhase = "ProvisionInProgress"
	ProvisionSucceed    ProvisionPhase = "ProvisionSucceed"
	ProvisionFailed     ProvisionPhase = "ProvisionFailed"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// K3sSpec defines the desired state of K3s
type K3sSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	IP          string `json:"ip,omitempty"`
	User        string `json:"user,omitempty"`
	SshPort     string `json:"sshPort,omitempty"`
	Type        string `json:"type,omitempty"`
	SkipInstall bool   `json:"skipInstall,omitempty"`
	ServerIP    string `json:"serverIp,omitempty"`
}

// K3sStatus defines the observed state of K3s
type K3sStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Type                string         `json:"type,omitempty"`
	Phase               ProvisionPhase `json:"status,omitempty"`
	LastUpdateTimestamp *metav1.Time   `json:"lastUpdateTimestamp,omitempty"`
}

// +kubebuilder:object:root=true

// K3s is the Schema for the k3s API
type K3s struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   K3sSpec   `json:"spec,omitempty"`
	Status K3sStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K3sList contains a list of K3s
type K3sList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K3s `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K3s{}, &K3sList{})
}
