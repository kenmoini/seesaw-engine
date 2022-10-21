package startup

import (
	"flag"
	"fmt"
	"os"

	"github.com/kenmoini/seesaw-engine/internal/logging"
	"gopkg.in/yaml.v2"
)

//=================================================================================================
// Preflight Functions
//=================================================================================================

// PreflightSetup just makes sure the stage is set before starting the application in general
func PreflightSetup() {
	logging.LogStdOutInfo("Preflight complete!")
}

// ServerPreflightSetup just makes sure the stage is set before starting the HTTP server
func ServerPreflightSetup() {
	logging.LogStdOutInfo("Server Mode Preflight complete!")
}

// AgentPreflightSetup just makes sure the stage is set before starting the seesaw agent
func AgentPreflightSetup() {
	logging.LogStdOutInfo("Agent Mode Preflight complete!")
}

// HybridPreflightSetup just makes sure the stage is set before starting each component
func HybridPreflightSetup() {
	// Run the server preflight
	ServerPreflightSetup()

	// Run the agent preflight
	AgentPreflightSetup()

	// Log that the preflight is complete
	logging.LogStdOutInfo("Hybrid Mode Preflight complete!")
}

//=================================================================================================
// CLI Option Parsing
//=================================================================================================

// ParseFlags will define and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (CLIOpts, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "", "path to config file, eg '-config=./config.yml'")

	// Actually parse the flags
	flag.Parse()

	if configPath == "" {
		return CLIOpts{}, logging.Stoerr("No server configuration defined! (-config=./config.yml)")
	} else {
		// Validate the path first
		if err := ValidateConfigPath(configPath); err != nil {
			return CLIOpts{}, err
		}
	}

	SetCLIOpts := CLIOpts{
		Config: configPath}

	// Return the configuration path
	return SetCLIOpts, nil
}

//=================================================================================================
// Configuration Loading
//=================================================================================================

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath CLIOpts) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath.Config)
	logging.CheckAndFail(err, "Failed to open config file")
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	//readConfig = config

	return config, nil
}

// StartEngine will start the engine in the mode specified in the configuration
func StartEngine() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := ParseFlags()
	logging.CheckAndFail(err, "Failed to parse CLI Opt flags")

	// Run general preflight
	PreflightSetup()

	// Setup engine config
	cfg, err := NewConfig(cfgPath)
	logging.CheckAndFail(err, "Failed to parse server configuration")

	// Run the engine in the mode specified in the configuration
	switch cfg.Seesaw.Config.Mode {
	case "server":
		// Run the server preflight
		ServerPreflightSetup()

		// Start the server
		cfg.RunHTTPServer()
	case "agent":
		// Run the agent preflight
		AgentPreflightSetup()

		// Start the agent
		cfg.RunAgent()

	case "hybrid":
		// Run the hybrid preflight
		HybridPreflightSetup()

		// Start the server
		cfg.RunHTTPServer()

		// Start the agent
		cfg.RunAgent()
	default:
		logging.LogStdErr("Invalid mode specified in configuration!")
	}

}
