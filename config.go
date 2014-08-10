package config

/**
 * Configuration abstraction
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-10
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Settings
	Preferences
	File string
}

func (config *Config) Load() error {
	defer config.Override()

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
		if *flag.Value != "" {
			config.Set(*flag.Value, flag.Keys...)
		}
	}
}
