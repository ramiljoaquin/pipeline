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

type PostLeaderElectionRequest struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip,omitempty"`
}