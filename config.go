package config

/**
 * Configuration abstraction
 *
 * Author: Casey Delow <cdelorme@gmail.com>
 * Date: 2014-8-5
 */

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DataBag
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
