package main

import (
    "fmt"
    "net/http"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

    route.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":8080", route)


    // http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Welcome to my website!")
    // })

    // fs := http.FileServer(http.Dir("static/"))
    // http.Handle("/static/", http.StripPrefix("/static/", fs))

    // http.ListenAndServe(":80", nil)
}