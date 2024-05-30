package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh/spinner"
)

func main() {
	var (
		argValidator    ArgValidator
		interactiveForm InteractiveForm
		path            string
		outputFile      string
		generators      []string
	)

	generators = []string{}
	outputFile = "Mage_Docs_README.md"

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = interactiveForm.Start(argValidator, &path, &outputFile, &generators)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path = AddTrailingSlash(path)

	cnf := Config{
		ModulePath: path,
	}

	file, err := os.Create(cnf.ModulePath + outputFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	err = spinner.New().
		Title("Generating documentation").
		Action(func() {
			err = NewRenderer().Render(file, generators, cnf)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}).
		Run()
}
