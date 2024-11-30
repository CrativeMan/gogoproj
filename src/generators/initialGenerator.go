package generators

import (
	"fmt"
	"gogo/logger"
	"gogo/util"
	"io"
	"net/http"
	"os/exec"
)

func GenerateInitialProjectFiles(projtype int, log logger.Logger) error {
	switch projtype {
	case util.C:
		generateCOrCppProject(projtype, log)
	case util.CPP:
		generateCOrCppProject(projtype, log)
	case util.GO:
		generateGoProject(log)
	case util.RUST:
		log.Info("Use 'cargo init' in the new folder to create project")
	case util.VLANG:
		generateVProjekt(log)
	default:
		return fmt.Errorf("invalid project type: %d", projtype)
	}

	log.Info("Generated initial project files")
	return nil
}

func getRawFileData(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code: %s", resp.Status)
	}

	return resp.Body, nil
}

func generateMakefile(projtype int, log logger.Logger) error {
	log.Info(fmt.Sprintf("Generating Makefile for %s", util.TypeIntToStr[projtype]))

	outputFile := "makefile"

	body, err := getRawFileData(util.TemplatesMakefile[projtype])
	if err != nil {
		return err
	}
	defer body.Close()

	util.WriteToFile(outputFile, body)

	log.Info("Generated Makefile")
	return nil
}

func generateMainFile(projtype int, log logger.Logger) error {
	log.Info(fmt.Sprintf("Generating main file for %s", util.TypeIntToStr[projtype]))

	outputFile := "main" + util.TypeExtension[projtype]

	body, err := getRawFileData(util.TemplatesMain[projtype])
	if err != nil {
		return err
	}
	defer body.Close()

	util.WriteToFile(outputFile, body)

	log.Info("Generated main file")
	return nil
}

func allowDirenv() error {
	cmd := exec.Command("direnv", "allow")
	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return err
}
