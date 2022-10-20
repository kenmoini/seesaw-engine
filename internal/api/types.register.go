package api

//==============================================================================
// GET /api/v1/register/{uuid} - Registration Checks
//==============================================================================

// RegistrationCheckRequest is the structure that houses the request for a registration check
type RegistrationCheckRequest struct {
	// Status is the status of the agent registration check, if found returns 'approved' or 'pending'
	Status string `json:"status,omitempty"`
	// Message is the message returned from the agent registration check
	Message string `json:"message,omitempty"`
}

//==============================================================================
// POST /api/v1/register - Registration of new Agent/Sensor
//==============================================================================

// Register is the structure that comprises the data expected from an agent when making a registration request
type Register struct {
	// Hostname is the hostname of the agent
	Hostname string `json:"hostname,omitempty"`
	// HostInventory is the inventory of the agent
	HostInventory HostInventory `json:"host_inventory,omitempty"`
	// Token is the token used to authenticate the agent
	Token string `json:"token,omitempty"`
}

// RegisterResponse is the structure that comprises the data returned from a registration POST request
type RegisterResponse struct {
	// Status is the status of the registration request
	Status string `json:"status,omitempty"`
	// Message is the message returned from the registration request
	Message string `json:"message,omitempty"`
	// AgentUUID is the agent UUID returned from the registration request
	AgentUUID string `json:"agent_uuid,omitempty"`
}

// HostInventory is the structure that houses the inventory of the agent
type HostInventory struct {
	// Network is the network inventory of the agent
	Network HostNetwork `json:"network,omitempty"`
	// CPU is the CPU inventory of the agent
	CPU HostCPU `json:"cpu,omitempty"`
	// Memory is the memory inventory of the agent
	Memory HostMemory `json:"memory,omitempty"`
	// OS is the OS inventory of the agent
	OS HostOS `json:"os,omitempty"`
	// Storage is the storage inventory of the agent
	Storage HostStorage `json:"storage,omitempty"`
	// Metadata is the metadata inventory of the agent
	Metadata map[string]string `json:"metadata,omitempty"`
}

// HostNetwork is the structure that houses the network inventory of the agent
type HostNetwork struct {
	// Interfaces is the list of interfaces on the agent
	Interfaces []HostInterface `json:"interfaces,omitempty"`
}

// HostInterface is the structure that houses the interface inventory of the agent
type HostInterface struct {
	// Name is the name of the interface
	Name string `json:"name,omitempty"`
	// MAC is the MAC address of the interface
	MAC string `json:"mac,omitempty"`
	// IP is the list of IP address(es) of the interface
	IP []string `json:"ip,omitempty"`
}

// HostCPU is the structure that houses the CPU inventory of the agent
type HostCPU struct {
	// Count is the number of CPUs on the agent
	Count int `json:"count,omitempty"`
	// Model is the model of the CPUs on the agent
	Model string `json:"model,omitempty"`
	// Architecture is the architecture of the CPUs on the agent
	Architecture string `json:"architecture,omitempty"`
}

// Memory is the structure that houses the memory inventory of the agent
type HostMemory struct {
	// Total is the total amount of memory on the agent
	Total int `json:"total,omitempty"`
}

// OS is the structure that houses the OS inventory of the agent
type HostOS struct {
	// Name is the name of the OS on the agent
	Name string `json:"name,omitempty"`

	// Version is the version of the OS on the agent
	Version string `json:"version,omitempty"`

	// Uptime is the uptime of the OS on the agent
	Uptime int `json:"uptime,omitempty"`
}

// Storage is the structure that houses the storage inventory of the agent
type HostStorage struct {
	// Disks is the list of disks on the agent
	Disks []HostDisk `json:"disks,omitempty"`
}

// Disk is the structure that houses the disk inventory of the agent
type HostDisk struct {
	// Name is the name of the disk
	Name string `json:"name,omitempty"`
	// Size is the size of the disk
	Size int `json:"size,omitempty"`
	// Type is the type of the disk
	Type string `json:"type,omitempty"`
}
