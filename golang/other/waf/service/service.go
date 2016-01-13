package service

import (
	"database/sql"
	"log"
	"net/http"
)

type Service interface {
}

type TxServer struct {
	*sql.DB
}

func TransactionHandle(pattern string, db *sql.DB, txHandler func(*sql.Tx, http.ResponseWriter, *http.Request)) {
	txServer := TxServer{db}
	http.Handle(pattern, txServer.TransactionHandlerFunc(txHandler))
}

func TransactionHandlerFunc(db *sql.DB, txHandler func(*sql.Tx, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	txServer := TxServer{db}
	return txServer.TransactionHandlerFunc(txHandler)
}

func (s TxServer) TransactionHandlerFunc(txHandler func(*sql.Tx, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx, err := s.Begin()
		if err != nil {
			log.Println("begin error Rollback", err.Error())
			s.HandlerError(w, r)
			return
		}
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover Rollback", err)
				tx.Rollback()
				s.HandlerError(w, r)
				return
			}
		}()

		txHandler(tx, w, r)

		if err := tx.Commit(); err != nil {
			log.Println("commit error Rollback", err.Error())
			tx.Rollback()
			s.HandlerError(w, r)
			return
		}

		log.Println("end")
	})
}

func (s TxServer) HandlerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("transaction error"))
}
