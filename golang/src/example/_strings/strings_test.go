package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	main()
}

func TestFieldsFunc(t *testing.T) {

	t1 := "aaaa112bbbb3333ccccc2ddddd3"
	fieldsResult := strings.FieldsFunc(t1, Is123)
	if len(fieldsResult) != 4 {
		t.Errorf("fields len error 4 != %d", len(fieldsResult))
	}
}
