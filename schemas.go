package main

import (
	md "github.com/nao1215/markdown"
)

type Schemas struct {

}

type Schema struct {
	Table []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Resource string `xml:"resource,attr"`
		Engine   string `xml:"engine,attr"`
		Comment  string `xml:"comment,attr"`
		Column   []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Name      string `xml:"name,attr"`
			Unsigned  string `xml:"unsigned,attr"`
			Nullable  string `xml:"nullable,attr"`
			Identity  string `xml:"identity,attr"`
			Comment   string `xml:"comment,attr"`
			Default   string `xml:"default,attr"`
			Length    string `xml:"length,attr"`
			OnUpdate  string `xml:"on_update,attr"`
			Scale     string `xml:"scale,attr"`
			Precision string `xml:"precision,attr"`
			Padding   string `xml:"padding,attr"`
		} `xml:"column"`
		Constraint []struct {
			Text            string `xml:",chardata"`
			Type            string `xml:"type,attr"`
			ReferenceId     string `xml:"referenceId,attr"`
			Table           string `xml:"table,attr"`
			AttrColumn      string `xml:"column,attr"`
			ReferenceTable  string `xml:"referenceTable,attr"`
			ReferenceColumn string `xml:"referenceColumn,attr"`
			OnDelete        string `xml:"onDelete,attr"`
			Column          []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"column"`
		} `xml:"constraint"`
		Index []struct {
			Text        string `xml:",chardata"`
			ReferenceId string `xml:"referenceId,attr"`
			IndexType   string `xml:"indexType,attr"`
			Column      []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"column"`
		} `xml:"index"`
	} `xml:"table"`
}

func (s *Schemas) Generate(cnf Config, markdown *md.Markdown) {
	areaMap := NewXml("db_schema").UnmarshalToMap(Schema{}, cnf)

	if len(areaMap) == 0 {
		return
	}

	markdown.H2("Schema")

	for _, schema := range areaMap {
		tables := schema.(*Schema).Table

		var rows [][]string
		for _, table := range tables {
			row := []string{table.Name, table.Comment}
			rows = append(rows, row)
		}

		if len(tables) == 0 {
			continue
		}

		markdown.CustomTable(
			md.TableSet{
				Header: []string{"Table Name", "Comment"},
				Rows:   rows,
			},
			md.TableOptions{
				AutoWrapText: false,
			},
		)
	}
}
