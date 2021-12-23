/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ProgressDetail struct {
	Current int32 `json:"current,omitempty"`

	Total int32 `json:"total,omitempty"`
}
