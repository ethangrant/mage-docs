package main

import (
	"github.com/charmbracelet/huh"
)

type InteractiveForm struct {
}

func (i *InteractiveForm) Start(a ArgValidator, path *string, outputFile *string, selected *[]string) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Absolute path to module dir").Value(path).Validate(func(str string) error {
				return a.ModulePath(AddTrailingSlash(*path))
			}),
			huh.NewInput().Title("Name of output file").Value(outputFile).Validate(func(str string) error {
				return a.OutputFile(*outputFile)
			}),
			huh.NewMultiSelect[string]().Title("Select documentation you would like to render").Options(
				huh.NewOption("Module", "Module").Selected(true),
				huh.NewOption("Routes", "Routes"),
				huh.NewOption("Webapi", "Webapi"),
				huh.NewOption("Layouts", "Layouts"),
				huh.NewOption("Mixins", "Mixins"),
				huh.NewOption("Schemas", "Schemas"),
				huh.NewOption("Dispatchers", "Dispatchers"),
				huh.NewOption("Observers", "Observers"),
				huh.NewOption("Plugins", "Plugins"),
				huh.NewOption("Preferences", "Preferences"),
			).Value(selected),
		),
	)

	err := form.Run()
	if err != nil {
		return err
	}

	return nil
}
