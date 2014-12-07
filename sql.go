package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
	"fmt"
	"math/rand"
	"time"
)


const DROP_TABLE   = `DROP TABLE TEST_TBL;`
const CREATE_TABLE = `
CREATE TABLE TEST_TBL (
	ID int primary key,
	Name text,
	Detail text
);
`

const INSERT_TEST_TBL = `
INSERT INTO TEST_TBL (ID, Name, Detail) VALUES($1, $2, $3);
`

func resetDB(db *sql.DB) {
	if res, err := db.Exec(DROP_TABLE); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%+v\n", res)
	}

	if res, err := db.Exec(CREATE_TABLE); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}

func sqlExample() {

	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 dbname=example_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(){
		fmt.Println("close.")
		if err := db.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

//	resetDB(db)

	if tx, err := db.Begin(); err != nil {
		log.Fatalln("begin ", err)
	} else {
		var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		if res, err := tx.Exec(INSERT_TEST_TBL, r.Int31(), "hoge", "fugafuga"); err != nil {
			log.Fatalln("insert query ", err)
		} else {
			fmt.Printf("%+v\n", res)
		}

		defer func(){
			fmt.Println("commit.")
			if err := tx.Commit(); err != nil {
				log.Fatalln(err)
			}
		}()
	}
}
