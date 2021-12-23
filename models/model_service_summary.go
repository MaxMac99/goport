/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ServiceSummary - Configuration of the docker-compose service.
type ServiceSummary struct {

	// The name of the service.
	Name string `json:"Name,omitempty"`

	// String representation of the service state.  Can be on of:   * \"running\": All of the containers are running.   * \"created\": One of the containers is created.   * \"paused\": One of the containers is paused.   * \"restarting\": One of the containers is restarting.   * \"dead\": One of the containers is dead.   * \"exited\": One of the containers is exited.   * \"removing\": One of the containers is removing.   * \"none\": None of the containers has been created or is in any other state.
	State string `json:"State,omitempty"`

	Health HealthStatus `json:"Health,omitempty"`

	// A list of all container-ids of this service.
	Containers []string `json:"Containers,omitempty"`

	// The command for the service.
	Command string `json:"Command,omitempty"`

	// The restart policy of the service.
	Restart string `json:"Restart,omitempty"`

	// PortMap describes the mapping of container ports to host ports, using the container's port-number and protocol as key in the format `<port>/<protocol>`, for example, `80/udp`.  If a container's port is mapped for multiple protocols, separate entries are added to the mapping table.
	Ports map[string][]PortBinding `json:"Ports,omitempty"`
}