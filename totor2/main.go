package main

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)
func main() {
	r:=mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]
		
		b := []byte("<h1>Book:" + title + "</h1><br><h2>: " + page + "</h2>")
	w.Write(b)
        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/cal/{num1}/plus/{num2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		num1, _ := strconv.Atoi(vars["num1"])
		num2, _ := strconv.Atoi(vars["num2"])
		fmt.Fprintf(w, "<h1>Total = %d</h1>", num1+num2)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	})
	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "woramet")
	})

fs:=http.FileServer(http.Dir("static/"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
	
http.ListenAndServe(":8080", r)
}