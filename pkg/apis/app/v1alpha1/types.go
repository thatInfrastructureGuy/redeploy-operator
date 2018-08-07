package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RedeployList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Redeploy `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Redeploy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              RedeploySpec   `json:"spec"`
	Status            RedeployStatus `json:"status,omitempty"`
}

type RedeploySpec struct {
	// Fill me
	RedeployNeeded      bool   `json: ", redeployNeeded"`
	DeploymentName      string `json: ", deploymentName"`
	DeploymentNamespace string `json: ", deploymentNamespace"`
}
type RedeployStatus struct {
	// Fill me
	Status string `json: ", status"`
	Date   string `json: ", status"`
}
