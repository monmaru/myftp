package server

import (
	"os"

	"github.com/pkg/errors"
)

// Config ...
type Config struct {
	Address     string
	Certificate string
	Key         string
	DestDir     string
	LogDir      string
}

func (cfg *Config) validate() error {
	if cfg.Address == "" {
		return errors.Errorf("Address must be specified")
	}

	_, err := os.Stat(cfg.DestDir)
	if err != nil {
		return err
	}

	_, err = os.Stat(cfg.LogDir)
	return err
}
