package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	md "github.com/nao1215/markdown"
)

type Dispatches struct {
	Title string
}

func (d *Dispatches) Generate(cnf Config, markdown *md.Markdown) {
	events := make(map[string]string)

	filepath.WalkDir(cnf.ModulePath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}

		if filepath.Ext(d.Name()) != ".php" {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		dataString := string(data)
		// flatten the file to single line, make it easier to check and pull dispatched events.
		replacer := strings.NewReplacer("\n", "", "\r", "", " ", "")
		dataString = replacer.Replace(dataString)

		// primitive but most common way eventManager is used.
		re := regexp.MustCompile(`(eventManager->dispatch\([^;]*)`)
		event := re.FindString(dataString)

		if event == "" {
			return err
		}

		path = strings.Replace(path, cnf.ModulePath, "", -1)

		events[path] = event

		return nil
	})

	if len(events) == 0 {
		return
	}

	rows := [][]string{}
	for path, event := range events {
		row := []string{md.Code(path), md.Code(event)}
		rows = append(rows, row)
	}

	d.Title = RenderTitle(d.Title, "Dispatched Events", markdown)
	RenderTable([]string{"Path", "Dispatched event"}, rows, "", markdown)
}
