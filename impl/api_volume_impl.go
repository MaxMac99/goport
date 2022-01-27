/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package impl

import (
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"gitlab.com/maxmac99/goport/context"
	"gitlab.com/maxmac99/goport/models"
)

// VolumeCreate - Create a volume
func VolumeCreate(c *gin.Context, opts *models.VolumeCreateOpts) (*types.Volume, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	options := volume.VolumeCreateBody{
		Driver:     opts.Driver,
		DriverOpts: opts.DriverOpts,
		Labels:     opts.Labels,
		Name:       opts.Name,
	}
	response, err := client.VolumeCreate(c, options)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// VolumeDelete - Remove a volume
func VolumeDelete(c *gin.Context, opts *models.VolumeDeleteOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	return client.VolumeRemove(c, opts.Name, opts.Force)
}

// VolumeInspect - Inspect a volume
func VolumeInspect(c *gin.Context, opts *models.VolumeInspectOpts) (*models.Volume, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	response, err := client.VolumeInspect(c, opts.Name)
	if err != nil {
		return nil, err
	}
	volume := buildVolume(response)
	return &volume, nil
}

// VolumeList - List volumes
func VolumeList(c *gin.Context, opts *models.VolumeListOpts) (*map[string]models.VolumeListResponse, error) {
	clients, err := context.ResolveContexts(opts.Context)
	if err != nil {
		return nil, err
	}
	parsedFilters, err := filters.FromJSON(opts.Filters)
	if err != nil {
		return nil, err
	}
	output := make(map[string]models.VolumeListResponse, len(clients))
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for context, cli := range clients {
		go func(context string, cli client.APIClient) {
			list, err := cli.VolumeList(c, parsedFilters)
			if err != nil {
				wg.Done()
				return
			}
			warnings := make([]string, 0)
			if len(list.Warnings) > 0 {
				warnings = list.Warnings
			}
			volumes := make([]models.Volume, 0)
			if len(list.Volumes) > 0 {
				for _, volume := range list.Volumes {
					volumes = append(volumes, buildVolume(*volume))
				}
			}
			mutex.Lock()
			output[context] = models.VolumeListResponse{
				Volumes:  volumes,
				Warnings: warnings,
			}
			mutex.Unlock()
			wg.Done()
		}(context, cli)
	}
	wg.Wait()
	return &output, nil
}

// VolumePrune - Delete unused volumes
func VolumePrune(c *gin.Context, opts *models.VolumePruneOpts) (*map[string]models.VolumePruneResponse, error) {
	clients, err := context.ResolveContexts(opts.Context)
	if err != nil {
		return nil, err
	}
	parsedFilters, err := filters.FromJSON(opts.Filters)
	if err != nil {
		return nil, err
	}
	response := make(map[string]models.VolumePruneResponse)
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for context, cli := range clients {
		go func(context string, cli client.APIClient) {
			pruneResponse, err := cli.VolumesPrune(c, parsedFilters)
			if err != nil {
				wg.Done()
				return
			}
			mutex.Lock()
			response[context] = models.VolumePruneResponse{
				VolumesDeleted: pruneResponse.VolumesDeleted,
				SpaceReclaimed: pruneResponse.SpaceReclaimed,
			}
			mutex.Unlock()
			wg.Done()
		}(context, cli)
	}
	wg.Wait()
	return &response, nil
}

func buildVolume(volume types.Volume) models.Volume {
	labels := make(map[string]string)
	if len(volume.Labels) > 0 {
		labels = volume.Labels
	}
	options := make(map[string]string)
	if len(volume.Options) > 0 {
		options = volume.Options
	}
	var usageData *models.VolumeUsageData
	if volume.UsageData != nil {
		usageData = &models.VolumeUsageData{
			Size:     volume.UsageData.Size,
			RefCount: volume.UsageData.RefCount,
		}
	}
	return models.Volume{
		Name:       volume.Name,
		Driver:     volume.Driver,
		Mountpoint: volume.Mountpoint,
		CreatedAt:  volume.CreatedAt,
		Status:     volume.Status,
		Labels:     labels,
		Scope:      volume.Scope,
		Options:    options,
		UsageData:  usageData,
	}
}
