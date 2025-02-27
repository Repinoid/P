package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	pgx "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	//create router
	router := mux.NewRouter()
	router.HandleFunc("/f", first).Methods("GET")
	router.HandleFunc("/s", second).Methods("GET")
	router.HandleFunc("/", func(rwr http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rwr, "Ruoter %+v\n", router)
	}).Methods("GET")

	//start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
func first(rwr http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(rwr, "cant Open sql %v\n", err)
		return
	}
	defer db.Close()

	fmt.Fprintf(rwr, "ok Open db %+v %+v\n", db, err)

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(rwr, "NO PING %v\n", err)
		return
	}
	fmt.Fprintf(rwr, "PING OK %v %v\n", err, time.Now())

}
func second(rwr http.ResponseWriter, req *http.Request) {
	//	var DBEndPoint = "postgres://postgres:postgres@go_db:5432/postgres"

	//	baza, err := pgx.Connect(context.Background(), DBEndPoint)
	baza, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(rwr, "NO pgx.Connect %v\n", err)
		return
	}
	fmt.Fprintf(rwr, "PING OK %v %v\n", baza, time.Now())

}
