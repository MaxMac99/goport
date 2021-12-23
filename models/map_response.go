package models

import (
	"encoding/json"
	"io/fs"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/blkiodev"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-units"
)

func MapConfigFromOptions(opts ContainerCreateConfig) containertypes.Config {
	return containertypes.Config{
		Hostname:     opts.Hostname,
		Domainname:   opts.Domainname,
		User:         opts.User,
		AttachStdin:  opts.AttachStdin,
		AttachStdout: opts.AttachStdout,
		AttachStderr: opts.AttachStderr,
		ExposedPorts: MapPortsFromOptions(opts.ExposedPorts),
		Tty:          opts.Tty,
		OpenStdin:    opts.OpenStdin,
		StdinOnce:    opts.StdinOnce,
		Env:          opts.Env,
		Cmd:          opts.Cmd,
		Healthcheck: &containertypes.HealthConfig{
			Test:        opts.Healthcheck.Test,
			Interval:    time.Duration(opts.Healthcheck.Interval),
			Timeout:     time.Duration(opts.Healthcheck.Timeout),
			StartPeriod: time.Duration(opts.Healthcheck.StartPeriod),
			Retries:     int(opts.Healthcheck.Retries),
		},
		ArgsEscaped:     opts.ArgsEscaped,
		Image:           opts.Image,
		Volumes:         MapDoubleMapToSingle(opts.Volumes),
		WorkingDir:      opts.WorkingDir,
		Entrypoint:      opts.Entrypoint,
		NetworkDisabled: opts.NetworkDisabled,
		MacAddress:      opts.MacAddress,
		OnBuild:         opts.OnBuild,
		Labels:          opts.Labels,
		StopSignal:      opts.StopSignal,
		StopTimeout:     opts.StopTimeout,
		Shell:           opts.Shell,
	}
}

func MapPortsFromOptions(opts map[string]map[string]interface{}) map[nat.Port]struct{} {
	var exposedPortsString []string
	for key := range opts {
		exposedPortsString = append(exposedPortsString, key)
	}
	exposedPorts, _, err := nat.ParsePortSpecs(exposedPortsString)
	if err != nil {
		return nil
	}
	return exposedPorts
}

func MapDoubleMapToSingle(opts map[string]map[string]interface{}) map[string]struct{} {
	volumes := make(map[string]struct{})
	for key := range opts {
		volumes[key] = struct{}{}
	}
	return volumes
}

