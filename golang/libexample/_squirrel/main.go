package main

import (
	"fmt"

	sq "github.com/lann/squirrel"
)

func main() {
	users := sq.Select("id, name, age, detail").From("users").Join("emails USING (email_id)")

	fmt.Println(users.Where("name IN (?, ?)", "Dumbo", "Verna").ToSql())
	fmt.Println(users.Where("age = ?", 100).ToSql())
}
