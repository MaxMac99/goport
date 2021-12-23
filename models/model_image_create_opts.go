/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ImageCreateOpts struct {
	Context      string  `form:"context"`
	FromImage    string  `form:"fromImage"`
	FromSrc      string  `form:"fromSrc"`
	Repo         string  `form:"repo"`
	Tag          *string `form:"tag"`
	Message      string  `form:"message"`
	Platform     string  `form:"platform"`
	Quiet        bool    `form:"quiet,default=true"`
	RegistryAuth string  `header:"X-Registry-Auth"`
}