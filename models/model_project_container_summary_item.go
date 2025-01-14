/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ProjectContainerSummaryItem struct {

	// The ID of this container
	Id string `json:"Id,omitempty"`

	// The names that this container has been given
	Name string `json:"Name,omitempty"`

	// Command to run when starting the container
	Command string `json:"Command,omitempty"`

	Project string `json:"Project,omitempty"`

	Service string `json:"Service,omitempty"`

	// The state of this container (e.g. `Exited`)
	State string `json:"State,omitempty"`

	Health string `json:"Health,omitempty"`

	ExitCode int `json:"ExitCode,omitempty"`

	Publishers []PortPublisher `json:"Publishers,omitempty"`
}
