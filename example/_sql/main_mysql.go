package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"flag"
	"fmt"
)

type Sample struct {
	ID int
	Name string
}

func main() {

	var user string
	var pass string
	var host string
	var dbName string
	flag.StringVar(&user, "u", "", "user name")
	flag.StringVar(&pass, "p", "", "password")
	flag.StringVar(&host, "h", "", "host name")
	flag.StringVar(&dbName, "D", "", "database name")
	flag.Parse()

	log.SetFlags(log.Llongfile)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var sample Sample
	err = db.QueryRow("SELECT * FROM sample LIMIT ?", 1).Scan(&sample.ID, &sample.Name)
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(sample)
}
