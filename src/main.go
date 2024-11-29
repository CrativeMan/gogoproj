package main

import (
	"fmt"
	"main/forms"
	"main/generators"
	"main/logger"
	"main/util"
	"os"
)

type User struct {
	homeDir string
}

type Project struct {
	projType   int
	folder     string
	flakeCheck bool
	user       User
}

var err error

var log logger.Logger
var project Project

func init() {
	project = Project{}
	log = logger.NewLogger()

	project.user.homeDir, err = os.UserHomeDir()
	if err != nil {
		log.Error("Could not get user home directory")
	}
}

func main() {
	// get project type, flake check and folder
	project.projType, project.flakeCheck, project.folder, err = forms.InitialForm()
	if err != nil {
		log.Error("Could not run initial form")
		return
	} else {
		log.Info(fmt.Sprintf("Project type: %s", util.TypeIntToStr[project.projType]))
	}

	// move to project folder
	err = gotoProjectFolder()
	if err != nil {
		log.Error(err.Error())
		return
	}

	generators.GenerateFlake(project.flakeCheck, project.projType, log)
	err = generators.GenerateInitialProjectFiles(project.projType, log)
	if err != nil {
		log.Error("Could not generate initial project files")
		return
	}
}
