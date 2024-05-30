package main

import (
	"github.com/charmbracelet/huh"
)

type InteractiveForm struct {
}

func (i *InteractiveForm) Start(a ArgValidator, path *string, outputFile *string) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Absolute path to module dir").Value(path).Validate(func(str string) error {
				return a.ModulePath(AddTrailingSlash(*path))
			}),
			huh.NewInput().Title("Name of output file").Value(outputFile).Validate(func(str string) error {
				return a.OutputFile(*outputFile)
			}),
		),
	)

	err := form.Run()
	if err != nil {
		return err
	}

	return nil
}
