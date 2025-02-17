/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ContainerState - ContainerState stores container's running state. It's part of ContainerJSONBase and will be returned by the \"inspect\" command.
type ContainerState struct {

	// String representation of the container state. Can be one of \"created\", \"running\", \"paused\", \"restarting\", \"removing\", \"exited\", or \"dead\".
	Status string `json:"Status,omitempty"`

	// Whether this container is running.  Note that a running container can be _paused_. The `Running` and `Paused` booleans are not mutually exclusive:  When pausing a container (on Linux), the freezer cgroup is used to suspend all processes in the container. Freezing the process requires the process to be running. As a result, paused containers are both `Running` _and_ `Paused`.  Use the `Status` field instead to determine if a container's state is \"running\".
	Running bool `json:"Running,omitempty"`

	// Whether this container is paused.
	Paused bool `json:"Paused,omitempty"`

	// Whether this container is restarting.
	Restarting bool `json:"Restarting,omitempty"`

	// Whether this container has been killed because it ran out of memory.
	OOMKilled bool `json:"OOMKilled,omitempty"`

	Dead bool `json:"Dead,omitempty"`

	// The process ID of this container
	Pid int `json:"Pid,omitempty"`

	// The last exit code of this container
	ExitCode int `json:"ExitCode,omitempty"`

	Error string `json:"Error,omitempty"`

	// The time when this container was last started.
	StartedAt string `json:"StartedAt,omitempty"`

	// The time when this container last exited.
	FinishedAt string `json:"FinishedAt,omitempty"`

	Health *Health `json:"Health,omitempty"`
}
