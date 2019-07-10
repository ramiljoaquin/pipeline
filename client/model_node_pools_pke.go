/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type NodePoolsPke struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	// user provided custom node labels to be placed onto the nodes of the node pool
	Labels map[string]string `json:"labels,omitempty"`
	// Enables/disables autoscaling of this node pool through Kubernetes cluster autoscaler.
	Autoscaling    bool                   `json:"autoscaling"`
	Provider       string                 `json:"provider"`
	ProviderConfig map[string]interface{} `json:"providerConfig"`
	Hosts          []PkeHosts             `json:"hosts,omitempty"`
}
