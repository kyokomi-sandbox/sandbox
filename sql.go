package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
	"fmt"
	"math/rand"
	"time"
	"errors"
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

type MItem struct {
	ItemID int
	ImageID string
	ItemType int
	Name string
	Detail string
	Param int
}

func MItemSelectByID(db *sql.DB, itemID int) (*MItem, error) {
	q := `
select
	item_id
	,image_id
	,type
	,name
	,detail
	,param
from
	m_item
where
	item_id = $1
`
	var rows *sql.Rows
	var err error
	rows, err = db.Query(q, itemID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var mItem MItem
		if err := rows.Scan(&mItem.ItemID, &mItem.ImageID, &mItem.ItemType, &mItem.Name, &mItem.Detail, &mItem.Param); err != nil {
			fmt.Println(err)
		}
		return &mItem, nil
	}

	return nil, errors.New("not found")
}


