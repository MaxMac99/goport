/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Health - Health stores information about the container's healthcheck results.
type Health struct {
	Status HealthStatus `json:"Status,omitempty"`

	// FailingStreak is the number of consecutive failures
	FailingStreak int `json:"FailingStreak,omitempty"`

	// Log contains the last few results (oldest first)
	Log []HealthcheckResult `json:"Log,omitempty"`
}
