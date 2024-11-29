package main

import (
	"errors"
	"fmt"
	"os"
)

func gotoProjectFolder() error {
	absPath, err := getAbsoluteProjectPath(project.folder, project.user)
	if absPath == "" || err != nil {
		nError := errors.New("could not get absolute path")
		return nError
	}
	log.Info(fmt.Sprintf("Generating project in: %s", absPath))

	// manouver to the project folder
	err = os.Chdir(absPath)
	if err != nil {
		nError := errors.New("could not change directory")
		return nError
	}

	return nil
}

func getAbsoluteProjectPath(path string, user User) (string, error) {
	// check if path is empty
	if path == "" {
		return "", errors.New("path is empty")
	}

	// get global path
	if path[0] == '~' {
		path = user.homeDir + path[1:]
	} else if len(path) == 1 && path[0] == '.' {
		wd, _ := os.Getwd()
		path = wd
	} else if path[0] == '.' && path[1] == '/' {
		wd, _ := os.Getwd()
		path = wd + path[1:]
	}

	// check if path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// generate folder
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return "", err
		} else {
			log.Info("Folder created")
		}
	}

	return path, nil
}
