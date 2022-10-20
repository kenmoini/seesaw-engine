package server

import (
	"github.com/kenmoini/seesaw-server/internal/config"
)

// CLIOpts contains the CLI options
type CLIOpts struct {
	Config string
}

// Config struct for webapp config at the top level
type Config struct {
	Seesaw config.Seesaw `yaml:"seesaw"`
}
