/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ImageSearchResponseItem struct {
	Description string `json:"description,omitempty"`

	IsOfficial bool `json:"is_official,omitempty"`

	IsAutomated bool `json:"is_automated,omitempty"`

	Name string `json:"name,omitempty"`

	StarCount int `json:"star_count,omitempty"`
}
