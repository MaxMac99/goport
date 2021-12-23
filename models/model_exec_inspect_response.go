/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ExecInspectResponse struct {
	CanRemove bool `json:"CanRemove,omitempty"`

	DetachKeys string `json:"DetachKeys,omitempty"`

	ID string `json:"ID,omitempty"`

	Running bool `json:"Running,omitempty"`

	ExitCode int `json:"ExitCode,omitempty"`

	ProcessConfig ProcessConfig `json:"ProcessConfig,omitempty"`

	OpenStdin bool `json:"OpenStdin,omitempty"`

	OpenStderr bool `json:"OpenStderr,omitempty"`

	OpenStdout bool `json:"OpenStdout,omitempty"`

	ContainerID string `json:"ContainerID,omitempty"`

	// The system process ID for the exec process.
	Pid int `json:"Pid,omitempty"`
}