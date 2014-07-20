package main

import (
	"fmt"
	"text/template"
	"bytes"
	"log")

const baseTemplateText = `
Base {{.A}} World {{template "sub" .}}
`

const subTemplateText = `
Sub {{.B}} World
`

type ExecuteData struct {
	A string
	B string
	t *template.Template
}

func (e *ExecuteData) Execute() string {

	t := template.Must(e.t.Parse(baseTemplateText))

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, e); err != nil {
		log.Fatalln(err)
	}
	return buffer.String()
}

func CreateTemplateTree(text1, text2 string) *ExecuteData {
	fmt.Println("DoTemplateTree")

	base := template.New("base")
	sub := template.Must(template.New("sub").Parse(subTemplateText))
	base.AddParseTree("sub", sub.Tree)

	fmt.Println("Lookup", base.Lookup("sub"))

	return &ExecuteData {
		A: text1,
		B: text2,
		t: base,
	}
}
