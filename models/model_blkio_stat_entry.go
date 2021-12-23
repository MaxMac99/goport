/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// BlkioStatEntry - BlkioStatEntry is one small entity to store a piece of Blkio stats. Not used on Windows.
type BlkioStatEntry struct {
	Major uint64 `json:"Major,omitempty"`

	Minor uint64 `json:"Minor,omitempty"`

	Op string `json:"Op,omitempty"`

	Value uint64 `json:"Value,omitempty"`
}