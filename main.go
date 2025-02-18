package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/u", first).Methods("GET")
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

	fmt.Fprintf(rwr, "ok Open %v", err)

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(rwr, "NO PING %v\n", err)
		return
	}
	fmt.Fprintf(rwr, "PING OK %v %v\n", err, time.Now())

}
