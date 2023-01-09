package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	isuconpv1 "isuconp/gen/isuconp/v1"
	"isuconp/gen/isuconp/v1/isuconpv1connect"

	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type IsuconpServer struct{}

func (s *IsuconpServer) Initialize(
	ctx context.Context,
	req *connect.Request[isuconpv1.InitializeRequest],
) (*connect.Response[isuconpv1.InitializeResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&isuconpv1.InitializeResponse{})
	res.Header().Set("Isuconp-Version", "v1")

	// Do
	dbInitialize()

	return res, nil
}

func (s *IsuconpServer) GetLogin(
	ctx context.Context,
	req *connect.Request[isuconpv1.GetLoginRequest],
) (*connect.Response[isuconpv1.GetLoginResponse], error) {
	log.Println("Request headers: ", req.Header())

	res := connect.NewResponse(&isuconpv1.GetLoginResponse{
		Id:                "1",
		Name:              "test",
		Passhash:          "passhash",
		Authority:         1,
		DelFlg:            false,
		CreatedAtUnixNano: time.Now().UnixNano(),
	})
	res.Header().Set("Isuconp-Version", "v1")

	//ID          int       `db:"id"`
	//AccountName string    `db:"account_name"`
	//Passhash    string    `db:"passhash"`
	//Authority   int       `db:"authority"`
	//DelFlg      int       `db:"del_flg"`
	//CreatedAt   time.Time `db:"created_at"`
	//
	return res, nil
}

func main() {
	host := os.Getenv("ISUCONP_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("ISUCONP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUCONP_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUCONP_DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("ISUCONP_DB_PASSWORD")
	dbname := os.Getenv("ISUCONP_DB_NAME")
	if dbname == "" {
		dbname = "isuconp"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()

	server := &IsuconpServer{}
	path, handler := isuconpv1connect.NewIsuconpServiceHandler(server)

	r := chi.NewRouter()

	r.Get("/login", getLogin)
	r.Post("/login", postLogin)
	r.Get("/register", getRegister)
	r.Post("/register", postRegister)
	r.Get("/logout", getLogout)
	r.Get("/", getIndex)
	r.Get("/posts", getPosts)
	r.Get("/posts/{id}", getPostsID)
	r.Post("/", postIndex)
	r.Get("/image/{id}.{ext}", getImage)
	r.Post("/comment", postComment)
	r.Get("/admin/banned", getAdminBanned)
	r.Post("/admin/banned", postAdminBanned)
	r.Get(`/@{accountName:[a-zA-Z]+}`, getAccountName)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("../public")).ServeHTTP(w, r)
	})

	//r.Get("/initialize", getInitialize)
	r.Method(http.MethodGet, "/initialize", connect.NewUnaryHandler(
		"/initialize",
		server.Initialize,
	))

	mux := http.NewServeMux()
	mux.Handle("/", r)
	mux.Handle(path, handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
