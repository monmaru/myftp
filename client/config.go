package client

import (
	"os"

	"github.com/pkg/errors"
)

// Config ...
type Config struct {
	Address     string
	Certificate string
	SrcDir      string
	Parallelism int
}

func (cfg *Config) validate() error {
	if cfg.Address == "" {
		return errors.New("Address must be specified")
	}

	if cfg.Parallelism <= 0 {
		return errors.New("Parallelism must be greater than zero")
	}

	_, err := os.Stat(cfg.SrcDir)
	return err
}
