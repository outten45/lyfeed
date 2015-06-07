package main

import (
	"fmt"
	"lyfeed/importer"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func myHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", ps.ByName("id"))
}

func main() {
	fmt.Println("Made it!")
	importer.Run()
	fmt.Println("----------")

	mux := httprouter.New()
	mux.GET("/page/:id", myHandler)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")

}