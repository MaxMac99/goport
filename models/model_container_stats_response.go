/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ``` 
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// ContainerStatsResponse - OK response to ContainerTop operation
type ContainerStatsResponse struct {

	// The ID of the container
	Id string `json:"Id,omitempty"`

	// The Name of the container
	Name string `json:"Name,omitempty"`

	Read time.Time `json:"Read,omitempty"`

	PreRead time.Time `json:"PreRead,omitempty"`

	PidsStats ContainersIdStatsPidsStats `json:"PidsStats,omitempty"`

	BlkioStats BlkioStats `json:"BlkioStats,omitempty"`

	CPUStats CpuStats `json:"CPUStats,omitempty"`

	PreCPUStats CpuStats `json:"PreCPUStats,omitempty"`

	MemoryStats MemoryStats `json:"MemoryStats,omitempty"`

	Networks ContainersIdStatsNetworks `json:"Networks,omitempty"`
}