func MapHostConfigFromOptions(opts HostConfig) containertypes.HostConfig {
	return containertypes.HostConfig{
		Binds:           opts.Binds,
		ContainerIDFile: opts.ContainerIDFile,
		LogConfig:       containertypes.LogConfig(opts.LogConfig),
		NetworkMode:     containertypes.NetworkMode(opts.NetworkMode),
		PortBindings:    MapPortBindingsFromOptions(opts.PortBindings),
		RestartPolicy: containertypes.RestartPolicy{
			Name:              opts.RestartPolicy.Name,
			MaximumRetryCount: opts.RestartPolicy.MaximumRetryCount,
		},
		AutoRemove:      opts.AutoRemove,
		VolumeDriver:    opts.VolumeDriver,
		VolumesFrom:     opts.VolumesFrom,
		CapAdd:          opts.CapAdd,
		CapDrop:         opts.CapDrop,
		CgroupnsMode:    containertypes.CgroupnsMode(opts.CgroupnsMode),
		DNS:             opts.Dns,
		DNSOptions:      opts.DnsOptions,
		DNSSearch:       opts.DnsSearch,
		ExtraHosts:      opts.ExtraHosts,
		GroupAdd:        opts.GroupAdd,
		IpcMode:         containertypes.IpcMode(opts.IpcMode),
		Cgroup:          containertypes.CgroupSpec(opts.Cgroup),
		Links:           opts.Links,
		OomScoreAdj:     opts.OomScoreAdj,
		PidMode:         containertypes.PidMode(opts.PidMode),
		Privileged:      opts.Privileged,
		PublishAllPorts: opts.PublishAllPorts,
		ReadonlyRootfs:  opts.ReadonlyRootfs,
		SecurityOpt:     opts.SecurityOpt,
		StorageOpt:      opts.StorageOpt,
		Tmpfs:           opts.Tmpfs,
		UTSMode:         containertypes.UTSMode(opts.UTSMode),
		UsernsMode:      containertypes.UsernsMode(opts.UsernsMode),
		ShmSize:         opts.ShmSize,
		Sysctls:         opts.Sysctls,
		Runtime:         opts.Runtime,
		ConsoleSize:     MapConsoleSizeFromOptions(opts.ConsoleSize),
		Isolation:       containertypes.Isolation(opts.Isolation),
		Resources: containertypes.Resources{
			CPUShares:            opts.CpuShares,
			Memory:               opts.Memory,
			NanoCPUs:             opts.NanoCpus,
			CgroupParent:         opts.CgroupParent,
			BlkioWeight:          opts.BlkioWeight,
			BlkioWeightDevice:    MapBlkioWeightDeviceFromOptions(opts.BlkioWeightDevice),
			BlkioDeviceReadBps:   MapBlkioThrottleDeviceFromOptions(opts.BlkioDeviceReadBps),
			BlkioDeviceWriteBps:  MapBlkioThrottleDeviceFromOptions(opts.BlkioDeviceWriteBps),
			BlkioDeviceReadIOps:  MapBlkioThrottleDeviceFromOptions(opts.BlkioDeviceReadIOps),
			BlkioDeviceWriteIOps: MapBlkioThrottleDeviceFromOptions(opts.BlkioDeviceWriteIOps),
			CPUPeriod:            opts.CpuPeriod,
			CPUQuota:             opts.CpuQuota,
			CPURealtimePeriod:    opts.CpuRealtimePeriod,
			CPURealtimeRuntime:   opts.CpuRealtimeRuntime,
			CpusetCpus:           opts.CpusetCpus,
			CpusetMems:           opts.CpusetMems,
			Devices:              MapDeviceMappingFromOptions(opts.Devices),
			DeviceCgroupRules:    opts.DeviceCgroupRules,
			DeviceRequests:       MapDeviceRequestsFromOptions(opts.DeviceRequests),
			KernelMemory:         opts.KernelMemory,
			KernelMemoryTCP:      opts.KernelMemoryTCP,
			MemoryReservation:    opts.MemoryReservation,
			MemorySwap:           opts.MemorySwap,
			MemorySwappiness:     opts.MemorySwappiness,
			OomKillDisable:       opts.OomKillDisable,
			PidsLimit:            opts.PidsLimit,
			Ulimits:              MapUlimitsFromOptions(opts.Ulimits),
			CPUCount:             opts.CpuCount,
			CPUPercent:           opts.CpuPercent,
			IOMaximumIOps:        opts.IOMaximumIOps,
			IOMaximumBandwidth:   opts.IOMaximumBandwidth,
		},
		Mounts:        MapMountsFromOptions(opts.Mounts),
		MaskedPaths:   opts.MaskedPaths,
		ReadonlyPaths: opts.ReadonlyPaths,
		Init:          opts.Init,
	}
}

func MapConsoleSizeFromOptions(opts []uint) [2]uint {
	var result [2]uint
	for i, item := range opts {
		if i >= 2 {
			return result
		}
		result[i] = item
	}
	return result
}

func MapPortBindingsFromOptions(opts map[string][]PortBinding) nat.PortMap {
	portBindings := make(nat.PortMap)
	for key, values := range opts {
		port, err := nat.ParsePortSpec(key)
		if err != nil {
			continue
		}

		var bindings []nat.PortBinding
		for _, value := range values {
			bindings = append(bindings, nat.PortBinding{
				HostIP:   value.HostIp,
				HostPort: value.HostPort,
			})
		}
		portBindings[port[0].Port] = bindings
	}
	return portBindings
}

