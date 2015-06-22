package main

import (
	"database/sql"
	"fmt"
	"log"
	"lyfeed"
	"lyfeed/importer"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gchaincl/dotsql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

var dot *dotsql.DotSql
var db *sql.DB

func myHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", ps.ByName("id"))
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dot, err = dotsql.LoadFromFile("queries.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dot>> %+v\n", dot)

	context := &lyfeed.Context{DB: db, DotSql: dot}

	importer.Run(context, []string{"http://blog.outten.net/index.xml"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	mux := httprouter.New()
	mux.GET("/page/:id", myHandler)
	n := negroni.Classic()
	// n := negroni.New()
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(":%s", port))

}
