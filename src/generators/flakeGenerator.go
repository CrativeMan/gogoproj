package generators

import (
	"bytes"
	"fmt"
	"gogo/logger"
	"gogo/util"
	"os/exec"
)

func GenerateFlake(projtype int, log logger.Logger) error {
	command := util.Flake
	command[4] = fmt.Sprintf("https://flakehub.com/f/the-nix-way/dev-templates/*#%s", util.TypeIntToStr[projtype])
	log.Info("Generating flake.nix and flake.lock files")
	cmd := exec.Command(command[0], command[1:]...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()
	if err != nil {
		log.Info(cmd.String())
		log.Error(fmt.Sprintf("Could not run command: %v", err))
		log.Error(fmt.Sprintf("Command stderr: %s", stderr.String()))
		return err
	}

	err = allowDirenv()
	if err != nil {
		log.Warn(fmt.Sprintf("Could not run command: %v", err))
		return err
	}
	log.Info("Flake.nix and flake.lock files generated")
	return nil
}