func MapBlkioWeightDeviceFromOptions(opts []ResourcesBlkioWeightDevice) []*blkiodev.WeightDevice {
	var blkioWeightDevice []*blkiodev.WeightDevice
	for _, device := range opts {
		weightDevice := blkiodev.WeightDevice{
			Path:   device.Path,
			Weight: device.Weight,
		}
		blkioWeightDevice = append(blkioWeightDevice, &weightDevice)
	}
	return blkioWeightDevice
}

func MapBlkioThrottleDeviceFromOptions(opts []ThrottleDevice) []*blkiodev.ThrottleDevice {
	var blkioDevice []*blkiodev.ThrottleDevice
	for _, device := range opts {
		throttleDevice := blkiodev.ThrottleDevice{
			Path: device.Path,
			Rate: device.Rate,
		}
		blkioDevice = append(blkioDevice, &throttleDevice)
	}
	return blkioDevice
}

func MapDeviceMappingFromOptions(opts []DeviceMapping) []containertypes.DeviceMapping {
	var devices []containertypes.DeviceMapping
	for _, device := range opts {
		devices = append(devices, containertypes.DeviceMapping{
			PathOnHost:        device.PathOnHost,
			PathInContainer:   device.PathInContainer,
			CgroupPermissions: device.CgroupPermissions,
		})
	}
	return devices
}

func MapDeviceRequestsFromOptions(opts []DeviceRequest) []containertypes.DeviceRequest {
	var deviceRequests []containertypes.DeviceRequest
	for _, device := range opts {
		deviceRequests = append(deviceRequests, containertypes.DeviceRequest{
			Driver:       device.Driver,
			Count:        device.Count,
			DeviceIDs:    device.DeviceIDs,
			Capabilities: device.Capabilities,
			Options:      device.Options,
		})
	}
	return deviceRequests
}

func MapUlimitsFromOptions(opts []ResourcesUlimits) []*units.Ulimit {
	var ulimits []*units.Ulimit
	for _, limit := range opts {
		ulimits = append(ulimits, &units.Ulimit{
			Name: limit.Name,
			Hard: limit.Hard,
			Soft: limit.Soft,
		})
	}
	return ulimits
}

func MapMountsFromOptions(opts []Mount) []mount.Mount {
	var mounts []mount.Mount
	for _, item := range opts {
		var bindOptions *mount.BindOptions
		if item.BindOptions != nil {
			bindOptions = &mount.BindOptions{
				Propagation:  mount.Propagation(item.BindOptions.Propagation),
				NonRecursive: item.BindOptions.NonRecursive,
			}
		}
		var volumeOptions *mount.VolumeOptions
		if item.VolumeOptions != nil {
			var driver *mount.Driver
			if item.VolumeOptions.DriverConfig != nil {
				driver = &mount.Driver{
					Name:    item.VolumeOptions.DriverConfig.Name,
					Options: item.VolumeOptions.DriverConfig.Options,
				}
			}
			volumeOptions = &mount.VolumeOptions{
				NoCopy:       item.VolumeOptions.NoCopy,
				Labels:       item.VolumeOptions.Labels,
				DriverConfig: driver,
			}
		}
		var tmpfsOptions *mount.TmpfsOptions
		if item.TmpfsOptions != nil {
			tmpfsOptions = &mount.TmpfsOptions{
				SizeBytes: item.TmpfsOptions.SizeBytes,
				Mode:      fs.FileMode(item.TmpfsOptions.Mode),
			}
		}
		mounts = append(mounts, mount.Mount{
			Type:          mount.Type(item.Type),
			Source:        item.Source,
			Target:        item.Target,
			ReadOnly:      item.ReadOnly,
			Consistency:   mount.Consistency(item.Consistency),
			BindOptions:   bindOptions,
			VolumeOptions: volumeOptions,
			TmpfsOptions:  tmpfsOptions,
		})
	}
	return mounts
}

