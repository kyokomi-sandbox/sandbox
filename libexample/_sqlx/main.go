package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

var queryString = `
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
and tinytext_columns = /*:tinytext_columns*/'piyo'

`

var argsMap = map[string]interface{}{
	"id":               1000,
	"text_columns":     "fuga",
	"tinytext_columns": "hoge",
}

func sqlxNamed(query string, args interface{}) (string, []interface{}, error) {
	return sqlx.Named(query, args)
}

func sqlxNamedReplaceQueryComment(query string, args interface{}) (string, []interface{}, error) {
	return sqlx.Named(replaceQueryComment(query), args)
}

func replaceQueryComment(query string) string {
	// */〜スペースまたは改行までの文字を削る
	reg1 := regexp.MustCompile(`\*/[^ |\n]*[ ]`)
	reg2 := regexp.MustCompile(`\*/[^ |\n]*[\n]`)
	query = strings.Replace(query, "/*", "", -1)
	if reg1.MatchString(query) {
		query = reg1.ReplaceAllString(query, " ")
	}

	if reg2.MatchString(query) {
		query = reg2.ReplaceAllString(query, "\n")
	}

	return query
}

func main() {
	query, args, err := sqlxNamedReplaceQueryComment(queryString, argsMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(query, args)
}
