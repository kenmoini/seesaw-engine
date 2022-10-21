package config

import "time"

const (
	// AppName is the name of the application
	AppName = "seesaw-engine"
	// AppVersion is the version of the application
	AppVersion = "0.0.1"
	// DefaultBasePath is the base path for the application API
	DefaultBasePath = "api"
)

// Seesaw is the structure that houses the root of the configuration
type Seesaw struct {
	// Config is the configuration of this instance
	Config Config `json:"config,omitempty"`
	// Server is the configuration for the Seesaw server
	Server Server `json:"server"`
	// Database is the configuration for the Seesaw database
	Database Database `json:"database"`
}

// Config is the structure that houses the general configuration
type Config struct {
	// Mode is the mode that the application is running in, can be "server", "agent", or "hybrid"
	Mode string `json:"mode"`
	// PKI is the structure that houses the PKI configuration
	PKI PKI `json:"pki,omitempty"`
}

// PKI is the structure that houses the PKI configuration
type PKI struct {
	// BasePath is the base path for the PKI files used in leveraging client certificates
	BasePath string `json:"base_path,omitempty"`

	// Clients is the client certificate configuration
	Clients Clients `json:"clients,omitempty"`

	// Signer is the signer certificate configuration
	Signer Signer `json:"signer,omitempty"`
}

// Clients is the structure that houses the client certificate configuration
type Clients struct {
	// Keysize is the key size for the client certificates
	Keysize int `json:"keysize,omitempty"`

	// Validity is the validity for the client certificates
	Validity int `json:"validity,omitempty"`
}

// Signer is the structure that houses the signer certificate configuration
type Signer struct {
	// Keysize is the key size for the signer certificate
	Keysize int `json:"keysize,omitempty"`

	// Validity is the validity for the signer certificate
	Validity int `json:"validity,omitempty"`

	// CommonName is the common name for the signer certificate
	CommonName string `json:"common_name,omitempty"`
}

// Server is the structure that houses the Seesaw server configuration
type Server struct {
	// BasePath is the base path for the Seesaw server
	BasePath string `json:"base_path,omitempty"`

	// Host is the address to listen on
	Host string `json:"host,omitempty"`

	// Port is the port to listen on
	Port string `json:"port,omitempty"`

	// TLS is the TLS configuration
	TLS TLS `json:"tls,omitempty"`

	// Timeout is the set of timeouts for the server
	Timeout struct {
		// Server is the general server timeout to use
		// for graceful shutdowns
		Server time.Duration `json:"server,omitempty"`

		// Write is the amount of time to wait until an HTTP server
		// write opperation is cancelled
		Write time.Duration `json:"write,omitempty"`

		// Read is the amount of time to wait until an HTTP server
		// read operation is cancelled
		Read time.Duration `json:"read,omitempty"`

		// Read is the amount of time to wait
		// until an IDLE HTTP session is closed
		Idle time.Duration `json:"idle,omitempty"`
	} `json:"timeout,omitempty"`
}

// TLS is the structure that houses the TLS configuration
type TLS struct {
	// Cert is the path to the TLS certificate
	Cert string `json:"cert,omitempty"`

	// Key is the path to the TLS key
	Key string `json:"key,omitempty"`

	// CA is the path to the TLS CA
	CA string `json:"ca,omitempty"`
}

// Database is the structure that houses the Seesaw database configuration
// This database is shared by this server and the frontend, both of which need write capabilities
type Database struct {
	// Type is the type of database to use
	Type string `json:"type,omitempty"`

	// Host is the host of the database
	Host string `json:"host,omitempty"`

	// Port is the port of the database
	Port int `json:"port,omitempty"`

	// User is the user to connect to the database
	Username string `json:"username,omitempty"`

	// Password is the password to connect to the database
	Password string `json:"password,omitempty"`

	// Database is the database to connect to
	Database string `json:"database,omitempty"`
}
