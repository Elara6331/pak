package main

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"path/filepath"
)

// Config contains the root of the TOML config file
type Config struct {
	ActiveManager string
	RootCommand   string
	Managers      map[string]Manager
}

// Manager contains the root of all manager sections in the TOML config file
type Manager struct {
	UseRoot   bool
	Commands  map[string]string
	Shortcuts map[string]string
}

// Create new Config{} with values from file path given
func NewConfig(path string) Config {
	// Read file at path
	data, err := ioutil.ReadFile(path)
	if err != nil {
		Log.Fatal().Err(err).Str("file", filepath.Base(path)).Msg("Error reading config")
	}
	// Create new Config{}
	cfg := Config{}
	// Unmarshal TOML in config
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		Log.Fatal().Err(err).Str("file", filepath.Base(path)).Msg("Error unmarshalling TOML")
	}
	// Return config
	return cfg
}
