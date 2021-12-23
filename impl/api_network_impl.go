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
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/maxmac99/goport/controllers"
	"github.com/maxmac99/goport/models"
)

// NetworkConnect - Connect a container to a network
func NetworkConnect(c *gin.Context, opts *models.NetworkConnectOpts) error {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	var ipamConfig *network.EndpointIPAMConfig
	if opts.EndpointConfig.IPAMConfig != nil {
		ipamConfig = &network.EndpointIPAMConfig{
			IPv4Address:  opts.EndpointConfig.IPAMConfig.IPv4Address,
			IPv6Address:  opts.EndpointConfig.IPAMConfig.IPv6Address,
			LinkLocalIPs: opts.EndpointConfig.IPAMConfig.LinkLocalIPs,
		}
	}
	options := network.EndpointSettings{
		IPAMConfig:          ipamConfig,
		Links:               opts.EndpointConfig.Links,
		Aliases:             opts.EndpointConfig.Aliases,
		NetworkID:           opts.EndpointConfig.NetworkID,
		EndpointID:          opts.EndpointConfig.EndpointID,
		Gateway:             opts.EndpointConfig.Gateway,
		IPAddress:           opts.EndpointConfig.IPAddress,
		IPPrefixLen:         opts.EndpointConfig.IPPrefixLen,
		IPv6Gateway:         opts.EndpointConfig.IPv6Gateway,
		GlobalIPv6Address:   opts.EndpointConfig.GlobalIPv6Address,
		GlobalIPv6PrefixLen: opts.EndpointConfig.GlobalIPv6PrefixLen,
		MacAddress:          opts.EndpointConfig.MacAddress,
		DriverOpts:          opts.EndpointConfig.DriverOpts,
	}
	return client.NetworkConnect(c, opts.Id, opts.Container, &options)
}

// NetworkCreate - Create a network
func NetworkCreate(c *gin.Context, opts *models.NetworkCreateOpts) (*models.NetworkCreateResponse, error) {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	var ipam *network.IPAM
	if opts.IPAM != nil {
		var config []network.IPAMConfig
		for _, item := range opts.IPAM.Config {
			config = append(config, network.IPAMConfig{
				Subnet:     item.Subnet,
				IPRange:    item.IPRange,
				Gateway:    item.Gateway,
				AuxAddress: item.AuxAddress,
			})
		}
		ipam = &network.IPAM{
			Driver:  opts.IPAM.Driver,
			Options: opts.IPAM.Options,
			Config:  config,
		}
	}
	options := types.NetworkCreate{
		CheckDuplicate: opts.CheckDuplicate,
		Driver:         opts.Driver,
		Scope:          "",
		EnableIPv6:     opts.EnableIPv6,
		IPAM:           ipam,
		Internal:       opts.Internal,
		Attachable:     opts.Attachable,
		Ingress:        opts.Ingress,
		ConfigOnly:     false,
		ConfigFrom:     nil,
		Options:        opts.Options,
		Labels:         opts.Labels,
	}
	response, err := client.NetworkCreate(c, opts.Name, options)
	if err != nil {
		return nil, err
	}
	return &models.NetworkCreateResponse{
		Id:      response.ID,
		Warning: response.Warning,
	}, nil
}

// NetworkDelete - Remove a network
func NetworkDelete(c *gin.Context, opts *models.NetworkDeleteOpts) error {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	return client.NetworkRemove(c, opts.Id)
}

// NetworkDisconnect - Disconnect a container from a network
func NetworkDisconnect(c *gin.Context, opts *models.NetworkDisconnectOpts) error {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	return client.NetworkDisconnect(c, opts.Id, opts.Container, opts.Force)
}

// NetworkInspect - Inspect a network
func NetworkInspect(c *gin.Context, opts *models.NetworkInspectOpts) (*types.NetworkResource, error) {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	options := types.NetworkInspectOptions{
		Scope:   opts.Scope,
		Verbose: opts.Verbose,
	}
	response, err := client.NetworkInspect(c, opts.Id, options)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// NetworkList - List networks
func NetworkList(c *gin.Context, opts *models.NetworkListOpts) (*map[string][]models.Network, error) {
	clients, err := controllers.ResolveContexts(opts.Context)
	if err != nil {
		return nil, err
	}
	parsedFilters, err := filters.FromJSON(opts.Filters)
	if err != nil {
		return nil, err
	}
	options := types.NetworkListOptions{
		Filters: parsedFilters,
	}
	output := make(map[string][]models.Network, len(clients))
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for context, cli := range clients {
		go func(context string, cli client.APIClient) {
			list, err := cli.NetworkList(c, options)
			if err != nil {
				wg.Done()
				return
			}
			var contextItems []models.Network
			for _, item := range list {
				contextItems = append(contextItems, models.MapToNetwork(item))
			}
			mutex.Lock()
			output[context] = contextItems
			mutex.Unlock()
			wg.Done()
		}(context, cli)
	}
	wg.Wait()
	return &output, nil
}

// NetworkPrune - Delete unused networks
func NetworkPrune(c *gin.Context, opts *models.NetworkPruneOpts) (*models.NetworkPruneResponse, error) {
	clients, err := controllers.ResolveContexts(opts.Context)
	if err != nil {
		return nil, err
	}
	parsedFilters, err := filters.FromJSON(opts.Filters)
	if err != nil {
		return nil, err
	}
	var networksDeleted []string
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for context, cli := range clients {
		go func(context string, cli client.APIClient) {
			response, err := cli.NetworksPrune(c, parsedFilters)
			if err != nil {
				wg.Done()
				return
			}
			mutex.Lock()
			networksDeleted = append(networksDeleted, response.NetworksDeleted...)
			mutex.Unlock()
			wg.Done()
		}(context, cli)
	}
	wg.Wait()
	return &models.NetworkPruneResponse{
		NetworksDeleted: networksDeleted,
	}, nil
}
