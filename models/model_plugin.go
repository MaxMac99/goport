/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Plugin - A plugin for the Engine API
type Plugin struct {
	Id string `json:"Id,omitempty"`

	Name string `json:"Name"`

	// True if the plugin is running. False if the plugin is not running, only installed.
	Enabled bool `json:"Enabled"`

	Settings PluginSettings `json:"Settings"`

	// plugin remote reference used to push/pull the plugin
	PluginReference string `json:"PluginReference,omitempty"`

	Config PluginConfig `json:"Config"`
}