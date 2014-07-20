package main

import (
	"testing"
	"strings"
)

func TestCreateTemplateTree(t *testing.T) {
	a := "Hoge"
	b := "Fuga"
	data := CreateTemplateTree(a, b)
	if data.A != a {
		t.Errorf("template text1 mismatch %s", data.A)
	}
	if data.B != b {
		t.Errorf("template text2 mismatch %s", data.B)
	}

	text := data.Execute()
	if !strings.Contains(text, a) {
		t.Errorf("template execute result contains %s", a)
	}
	if !strings.Contains(text, b) {
		t.Errorf("template execute result contains %s", b)
	}
}

