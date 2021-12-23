/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ContainerConfig - Configuration for a container that is portable between hosts
type ContainerConfig struct {

	// The hostname to use for the container, as a valid RFC 1123 hostname.
	Hostname string `json:"Hostname,omitempty"`

	// The domain name to use for the container.
	Domainname string `json:"Domainname,omitempty"`

	// The user that commands are run as inside the container.
	User string `json:"User,omitempty"`

	// Whether to attach to `stdin`.
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Whether to attach to `stdout`.
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Whether to attach to `stderr`.
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// An object mapping ports to an empty object in the form:  `{\"<port>/<tcp|udp|sctp>\": {}}`
	ExposedPorts map[string]map[string]interface{} `json:"ExposedPorts,omitempty"`

	// Attach standard streams to a TTY, including `stdin` if it is not closed.
	Tty bool `json:"Tty,omitempty"`

	// Open `stdin`
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// Close `stdin` after one attached client disconnects
	StdinOnce bool `json:"StdinOnce,omitempty"`

	// A list of environment variables to set inside the container in the form `[\"VAR=value\", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value.
	Env []string `json:"Env,omitempty"`

	// Command to run specified as a string or an array of strings.
	Cmd []string `json:"Cmd,omitempty"`

	Healthcheck *HealthConfig `json:"Healthcheck,omitempty"`

	// Command is already escaped (Windows only)
	ArgsEscaped bool `json:"ArgsEscaped,omitempty"`

	// The name of the image to use when creating the container/
	Image string `json:"Image,omitempty"`

	// An object mapping mount point paths inside the container to empty objects.
	Volumes map[string]map[string]interface{} `json:"Volumes,omitempty"`

	// The working directory for commands to run in.
	WorkingDir string `json:"WorkingDir,omitempty"`

	// The entry point for the container as a string or an array of strings.  If the array consists of exactly one empty string (`[\"\"]`) then the entry point is reset to system default (i.e., the entry point used by docker when there is no `ENTRYPOINT` instruction in the `Dockerfile`).
	Entrypoint []string `json:"Entrypoint,omitempty"`

	// Disable networking for the container.
	NetworkDisabled bool `json:"NetworkDisabled,omitempty"`

	// MAC address of the container.
	MacAddress string `json:"MacAddress,omitempty"`

	// `ONBUILD` metadata that were defined in the image's `Dockerfile`.
	OnBuild []string `json:"OnBuild,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// Signal to stop a container as a string or unsigned integer.
	StopSignal string `json:"StopSignal,omitempty"`

	// Timeout to stop a container in seconds.
	StopTimeout *int `json:"StopTimeout,omitempty"`

	// Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell.
	Shell []string `json:"Shell,omitempty"`
}