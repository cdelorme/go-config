// Package config provides a stupid-simple abstraction to loading
// and saving json configuration files in the format of
// a map[string]interface

package config

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

var exts = []string{"", ".json", ".conf"}
var paths []string
var ConfigFile string

// init prepares the package to deal with standard paths
// including APPDATA, XDG_CONFIG_DIR, and user home for
// configuration storage, using conditions to prepare
func init() {
	var list []string
	appName := filepath.Base(os.Args[0])

	if p, e := filepath.EvalSymlinks(os.Args[0]); e == nil {
		if a, e := filepath.Abs(p); e == nil {
			for _, e := range exts[1:] {
				paths = append(paths, filepath.Join(filepath.Dir(a), appName+e))
			}
		}
	}
	if p := os.Getenv("APPDATA"); p != "" {
		list = append(list, filepath.Join(p, appName)+"/")
		ConfigFile = filepath.Join(p, appName, appName+".json")
	}
	home := os.Getenv("HOME")
	if home == "" {
		if u, err := user.Current(); err == nil {
			home = u.HomeDir
		}
	}
	xdg := os.Getenv("XDG_CONFIG_DIR")
	if xdg == "" {
		xdg = filepath.Join(home, ".config", appName)
	}
	list = append(list, xdg, home+"/.", filepath.Join("/etc/", appName)+"/", "/etc/")
	if ConfigFile == "" {
		ConfigFile = filepath.Join(xdg, appName+".json")
	}

	for _, p := range list {
		for _, e := range exts {
			paths = append(paths, p+appName+e)
		}
	}
}

// Load checks suggestedPaths then standard package-local paths
// and returns first results found, capturing the file for saving later
func Load(suggestedPaths ...string) (map[string]interface{}, error) {
	check := append(suggestedPaths, paths...)

	for _, f := range check {
		openFile, fileErr := os.Open(f)
		defer openFile.Close()
		if fileErr == nil {
			conf := make(map[string]interface{})
			decoder := json.NewDecoder(openFile)
			jsonErr := decoder.Decode(&conf)
			if jsonErr != nil {
				return nil, jsonErr
			}
			ConfigFile = f
			return conf, nil
		}
	}

	return nil, errors.New("Unable to locate file for loading")
}

// Save accepts an interface{} and optionally a filepath and will
// prepare the expected path, attempt to json marshal  the data
// then save to the supplied path, or fallback path which is
// APPDATA or XDG, returning any errors immediately
func Save(data interface{}, suggestedPaths ...string) error {
	check := append(suggestedPaths, ConfigFile)

	for _, f := range check {
		os.MkdirAll(filepath.Dir(f), os.ModePerm)
		o, e := os.Create(f)
		if e != nil {
			return e
		}
		js, e := json.MarshalIndent(data, "", "	")
		if e != nil {
			return e
		}
		if _, e := o.Write(js); e != nil {
			return e
		}
	}

	return nil
}
