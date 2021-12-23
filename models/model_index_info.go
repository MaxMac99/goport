/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// IndexInfo - IndexInfo contains information about a registry.
type IndexInfo struct {

	// Name of the registry, such as \"docker.io\".
	Name string `json:"Name,omitempty"`

	// List of mirrors, expressed as URIs.
	Mirrors []string `json:"Mirrors,omitempty"`

	// Indicates if the registry is part of the list of insecure registries.  If `false`, the registry is insecure. Insecure registries accept un-encrypted (HTTP) and/or untrusted (HTTPS with certificates from unknown CAs) communication.  > **Warning**: Insecure registries can be useful when running a local > registry. However, because its use creates security vulnerabilities > it should ONLY be enabled for testing purposes. For increased > security, users should add their CA to their system's list of > trusted CAs instead of enabling this option.
	Secure bool `json:"Secure,omitempty"`

	// Indicates whether this is an official registry (i.e., Docker Hub / docker.io)
	Official bool `json:"Official,omitempty"`
}
