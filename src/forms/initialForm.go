package forms

import (
	"main/util"

	"github.com/charmbracelet/huh"
)

func InitialForm() (int, bool, string, error) {
	var projType int
	var flakeCheck bool
	var folder string
	var initialForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose programming language: ").
				Options(
					huh.NewOption("C", util.C),
					huh.NewOption("C++", util.CPP),
					huh.NewOption("Go", util.GO),
					huh.NewOption("Rust", util.RUST),
				).
				Value(&projType),

			huh.NewConfirm().
				Title("Don't generate a flake.nix and flake.lock file?").
				Value(&flakeCheck).
				Affirmative("Yes").
				Negative("No"),

			huh.NewInput().
				Title("Which folder do you want to use?").
				Placeholder("Folder").
				Description("The folder where the project will be generated").
				Value(&folder),
		),
	)

	err := initialForm.Run()
	if err != nil {
		return 0, false, "", err
	} else {
		return projType, flakeCheck, folder, nil
	}
}
