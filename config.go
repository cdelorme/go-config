package config

/**
 * Configuration abstraction
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-7
 */

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Settings
	Preferences
	File string
}

func (config *Config) Load() error {
	if config.File == "" {
		return errors.New("No file supplied...")
	}

	openFile, fileErr := os.Open(config.File)
	defer openFile.Close()
	if fileErr != nil {
		return fileErr
	}

	decoder := json.NewDecoder(openFile)
	jsonErr := decoder.Decode(&config.Data)
	if jsonErr != nil {
		return jsonErr
	}

	config.Override()

	return nil
}

func (config *Config) Save() error {
	if config.File == "" {
		return errors.New("No file supplied...")
	}

	openFile, fileErr := os.Create(config.File)
	defer openFile.Close()
	if fileErr != nil {
		return fileErr
	}

	data, jsonErr := json.MarshalIndent(&config.Data, "", "    ")
	if jsonErr != nil {
		return jsonErr
	}

	_, writeErr := openFile.Write(data)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func (config *Config) Override() {
	config.Parse()
	for _, flag := range config.Options {
		config.Set(*flag.Value, flag.Keys...)
	}
}
