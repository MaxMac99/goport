/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// HostConfig - Container configuration that depends on the host we are running on
type HostConfig struct {

	// An integer value representing this container's relative CPU weight versus other containers.
	CpuShares int64 `json:"CpuShares,omitempty"`

	// Memory limit in bytes.
	Memory int64 `json:"Memory,omitempty"`

	// Path to `cgroups` under which the container's `cgroup` is created. If the path is not absolute, the path is considered to be relative to the `cgroups` path of the init process. Cgroups are created if they do not already exist.
	CgroupParent string `json:"CgroupParent,omitempty"`

	// Block IO weight (relative weight).
	BlkioWeight uint16 `json:"BlkioWeight,omitempty"`

	// Block IO weight (relative device weight) in the form:  ``` [{\"Path\": \"device_path\", \"Weight\": weight}] ```
	BlkioWeightDevice []ResourcesBlkioWeightDevice `json:"BlkioWeightDevice,omitempty"`

	// Limit read rate (bytes per second) from a device, in the form:  ``` [{\"Path\": \"device_path\", \"Rate\": rate}] ```
	BlkioDeviceReadBps []ThrottleDevice `json:"BlkioDeviceReadBps,omitempty"`

	// Limit write rate (bytes per second) to a device, in the form:  ``` [{\"Path\": \"device_path\", \"Rate\": rate}] ```
	BlkioDeviceWriteBps []ThrottleDevice `json:"BlkioDeviceWriteBps,omitempty"`

	// Limit read rate (IO per second) from a device, in the form:  ``` [{\"Path\": \"device_path\", \"Rate\": rate}] ```
	BlkioDeviceReadIOps []ThrottleDevice `json:"BlkioDeviceReadIOps,omitempty"`

	// Limit write rate (IO per second) to a device, in the form:  ``` [{\"Path\": \"device_path\", \"Rate\": rate}] ```
	BlkioDeviceWriteIOps []ThrottleDevice `json:"BlkioDeviceWriteIOps,omitempty"`

	// The length of a CPU period in microseconds.
	CpuPeriod int64 `json:"CpuPeriod,omitempty"`

	// Microseconds of CPU time that the container can get in a CPU period.
	CpuQuota int64 `json:"CpuQuota,omitempty"`

	// The length of a CPU real-time period in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	CpuRealtimePeriod int64 `json:"CpuRealtimePeriod,omitempty"`

	// The length of a CPU real-time runtime in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	CpuRealtimeRuntime int64 `json:"CpuRealtimeRuntime,omitempty"`

	// CPUs in which to allow execution (e.g., `0-3`, `0,1`).
	CpusetCpus string `json:"CpusetCpus,omitempty"`

	// Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
	CpusetMems string `json:"CpusetMems,omitempty"`

	// A list of devices to add to the container.
	Devices []DeviceMapping `json:"Devices,omitempty"`

	// a list of cgroup rules to apply to the container
	DeviceCgroupRules []string `json:"DeviceCgroupRules,omitempty"`

	// A list of requests for devices to be sent to device drivers.
	DeviceRequests []DeviceRequest `json:"DeviceRequests,omitempty"`

	// Kernel memory limit in bytes.  <p><br /></p>  > **Deprecated**: This field is deprecated as the kernel 5.4 deprecated > `kmem.limit_in_bytes`.
	KernelMemory int64 `json:"KernelMemory,omitempty"`

	// Hard limit for kernel TCP buffer memory (in bytes).
	KernelMemoryTCP int64 `json:"KernelMemoryTCP,omitempty"`

	// Memory soft limit in bytes.
	MemoryReservation int64 `json:"MemoryReservation,omitempty"`

	// Total memory limit (memory + swap). Set as `-1` to enable unlimited swap.
	MemorySwap int64 `json:"MemorySwap,omitempty"`

	// Tune a container's memory swappiness behavior. Accepts an integer between 0 and 100.
	MemorySwappiness *int64 `json:"MemorySwappiness,omitempty"`

	// CPU quota in units of 10<sup>-9</sup> CPUs.
	NanoCpus int64 `json:"NanoCpus,omitempty"`

	// Disable OOM Killer for the container.
	OomKillDisable *bool `json:"OomKillDisable,omitempty"`

	// Run an init inside the container that forwards signals and reaps processes. This field is omitted if empty, and the default (as configured on the daemon) is used.
	Init *bool `json:"Init,omitempty"`

	// Tune a container's PIDs limit. Set `0` or `-1` for unlimited, or `null` to not change.
	PidsLimit *int64 `json:"PidsLimit,omitempty"`

	// A list of resource limits to set in the container. For example:  ``` {\"Name\": \"nofile\", \"Soft\": 1024, \"Hard\": 2048} ```
	Ulimits []ResourcesUlimits `json:"Ulimits,omitempty"`

	// The number of usable CPUs (Windows only).  On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	CpuCount int64 `json:"CpuCount,omitempty"`

	// The usable percentage of the available CPUs (Windows only).  On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	CpuPercent int64 `json:"CpuPercent,omitempty"`

	// Maximum IOps for the container system drive (Windows only)
	IOMaximumIOps uint64 `json:"IOMaximumIOps,omitempty"`

	// Maximum IO in bytes per second for the container system drive (Windows only).
	IOMaximumBandwidth uint64 `json:"IOMaximumBandwidth,omitempty"`

	// A list of volume bindings for this container. Each volume binding is a string in one of these forms:  - `host-src:container-dest[:options]` to bind-mount a host path   into the container. Both `host-src`, and `container-dest` must   be an _absolute_ path. - `volume-name:container-dest[:options]` to bind-mount a volume   managed by a volume driver into the container. `container-dest`   must be an _absolute_ path.  `options` is an optional, comma-delimited list of:  - `nocopy` disables automatic copying of data from the container   path to the volume. The `nocopy` flag only applies to named volumes. - `[ro|rw]` mounts a volume read-only or read-write, respectively.   If omitted or set to `rw`, volumes are mounted read-write. - `[z|Z]` applies SELinux labels to allow or deny multiple containers   to read and write to the same volume.     - `z`: a _shared_ content label is applied to the content. This       label indicates that multiple containers can share the volume       content, for both reading and writing.     - `Z`: a _private unshared_ label is applied to the content.       This label indicates that only the current container can use       a private volume. Labeling systems such as SELinux require       proper labels to be placed on volume content that is mounted       into a container. Without a label, the security system can       prevent a container's processes from using the content. By       default, the labels set by the host operating system are not       modified. - `[[r]shared|[r]slave|[r]private]` specifies mount   [propagation behavior](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt).   This only applies to bind-mounted volumes, not internal volumes   or named volumes. Mount propagation requires the source mount   point (the location where the source directory is mounted in the   host operating system) to have the correct propagation properties.   For shared volumes, the source mount point must be set to `shared`.   For slave volumes, the mount must be set to either `shared` or   `slave`.
	Binds []string `json:"Binds,omitempty"`

	// Path to a file where the container ID is written
	ContainerIDFile string `json:"ContainerIDFile,omitempty"`

	LogConfig HostConfigAllOfLogConfig `json:"LogConfig,omitempty"`

	// Network mode to use for this container. Supported standard values are: `bridge`, `host`, `none`, and `container:<name|id>`. Any other value is taken as a custom network's name to which this container should connect to.
	NetworkMode string `json:"NetworkMode,omitempty"`

	// PortMap describes the mapping of container ports to host ports, using the container's port-number and protocol as key in the format `<port>/<protocol>`, for example, `80/udp`.  If a container's port is mapped for multiple protocols, separate entries are added to the mapping table.
	PortBindings map[string][]PortBinding `json:"PortBindings,omitempty"`

	RestartPolicy RestartPolicy `json:"RestartPolicy,omitempty"`

	// Automatically remove the container when the container's process exits. This has no effect if `RestartPolicy` is set.
	AutoRemove bool `json:"AutoRemove,omitempty"`

	// Driver that this container uses to mount volumes.
	VolumeDriver string `json:"VolumeDriver,omitempty"`

	// A list of volumes to inherit from another container, specified in the form `<container name>[:<ro|rw>]`.
	VolumesFrom []string `json:"VolumesFrom,omitempty"`

	// Specification for mounts to be added to the container.
	Mounts []Mount `json:"Mounts,omitempty"`

	// A list of kernel capabilities to add to the container. Conflicts with option 'Capabilities'.
	CapAdd []string `json:"CapAdd,omitempty"`

	// A list of kernel capabilities to drop from the container. Conflicts with option 'Capabilities'.
	CapDrop []string `json:"CapDrop,omitempty"`

	// cgroup namespace mode for the container. Possible values are:  - `\"private\"`: the container runs in its own private cgroup namespace - `\"host\"`: use the host system's cgroup namespace  If not specified, the daemon default is used, which can either be `\"private\"` or `\"host\"`, depending on daemon version, kernel support and configuration.
	CgroupnsMode string `json:"CgroupnsMode,omitempty"`

	// A list of DNS servers for the container to use.
	Dns []string `json:"Dns,omitempty"`

	// A list of DNS options.
	DnsOptions []string `json:"DnsOptions,omitempty"`

	// A list of DNS search domains.
	DnsSearch []string `json:"DnsSearch,omitempty"`

	// A list of hostnames/IP mappings to add to the container's `/etc/hosts` file. Specified in the form `[\"hostname:IP\"]`.
	ExtraHosts []string `json:"ExtraHosts,omitempty"`

	// A list of additional groups that the container process will run as.
	GroupAdd []string `json:"GroupAdd,omitempty"`

	// IPC sharing mode for the container. Possible values are:  - `\"none\"`: own private IPC namespace, with /dev/shm not mounted - `\"private\"`: own private IPC namespace - `\"shareable\"`: own private IPC namespace, with a possibility to share it with other containers - `\"container:<name|id>\"`: join another (shareable) container's IPC namespace - `\"host\"`: use the host system's IPC namespace  If not specified, daemon default is used, which can either be `\"private\"` or `\"shareable\"`, depending on daemon version and configuration.
	IpcMode string `json:"IpcMode,omitempty"`

	// Cgroup to use for the container.
	Cgroup string `json:"Cgroup,omitempty"`

	// A list of links for the container in the form `container_name:alias`.
	Links []string `json:"Links,omitempty"`

	// An integer value containing the score given to the container in order to tune OOM killer preferences.
	OomScoreAdj int `json:"OomScoreAdj,omitempty"`

	// Set the PID (Process) Namespace mode for the container. It can be either:  - `\"container:<name|id>\"`: joins another container's PID namespace - `\"host\"`: use the host's PID namespace inside the container
	PidMode string `json:"PidMode,omitempty"`

	// Gives the container full access to the host.
	Privileged bool `json:"Privileged,omitempty"`

	// Allocates an ephemeral host port for all of a container's exposed ports.  Ports are de-allocated when the container stops and allocated when the container starts. The allocated port might be changed when restarting the container.  The port is selected from the ephemeral port range that depends on the kernel. For example, on Linux the range is defined by `/proc/sys/net/ipv4/ip_local_port_range`.
	PublishAllPorts bool `json:"PublishAllPorts,omitempty"`

	// Mount the container's root filesystem as read only.
	ReadonlyRootfs bool `json:"ReadonlyRootfs,omitempty"`

	// A list of string values to customize labels for MLS systems, such as SELinux.
	SecurityOpt []string `json:"SecurityOpt,omitempty"`

	// Storage driver options for this container, in the form `{\"size\": \"120G\"}`.
	StorageOpt map[string]string `json:"StorageOpt,omitempty"`

	// A map of container directories which should be replaced by tmpfs mounts, and their corresponding mount options. For example:  ``` { \"/run\": \"rw,noexec,nosuid,size=65536k\" } ```
	Tmpfs map[string]string `json:"Tmpfs,omitempty"`

	// UTS namespace to use for the container.
	UTSMode string `json:"UTSMode,omitempty"`

	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode string `json:"UsernsMode,omitempty"`

	// Size of `/dev/shm` in bytes. If omitted, the system uses 64MB.
	ShmSize int64 `json:"ShmSize,omitempty"`

	// A list of kernel parameters (sysctls) to set in the container. For example:  ``` {\"net.ipv4.ip_forward\": \"1\"} ```
	Sysctls map[string]string `json:"Sysctls,omitempty"`

	// Runtime to use with this container.
	Runtime string `json:"Runtime,omitempty"`

	// Initial console size, as an `[height, width]` array. (Windows only)
	ConsoleSize []uint `json:"ConsoleSize,omitempty"`

	// Isolation technology of the container. (Windows only)
	Isolation string `json:"Isolation,omitempty"`

	// The list of paths to be masked inside the container (this overrides the default set of paths).
	MaskedPaths []string `json:"MaskedPaths,omitempty"`

	// The list of paths to be set as read-only inside the container (this overrides the default set of paths).
	ReadonlyPaths []string `json:"ReadonlyPaths,omitempty"`
}
