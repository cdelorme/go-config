package config

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path"
)

func Load(file string) (map[string]interface{}, error) {
	conf := make(map[string]interface{})

	files := []string{file}
	appName := path.Base(os.Args[0])
	usr, err := user.Current()
	if err == nil {
		files = append(files, usr.HomeDir+"/."+appName)
	}
	files = append(files, "/etc/"+appName)

	for _, f := range files {
		openFile, fileErr := os.Open(f)
		defer openFile.Close()
		if fileErr == nil {
			decoder := json.NewDecoder(openFile)
			jsonErr := decoder.Decode(&conf)
			if jsonErr != nil {
				return conf, jsonErr
			}
		}
	}

	return conf, nil
}

func Save(file string, data *map[string]interface{}) error {
	if file == "" {
		usr, err := user.Current()
		if err != nil {
			return errors.New("No supplied or accessible files for save operation")
		}
		file = usr.HomeDir + "/." + path.Base(os.Args[0])
	}

	openFile, fileErr := os.Create(file)
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
