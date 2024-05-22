package main

import (
	"fmt"
	"time"

	md "github.com/nao1215/markdown"
)

type Readme struct {
	Markdown *md.Markdown
	GeneratedAt time.Time
}

type Section struct {
	Header string
	Content string
}

func (s *Section) OutputHeader() string {
	return "##" + s.Header
}

func (s * Section) Render() (o string) {
	o += fmt.Sprintf("##%s", s.Header)

	return o;
}

// func (r * Readme) Write(md) (o string) {
// 	o += r.Content
// 	o += r.GeneratedAt.String()

// 	return o;
// }