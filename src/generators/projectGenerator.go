package generators

import (
	"fmt"
	"main/logger"
	"main/util"
)

func GenerateInitialProjectFiles(projtype int, log logger.Logger) error {
	log.Info(fmt.Sprintf("Generating initial project files for %s", util.TypeIntToStr[projtype]))
	switch projtype {
	case util.C:
		log.Info("C proj")
	case util.CPP:
		log.Info("CPP proj")
	case util.GO:
		log.Info("GO proj")
	case util.RUST:
		log.Info("RUST proj")
	default:
		log.Error("Invalid project type")
	}

	return nil
}
