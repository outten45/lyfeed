package main

import (
	"database/sql"
	"fmt"
	"log"
	"lyfeed/importer"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func myHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", ps.ByName("id"))
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Made it!")
	importer.Run([]string{"http://blog.outten.net/index.xml"})
	fmt.Println("----------")

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
