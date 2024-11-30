package generators

import (
	"fmt"
	"gogo/logger"
	"gogo/util"
	"os"
	"os/exec"
)

func generateCOrCppProject(projtype int, log logger.Logger) {
	log.Info(fmt.Sprintf("Generating %s project", util.TypeIntToStr[projtype]))
	util.CreateFolder("src")
	var ptype int
	if projtype == util.C {
		ptype = util.C
	} else {
		ptype = util.CPP
	}
	err := generateMakefile(ptype, log)
	if err != nil {
		log.Error(err.Error())
	}

	curWd, _ := os.Getwd()
	util.GotoFolder("src")
	err = generateMainFile(ptype, log)
	if err != nil {
		log.Error(err.Error())
	}
	util.GotoFolder(curWd)
}

func generateGoProject(log logger.Logger) {
	util.CreateFolder("src")
	cmd := exec.Command("go", "mod", "init", "main")
	err := cmd.Run()
	if err != nil {
		log.Error(fmt.Sprintf("Could not run command: %v", err))
	}
	err = generateMakefile(util.GO, log)
	if err != nil {
		log.Warn(err.Error())
	}

	curWd, _ := os.Getwd()
	util.GotoFolder("src")
	generateMainFile(util.GO, log)
	util.GotoFolder(curWd)
}

func generateVProjekt(log logger.Logger) {
	util.CreateFolder("src")
	err := generateMakefile(util.VLANG, log)
	if err != nil {
		log.Warn(err.Error())
	}

	curWd, _ := os.Getwd()
	util.GotoFolder("src")
	generateMainFile(util.VLANG, log)
	util.GotoFolder(curWd)
}
