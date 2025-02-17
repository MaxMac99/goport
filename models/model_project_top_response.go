/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ProjectTopResponse - OK response to ProjectTop operation
type ProjectTopResponse struct {

	// The ids
	Id string `json:"ID,omitempty"`

	Name string `json:"Name,omitempty"`

	// The ps column titles
	Titles []string `json:"Titles,omitempty"`

	// Each process running in the project, where each is process is an array of values corresponding to the titles.
	Processes [][]string `json:"Processes,omitempty"`
}