func MapNetworkingConfigFromOptions(opts NetworkingConfig) network.NetworkingConfig {
	endpointsConfig := make(map[string]*network.EndpointSettings)
	for key, settings := range opts.EndpointsConfig {
		var ipamConfig *network.EndpointIPAMConfig
		if settings.IPAMConfig != nil {
			ipamConfig = &network.EndpointIPAMConfig{
				IPv4Address:  settings.IPAMConfig.IPv4Address,
				IPv6Address:  settings.IPAMConfig.IPv6Address,
				LinkLocalIPs: settings.IPAMConfig.LinkLocalIPs,
			}
		}
		endpointsConfig[key] = &network.EndpointSettings{
			IPAMConfig:          ipamConfig,
			Links:               settings.Links,
			Aliases:             settings.Aliases,
			NetworkID:           settings.NetworkID,
			EndpointID:          settings.EndpointID,
			Gateway:             settings.Gateway,
			IPAddress:           settings.IPAddress,
			IPPrefixLen:         settings.IPPrefixLen,
			IPv6Gateway:         settings.IPv6Gateway,
			GlobalIPv6Address:   settings.GlobalIPv6Address,
			GlobalIPv6PrefixLen: settings.GlobalIPv6PrefixLen,
			MacAddress:          settings.MacAddress,
			DriverOpts:          settings.DriverOpts,
		}
	}
	return network.NetworkingConfig{
		EndpointsConfig: endpointsConfig,
	}
}

func MapContainerUpdateConfigFromOptions(opts ContainerUpdateOpts) containertypes.UpdateConfig {
	return containertypes.UpdateConfig{
		Resources: containertypes.Resources{
			CPUShares:            opts.Update.CpuShares,
			Memory:               opts.Update.Memory,
			NanoCPUs:             opts.Update.NanoCpus,
			CgroupParent:         opts.Update.CgroupParent,
			BlkioWeight:          opts.Update.BlkioWeight,
			BlkioWeightDevice:    MapBlkioWeightDeviceFromOptions(opts.Update.BlkioWeightDevice),
			BlkioDeviceReadBps:   MapBlkioThrottleDeviceFromOptions(opts.Update.BlkioDeviceReadBps),
			BlkioDeviceWriteBps:  MapBlkioThrottleDeviceFromOptions(opts.Update.BlkioDeviceWriteBps),
			BlkioDeviceReadIOps:  MapBlkioThrottleDeviceFromOptions(opts.Update.BlkioDeviceReadIOps),
			BlkioDeviceWriteIOps: MapBlkioThrottleDeviceFromOptions(opts.Update.BlkioDeviceWriteIOps),
			CPUPeriod:            opts.Update.CpuPeriod,
			CPUQuota:             opts.Update.CpuQuota,
			CPURealtimePeriod:    opts.Update.CpuRealtimePeriod,
			CPURealtimeRuntime:   opts.Update.CpuRealtimeRuntime,
			CpusetCpus:           opts.Update.CpusetCpus,
			CpusetMems:           opts.Update.CpusetMems,
			Devices:              MapDeviceMappingFromOptions(opts.Update.Devices),
			DeviceCgroupRules:    opts.Update.DeviceCgroupRules,
			DeviceRequests:       MapDeviceRequestsFromOptions(opts.Update.DeviceRequests),
			KernelMemory:         opts.Update.KernelMemory,
			KernelMemoryTCP:      opts.Update.KernelMemoryTCP,
			MemoryReservation:    opts.Update.MemoryReservation,
			MemorySwap:           opts.Update.MemorySwap,
			MemorySwappiness:     &opts.Update.MemorySwappiness,
			OomKillDisable:       &opts.Update.OomKillDisable,
			PidsLimit:            opts.Update.PidsLimit,
			Ulimits:              MapUlimitsFromOptions(opts.Update.Ulimits),
			CPUCount:             opts.Update.CpuCount,
			CPUPercent:           opts.Update.CpuPercent,
			IOMaximumIOps:        opts.Update.IOMaximumIOps,
			IOMaximumBandwidth:   opts.Update.IOMaximumBandwidth,
		},
		RestartPolicy: containertypes.RestartPolicy{
			Name:              opts.Update.RestartPolicy.Name,
			MaximumRetryCount: opts.Update.RestartPolicy.MaximumRetryCount,
		},
	}
}

