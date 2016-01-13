package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
	"github.com/kyokomi-sandbox/go-sandbox/other/waf/service"
)

type Quest struct {
	ID        int
	Name      string
	Detail    string
	CreatedAt time.Time
}

func main() {
	http.HandleFunc("/", service.Account.Index)

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	service.TransactionHandle("/test", db, func(tx *sql.Tx, w http.ResponseWriter, r *http.Request) {
		log.Println("insert")
		if _, err := tx.Exec("insert into quest values(1111, 'hoge', 'hgoe', now())"); err != nil {
			panic(err)
		}
	})

	service.TransactionHandle("/quest", db, func(tx *sql.Tx, w http.ResponseWriter, r *http.Request) {
		log.Println("select")
		rows, err := tx.Query("select * from quest where id = 1111")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var q Quest
			rows.Scan(&q.ID, &q.Name, &q.Detail, &q.CreatedAt)
			log.Println(q)
		}
	})

	http.Handle("/hoge", service.TransactionHandlerFunc(db, func(tx *sql.Tx, w http.ResponseWriter, r *http.Request) {
		log.Println("delete")
		if res, err := tx.Exec("delete from quest where id = 1111"); err != nil {
			log.Println("delete err", err)
			panic(err)
		} else if num, err := res.RowsAffected(); err != nil || num == 0 {
			log.Println("delete row err", err)
			panic(errors.New("delete row err"))
		}

		log.Println("select")
		row := tx.QueryRow("select * from quest where id = 1111")
		if err != nil {
			panic(err)
		}
		var q Quest
		row.Scan(&q.ID, &q.Name, &q.Detail, &q.CreatedAt)
		log.Println(q)

		log.Println("insert")
		if _, err := tx.Exec("insert into quest values(1111, 'hoge', 'hgoe', now())"); err != nil {
			panic(err)
		}
		log.Println("end")
	}))

	http.ListenAndServe(":8000", nil)
}
