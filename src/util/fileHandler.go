package util

import (
	"errors"
	"io"
	"os"
)

func CreateFolder(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		// generate folder
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func WriteToFile(fileName string, fileContents io.ReadCloser) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, fileContents)
	if err != nil {
		return err
	}

	return nil
}

func GotoFolder(folder string) error {
	absPath, err := getAbsoluteProjectPath(folder)
	if absPath == "" || err != nil {
		nError := errors.New("could not get absolute path")
		return nError
	}

	err = os.Chdir(absPath)
	if err != nil {
		nError := errors.New("could not change directory")
		return nError
	}

	return nil
}

func getAbsoluteProjectPath(path string) (string, error) {
	// check if path is empty
	if path == "" {
		return "", errors.New("path is empty")
	}

	// get global path
	if path[0] == '~' {
		path = os.Getenv("HOME") + path[1:]
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
		}
	}

	return path, nil
}
