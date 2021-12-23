/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ContainerInspectResponse struct {

	// The ID of the container
	Id string `json:"Id,omitempty"`

	// The time the container was created
	Created string `json:"Created,omitempty"`

	// The path to the command being run
	Path string `json:"Path,omitempty"`

	// The arguments to the command being run
	Args []string `json:"Args,omitempty"`

	State *ContainerState `json:"State,omitempty"`

	// The container's image ID
	Image string `json:"Image,omitempty"`

	ResolvConfPath string `json:"ResolvConfPath,omitempty"`

	HostnamePath string `json:"HostnamePath,omitempty"`

	HostsPath string `json:"HostsPath,omitempty"`

	LogPath string `json:"LogPath,omitempty"`

	Name string `json:"Name,omitempty"`

	RestartCount int `json:"RestartCount,omitempty"`

	Driver string `json:"Driver,omitempty"`

	Platform string `json:"Platform,omitempty"`

	MountLabel string `json:"MountLabel,omitempty"`

	ProcessLabel string `json:"ProcessLabel,omitempty"`

	AppArmorProfile string `json:"AppArmorProfile,omitempty"`

	// IDs of exec instances that are running in the container.
	ExecIDs []string `json:"ExecIDs,omitempty"`

	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	GraphDriver GraphDriverData `json:"GraphDriver,omitempty"`

	// The size of files that have been created or changed by this container.
	SizeRw *int64 `json:"SizeRw,omitempty"`

	// The total size of all the files in this container.
	SizeRootFs *int64 `json:"SizeRootFs,omitempty"`

	Mounts []MountPoint `json:"Mounts,omitempty"`

	Config *ContainerConfig `json:"Config,omitempty"`

	NetworkSettings *NetworkSettings `json:"NetworkSettings,omitempty"`
}