func MapToContainerSummary(container types.Container) ContainerSummary {
	var ports []Port
	for _, port := range container.Ports {
		ports = append(ports, Port{
			IP:          port.IP,
			PrivatePort: port.PrivatePort,
			PublicPort:  port.PublicPort,
			Type:        port.Type,
		})
	}
	networks := make(map[string]EndpointSettings)
	for key, value := range container.NetworkSettings.Networks {
		var endpointIPAMConfig *EndpointIpamConfig
		if value.IPAMConfig != nil {
			endpointIPAMConfig = &EndpointIpamConfig{
				IPv4Address:  value.IPAMConfig.IPv4Address,
				IPv6Address:  value.IPAMConfig.IPv6Address,
				LinkLocalIPs: value.IPAMConfig.LinkLocalIPs,
			}
		}
		networks[key] = EndpointSettings{
			IPAMConfig:          endpointIPAMConfig,
			Links:               value.Links,
			Aliases:             value.Aliases,
			NetworkID:           value.NetworkID,
			EndpointID:          value.EndpointID,
			Gateway:             value.Gateway,
			IPAddress:           value.IPAddress,
			IPPrefixLen:         value.IPPrefixLen,
			IPv6Gateway:         value.IPv6Gateway,
			GlobalIPv6Address:   value.GlobalIPv6Address,
			GlobalIPv6PrefixLen: value.GlobalIPv6PrefixLen,
			MacAddress:          value.MacAddress,
			DriverOpts:          value.DriverOpts,
		}
	}
	var mounts []MountPoint
	for _, mount := range container.Mounts {
		mounts = append(mounts, MountPoint{
			Type:        string(mount.Type),
			Name:        mount.Name,
			Source:      mount.Source,
			Destination: mount.Destination,
			Driver:      mount.Driver,
			Mode:        mount.Mode,
			RW:          mount.RW,
			Propagation: string(mount.Propagation),
		})
	}
	return ContainerSummary{
		Id:         container.ID,
		Names:      container.Names,
		Image:      container.Image,
		ImageID:    container.ImageID,
		Command:    container.Command,
		Created:    container.Created,
		Ports:      ports,
		SizeRw:     container.SizeRw,
		SizeRootFs: container.SizeRootFs,
		Labels:     container.Labels,
		State:      container.State,
		Status:     container.Status,
		HostConfig: ContainerSummaryHostConfig{
			NetworkMode: container.HostConfig.NetworkMode,
		},
		NetworkSettings: ContainerSummaryNetworkSettings{
			Networks: networks,
		},
		Mounts: mounts,
	}
}

func MapImageBuildFromOptions(opts ImageBuildOpts) types.ImageBuildOptions {
	buildArgs := make(map[string]*string)
	json.Unmarshal([]byte(opts.Buildargs), &buildArgs)
	labels := make(map[string]string)
	json.Unmarshal([]byte(opts.Labels), &labels)
	var outputs []types.ImageBuildOutput
	json.Unmarshal([]byte(opts.Outputs), &outputs)
	return types.ImageBuildOptions{
		Tags:           opts.Tags,
		SuppressOutput: opts.Quiet,
		RemoteContext:  opts.Remote,
		NoCache:        opts.Nocache,
		Remove:         opts.Remove,
		ForceRemove:    opts.ForceRemove,
		PullParent:     opts.Pull,
		Isolation:      "",
		CPUSetCPUs:     opts.Cpusetcpus,
		CPUSetMems:     "",
		CPUShares:      opts.Cpushares,
		CPUQuota:       opts.Cpuquota,
		CPUPeriod:      opts.Cpuperiod,
		Memory:         opts.Memory,
		MemorySwap:     opts.Memswap,
		CgroupParent:   "",
		NetworkMode:    opts.Networkmode,
		ShmSize:        opts.Shmsize,
		Dockerfile:     opts.Dockerfile,
		Ulimits:        nil,
		BuildArgs:      buildArgs,
		AuthConfigs:    nil,
		Context:        nil,
		Labels:         labels,
		Squash:         opts.Squash,
		CacheFrom:      opts.Cachefrom,
		SecurityOpt:    nil,
		ExtraHosts:     opts.Extrahosts,
		Target:         opts.Target,
		SessionID:      "",
		Platform:       opts.Platform,
		Version:        "",
		BuildID:        "",
		Outputs:        outputs,
	}
}

