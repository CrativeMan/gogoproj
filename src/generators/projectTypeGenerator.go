package generators

import (
	"fmt"
	"main/logger"
	"main/util"
	"os"
	"os/exec"
)

func GenerateCOrCppProject(projtype int, log logger.Logger) {
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

func GenerateGoProject(log logger.Logger) {
	util.CreateFolder("src")
	cmd := exec.Command("go", "mod", "init", "main")
	err := cmd.Run()
	if err != nil {
		log.Error(fmt.Sprintf("Could not run command: %v", err))
	}
	generateMakefile(util.GO, log)

	curWd, _ := os.Getwd()
	util.GotoFolder("src")
	generateMainFile(util.GO, log)
	util.GotoFolder(curWd)
}
