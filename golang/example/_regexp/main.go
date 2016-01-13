package main

import (
	"fmt"
	"regexp"
	"strings"
)

var query = `
select
  id
, text_columns
, tinytext_columns
, mediumtext_columns
, longtext_columns
, char_columns
, varchar_columns
, enum_columns
FROM
  goma_string_types
WHERE
  id = /*:id*/1 and text_columns = /* :text_columns */'hoge'

`

func main() {
	reg1 := regexp.MustCompile(`\*/[^ ]*[ ]`)
	reg2 := regexp.MustCompile(`\*/[^ ]*[\n]`)
	query := strings.Replace(query, "/*", "", -1)
	if reg1.MatchString(query) {
		query = reg1.ReplaceAllString(query, " ")
	}

	if reg2.MatchString(query) {
		query = reg2.ReplaceAllString(query, "\n")
	}

	fmt.Println(query)
}