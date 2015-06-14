// Package config provides a stupid-simple abstraction to loading
// and saving json configuration files in the format of
// a map[string]interface

package config

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path"
)

// Load looks for filePath, and loops through standard storage paths
// starting with XDG_CONFIG_DIR, then HOME, and finally /etc
// it returns after finding the first file, even if an error occurs
func Load(filePath string) (map[string]interface{}, error) {
	try := []string{filePath}

	appName := path.Base(os.Args[0])

	home := os.Getenv("HOME")
	if home == "" {
		if usr, err := user.Current(); err == nil {
			home = usr.HomeDir
		}
	}

	configDir := os.Getenv("XDG_CONFIG_DIR")
	if configDir == "" {
		configDir = ".config"
	}

	// I could have used a loop, but this was much more strait forward
	try = append(try, path.Join(home, configDir, appName, appName))
	try = append(try, path.Join(home, configDir, appName, appName+".json"))
	try = append(try, path.Join(home, configDir, appName, appName+".conf"))
	try = append(try, path.Join(home, "."+appName+".conf"))
	try = append(try, path.Join(home, "."+appName+".conf"))
	try = append(try, path.Join(home, "."+appName+".conf"))
	try = append(try, path.Join("/", "etc", appName))
	try = append(try, path.Join("/", "etc", appName+".json"))
	try = append(try, path.Join("/", "etc", appName+".conf"))

	for _, f := range try {
		openFile, fileErr := os.Open(f)
		defer openFile.Close()
		if fileErr == nil {
			conf := make(map[string]interface{})
			decoder := json.NewDecoder(openFile)
			jsonErr := decoder.Decode(&conf)
			if jsonErr != nil {
				return nil, jsonErr
			}
			return conf, nil
		}
	}

	return nil, errors.New("Unable to locate file for loading")
}

// Save accepts a filePath and a map[string]interface{} of data
// and attempts to save to the filePath, or if filePath is empty
// it attempts to save to the default XDG_CONFIG_DIR path
func Save(filePath string, data map[string]interface{}) error {
	if filePath == "" {
		appName := path.Base(os.Args[0])
		home := os.Getenv("HOME")
		if home == "" {
			if usr, err := user.Current(); err == nil {
				home = usr.HomeDir
			}
		}
		configDir := os.Getenv("XDG_CONFIG_DIR")
		if configDir == "" {
			configDir = ".config"
		}
		filePath = path.Join(home, configDir, appName, appName)
	}

	os.MkdirAll(path.Dir(filePath), os.ModePerm)

	openFile, fileErr := os.Create(filePath)
	defer openFile.Close()
	if fileErr != nil {
		return fileErr
	}

	js, jsonErr := json.MarshalIndent(data, "", "    ")
	if jsonErr != nil {
		return jsonErr
	}

	_, writeErr := openFile.Write(js)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
