package main

import (
	"fmt"
	"lyfeed/importer"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func myHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %v\n", ps.ByName("id"))
}

func main() {
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