func MapImageCommitFromOptions(opts ImageCommitOpts) types.ContainerCommitOptions {
	exposedPorts := make(nat.PortSet)
	for key := range opts.ExposedPorts {
		exposedPorts[nat.Port(key)] = struct{}{}
	}
	var healthcheck *containertypes.HealthConfig
	if opts.Healthcheck != nil {
		healthcheck = &containertypes.HealthConfig{
			Test:        opts.Healthcheck.Test,
			Interval:    time.Duration(opts.Healthcheck.Interval),
			Timeout:     time.Duration(opts.Healthcheck.Timeout),
			StartPeriod: time.Duration(opts.Healthcheck.StartPeriod),
			Retries:     opts.Healthcheck.Retries,
		}
	}
	volumes := make(map[string]struct{})
	for key := range opts.Volumes {
		volumes[key] = struct{}{}
	}
	return types.ContainerCommitOptions{
		Reference: "",
		Comment:   opts.Comment,
		Author:    opts.Author,
		Changes:   opts.Changes,
		Pause:     opts.Pause,
		Config: &containertypes.Config{
			Hostname:        opts.Hostname,
			Domainname:      opts.Domainname,
			User:            opts.User,
			AttachStdin:     opts.AttachStdin,
			AttachStdout:    opts.AttachStdout,
			AttachStderr:    opts.AttachStderr,
			ExposedPorts:    exposedPorts,
			Tty:             opts.Tty,
			OpenStdin:       opts.OpenStdin,
			StdinOnce:       opts.StdinOnce,
			Env:             opts.Env,
			Cmd:             opts.Cmd,
			Healthcheck:     healthcheck,
			ArgsEscaped:     opts.ArgsEscaped,
			Image:           opts.Image,
			Volumes:         volumes,
			WorkingDir:      opts.WorkingDir,
			Entrypoint:      opts.Entrypoint,
			NetworkDisabled: opts.NetworkDisabled,
			MacAddress:      opts.MacAddress,
			OnBuild:         opts.OnBuild,
			Labels:          opts.Labels,
			StopSignal:      opts.StopSignal,
			StopTimeout:     opts.StopTimeout,
			Shell:           opts.Shell,
		},
	}
}

func MapToImageSummary(image types.ImageSummary) ImageSummary {
	return ImageSummary{
		Id:          image.ID,
		ParentId:    image.ParentID,
		RepoTags:    image.RepoTags,
		RepoDigests: image.RepoDigests,
		Created:     image.Created,
		Size:        image.Size,
		SharedSize:  image.SharedSize,
		VirtualSize: image.VirtualSize,
		Labels:      image.Labels,
		Containers:  image.Containers,
	}
}

func MapToNetwork(network types.NetworkResource) Network {
	var ipamConfig []IpamConfig
	for _, item := range network.IPAM.Config {
		ipamConfig = append(ipamConfig, IpamConfig{
			Subnet:     item.Subnet,
			IPRange:    item.IPRange,
			Gateway:    item.Gateway,
			AuxAddress: item.AuxAddress,
		})
	}
	containers := make(map[string]NetworkContainer)
	for key, value := range network.Containers {
		containers[key] = NetworkContainer{
			Name:        value.Name,
			EndpointID:  value.EndpointID,
			MacAddress:  value.MacAddress,
			IPv4Address: value.IPv4Address,
			IPv6Address: value.IPv6Address,
		}
	}
	return Network{
		Name:       network.Name,
		Id:         network.ID,
		Created:    network.Created.String(),
		Scope:      network.Scope,
		Driver:     network.Driver,
		EnableIPv6: network.EnableIPv6,
		IPAM: Ipam{
			Driver:  network.IPAM.Driver,
			Config:  ipamConfig,
			Options: network.IPAM.Options,
		},
		Internal:   network.Internal,
		Attachable: network.Attachable,
		Ingress:    network.Ingress,
		Containers: containers,
		Options:    network.Options,
		Labels:     network.Labels,
	}
}
