package dao

import (
	"database/sql"
	"log"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type QueryArgs map[string]interface{}

func NewDao(options Options) (*Dao, error) {

	var d Dao
	d.options = options

	// TODO: 引数でもらうようにする
	db, err := sql.Open("mysql", "admin:password@tcp(localhost:3306)/test")
	if err != nil {
		return nil, err
	}

	d.DB = db

	d.setupDao()

	return &d, nil
}

func (d *Dao) Close() error {
	d.Println("dao close")

	err := d.DB.Close()

	return err
}

func (d *Dao) queryArgs(queryString string, args QueryArgs) string {

	d.Println("old: ", queryString)

	for key, val := range args {
		re := regexp.MustCompile(`\/\* ` + key + ` \*\/.*`)

		replaceWord := ""
		switch val.(type) {
		case int:
			replaceWord = strconv.Itoa(val.(int))
		case string:
			replaceWord = "\"" + val.(string) + "\""
		}
		queryString = re.ReplaceAllString(queryString, replaceWord)
	}

	d.Println("new: ", queryString)

	return queryString
}

func (d *Dao) Println(v ...interface{}) {
	if d.options.Debug {
		log.Println(v)
	}
}
