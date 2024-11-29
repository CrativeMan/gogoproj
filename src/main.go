package gogo

import (
	"fmt"
	"gogo/forms"
	"gogo/generators"
	"gogo/logger"
	"gogo/util"
)

type Project struct {
	projType int
	folder   string
}

var err error

var log logger.Logger
var project Project

func init() {
	project = Project{}
	log = logger.NewLogger()
}

func main() {
	project.projType, project.folder, err = forms.InitialForm()
	if err != nil {
		log.Error("Could not run initial form")
		return
	} else {
		log.Info(fmt.Sprintf("Project type: %s", util.TypeIntToStr[project.projType]))
	}

	err = util.GotoFolder(project.folder)
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = generators.GenerateFlake(project.projType, log)
	if err != nil {
		log.Error("Could not generate flake.nix and flake.lock files")
		return
	}

	err = generators.GenerateInitialProjectFiles(project.projType, log)
	if err != nil {
		log.Error("Could not generate initial project files")
		return
	}
}
