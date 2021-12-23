/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

var routes = Routes{
	{
		"ContainerChanges",
		http.MethodGet,
		"/v1.0/containers/:id/changes",
		ContainerChangesHandler,
	},

	{
		"ContainerCreate",
		http.MethodPost,
		"/v1.0/containers/create",
		ContainerCreateHandler,
	},

	{
		"ContainerDelete",
		http.MethodDelete,
		"/v1.0/containers/:id",
		ContainerDeleteHandler,
	},

	{
		"ContainerExport",
		http.MethodGet,
		"/v1.0/containers/:id/export",
		ContainerExportHandler,
	},

	{
		"ContainerInspect",
		http.MethodGet,
		"/v1.0/containers/:id/json",
		ContainerInspectHandler,
	},

	{
		"ContainerKill",
		http.MethodPost,
		"/v1.0/containers/:id/kill",
		ContainerKillHandler,
	},

	{
		"ContainerList",
		http.MethodGet,
		"/v1.0/containers/json",
		ContainerListHandler,
	},

	{
		"ContainerLogs",
		http.MethodGet,
		"/v1.0/containers/:id/logs",
		ContainerLogsHandler,
	},

	{
		"ContainerPause",
		http.MethodPost,
		"/v1.0/containers/:id/pause",
		ContainerPauseHandler,
	},

	{
		"ContainerPrune",
		http.MethodPost,
		"/v1.0/containers/prune",
		ContainerPruneHandler,
	},

	{
		"ContainerRename",
		http.MethodPost,
		"/v1.0/containers/:id/rename",
		ContainerRenameHandler,
	},

	{
		"ContainerResize",
		http.MethodPost,
		"/v1.0/containers/:id/resize",
		ContainerResizeHandler,
	},

	{
		"ContainerRestart",
		http.MethodPost,
		"/v1.0/containers/:id/restart",
		ContainerRestartHandler,
	},

	{
		"ContainerStart",
		http.MethodPost,
		"/v1.0/containers/:id/start",
		ContainerStartHandler,
	},

	{
		"ContainerStats",
		http.MethodGet,
		"/v1.0/containers/:id/stats",
		ContainerStatsHandler,
	},

	{
		"ContainerStop",
		http.MethodPost,
		"/v1.0/containers/:id/stop",
		ContainerStopHandler,
	},

	{
		"ContainerTop",
		http.MethodGet,
		"/v1.0/containers/:id/top",
		ContainerTopHandler,
	},

	{
		"ContainerUnpause",
		http.MethodPost,
		"/v1.0/containers/:id/unpause",
		ContainerUnpauseHandler,
	},

	{
		"ContainerUpdate",
		http.MethodPost,
		"/v1.0/containers/:id/update",
		ContainerUpdateHandler,
	},

	{
		"ContainerWait",
		http.MethodPost,
		"/v1.0/containers/:id/wait",
		ContainerWaitHandler,
	},

	{
		"ContextCreate",
		http.MethodPost,
		"/v1.0/contexts/create",
		ContextCreateHandler,
	},

	{
		"ContextDelete",
		http.MethodDelete,
		"/v1.0/contexts/:name",
		ContextDeleteHandler,
	},

	{
		"ContextInspect",
		http.MethodGet,
		"/v1.0/contexts/:name/json",
		ContextInspectHandler,
	},

	{
		"ContextList",
		http.MethodGet,
		"/v1.0/contexts/json",
		ContextListHandler,
	},

	{
		"ContextUpdate",
		http.MethodPost,
		"/v1.0/contexts/:name/update",
		ContextUpdateHandler,
	},

	{
		"ContainerExec",
		http.MethodPost,
		"/v1.0/containers/:id/exec",
		ContainerExecHandler,
	},

	{
		"ExecInspect",
		http.MethodGet,
		"/v1.0/exec/:id/json",
		ExecInspectHandler,
	},

	{
		"ExecResize",
		http.MethodPost,
		"/v1.0/exec/:id/resize",
		ExecResizeHandler,
	},

	{
		"ExecStart",
		http.MethodPost,
		"/v1.0/exec/:id/start",
		ExecStartHandler,
	},

	{
		"BuildPrune",
		http.MethodPost,
		"/v1.0/build/prune",
		BuildPruneHandler,
	},

	{
		"ImageBuild",
		http.MethodPost,
		"/v1.0/build",
		ImageBuildHandler,
	},

	{
		"ImageCommit",
		http.MethodPost,
		"/v1.0/commit",
		ImageCommitHandler,
	},

	{
		"ImageCreate",
		http.MethodPost,
		"/v1.0/images/create",
		ImageCreateHandler,
	},

	{
		"ImageDelete",
		http.MethodDelete,
		"/v1.0/images/:name",
		ImageDeleteHandler,
	},

	{
		"ImageHistory",
		http.MethodGet,
		"/v1.0/images/:name/history",
		ImageHistoryHandler,
	},

	{
		"ImageInspect",
		http.MethodGet,
		"/v1.0/images/:name/json",
		ImageInspectHandler,
	},

	{
		"ImageList",
		http.MethodGet,
		"/v1.0/images/json",
		ImageListHandler,
	},

	{
		"ImageLoad",
		http.MethodPost,
		"/v1.0/images/load",
		ImageLoadHandler,
	},

	{
		"ImagePrune",
		http.MethodPost,
		"/v1.0/images/prune",
		ImagePruneHandler,
	},

	{
		"ImagePush",
		http.MethodPost,
		"/v1.0/images/:name/push",
		ImagePushHandler,
	},

	{
		"ImageSearch",
		http.MethodGet,
		"/v1.0/images/search",
		ImageSearchHandler,
	},

	{
		"ImageTag",
		http.MethodPost,
		"/v1.0/images/:name/tag",
		ImageTagHandler,
	},

	{
		"NetworkConnect",
		http.MethodPost,
		"/v1.0/networks/:id/connect",
		NetworkConnectHandler,
	},

	{
		"NetworkCreate",
		http.MethodPost,
		"/v1.0/networks/create",
		NetworkCreateHandler,
	},

	{
		"NetworkDelete",
		http.MethodDelete,
		"/v1.0/networks/:id",
		NetworkDeleteHandler,
	},

	{
		"NetworkDisconnect",
		http.MethodPost,
		"/v1.0/networks/:id/disconnect",
		NetworkDisconnectHandler,
	},

	{
		"NetworkInspect",
		http.MethodGet,
		"/v1.0/networks/:id",
		NetworkInspectHandler,
	},

	{
		"NetworkList",
		http.MethodGet,
		"/v1.0/networks",
		NetworkListHandler,
	},

	{
		"NetworkPrune",
		http.MethodPost,
		"/v1.0/networks/prune",
		NetworkPruneHandler,
	},

	{
		"ProjectBuild",
		http.MethodPost,
		"/v1.0/projects/:name/build",
		ProjectBuildHandler,
	},

	{
		"ProjectCreate",
		http.MethodPost,
		"/v1.0/projects/:name",
		ProjectCreateHandler,
	},

	{
		"ProjectDown",
		http.MethodPost,
		"/v1.0/projects/:name/down",
		ProjectDownHandler,
	},

	{
		"ProjectEvents",
		http.MethodGet,
		"/v1.0/projects/:name/events",
		ProjectEventsHandler,
	},

	{
		"ProjectImages",
		http.MethodGet,
		"/v1.0/projects/:name/images",
		ProjectImagesHandler,
	},

	{
		"ProjectInspect",
		http.MethodGet,
		"/v1.0/projects/:name",
		ProjectInspectHandler,
	},

	{
		"ProjectKill",
		http.MethodPost,
		"/v1.0/projects/:name/kill",
		ProjectKillHandler,
	},

	{
		"ProjectList",
		http.MethodGet,
		"/v1.0/projects/json",
		ProjectListHandler,
	},

	{
		"ProjectLogs",
		http.MethodGet,
		"/v1.0/projects/:name/logs",
		ProjectLogsHandler,
	},

	{
		"ProjectPause",
		http.MethodPost,
		"/v1.0/projects/:name/pause",
		ProjectPauseHandler,
	},

	{
		"ProjectPs",
		http.MethodPost,
		"/v1.0/projects/:name/ps",
		ProjectPsHandler,
	},

	{
		"ProjectPull",
		http.MethodPost,
		"/v1.0/projects/:name/pull",
		ProjectPullHandler,
	},

	{
		"ProjectPush",
		http.MethodPost,
		"/v1.0/projects/:name/push",
		ProjectPushHandler,
	},

	{
		"ProjectRemove",
		http.MethodPost,
		"/v1.0/projects/:name/rm",
		ProjectRemoveHandler,
	},

	{
		"ProjectRestart",
		http.MethodPost,
		"/v1.0/projects/:name/restart",
		ProjectRestartHandler,
	},

	{
		"ProjectRun",
		http.MethodPost,
		"/v1.0/projects/:name/run/:service",
		ProjectRunHandler,
	},

	{
		"ProjectStart",
		http.MethodPost,
		"/v1.0/projects/:name/start",
		ProjectStartHandler,
	},

	{
		"ProjectStop",
		http.MethodPost,
		"/v1.0/projects/:name/stop",
		ProjectStopHandler,
	},

	{
		"ProjectTop",
		http.MethodGet,
		"/v1.0/projects/:name/top",
		ProjectTopHandler,
	},

	{
		"ProjectUnpause",
		http.MethodPost,
		"/v1.0/projects/:name/unpause",
		ProjectUnpauseHandler,
	},

	{
		"ProjectUp",
		http.MethodPost,
		"/v1.0/projects/:name/up",
		ProjectUpHandler,
	},

	{
		"SystemDataUsage",
		http.MethodGet,
		"/v1.0/system/df",
		SystemDataUsageHandler,
	},

	{
		"SystemEvents",
		http.MethodGet,
		"/v1.0/events",
		SystemEventsHandler,
	},

	{
		"SystemInfo",
		http.MethodGet,
		"/v1.0/info",
		SystemInfoHandler,
	},

	{
		"SystemPing",
		http.MethodGet,
		"/v1.0/_ping",
		SystemPingHandler,
	},

	{
		"SystemPingHead",
		http.MethodHead,
		"/v1.0/_ping",
		SystemPingHeadHandler,
	},

	{
		"SystemVersion",
		http.MethodGet,
		"/v1.0/version",
		SystemVersionHandler,
	},

	{
		"VolumeCreate",
		http.MethodPost,
		"/v1.0/volumes/create",
		VolumeCreateHandler,
	},

	{
		"VolumeDelete",
		http.MethodDelete,
		"/v1.0/volumes/:name",
		VolumeDeleteHandler,
	},

	{
		"VolumeInspect",
		http.MethodGet,
		"/v1.0/volumes/:name",
		VolumeInspectHandler,
	},

	{
		"VolumeList",
		http.MethodGet,
		"/v1.0/volumes",
		VolumeListHandler,
	},

	{
		"VolumePrune",
		http.MethodPost,
		"/v1.0/volumes/prune",
		VolumePruneHandler,
	},
}