package main

import (
	"fmt"
	"text/template"
	"log"
	"bytes"
	"github.com/hoisie/mustache"
)

type TemplateExec struct {
	Title string
	FooList []string
	FooMap map[string]int
	Content string
}

var templateData = TemplateExec{
	Title: "test",
	FooList: []string{"abc", "def", "ghi", "foo"},
	FooMap: map[string]int{"a": 111, "b": 222, "c": 333},
}

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	fmt.Println("----------------------------")
	fmt.Println("text/template")
	fmt.Println("----------------------------")

	textTemplate()

	fmt.Println("----------------------------")
	fmt.Println("github.com/hoisie/mustache")
	fmt.Println("----------------------------")

	hoisieMustache()

	fmt.Println("----------------------------")
	fmt.Println("github.com/benbjohnson/ego")
	fmt.Println("----------------------------")
	benbjohnsonEgo()
}

func textTemplate() {

	var sampleTemplate = `

## 表示
{{.Title}}

## 変数の代入
{{$title := .Title }}
$title = {{$title}}

## for文
{{range .FooList}}{{.}}
{{end}}

## mapのfor文
{{range $key, $val := .FooMap}}キー = {{$key}} 値 = {{$val}} $変数のスコープ = {{$title}}
{{end}}

## templateの入れ子
{{template "footer" $title}}
`

	var sampleFooterTemplate = `
> 子要素への変数渡し {{.}}`


	t := template.Must(template.New("foo").Parse(sampleTemplate))
	sub := template.Must(template.New("footer").Parse(sampleFooterTemplate))
	t.AddParseTree("footer", sub.Tree)

	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())
}

func hoisieMustache() {

	var sampleTemplate = `

## 表示
{{Title}}

## 変数の代入
分からなかった...

## for文
{{#FooList}}
{{.}}
{{/FooList}}

## mapのfor文
分からなかった...

## templateの入れ子
{{{content}}}
`

	var sampleFooterTemplate = `
> 子要素への変数渡し {{Title}}`

	text := mustache.RenderInLayout(sampleFooterTemplate, sampleTemplate, templateData)
	fmt.Println(text)
}

//go:generate ego -package main templates
func benbjohnsonEgo() {

	var subBuf bytes.Buffer
	if err := MySubTmpl(&subBuf, "Hoge"); err != nil {
		log.Fatalln(err)
	}
	templateData.Content = subBuf.String()

	var buf bytes.Buffer
	if err := MyTmpl(&buf, templateData); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())
}
