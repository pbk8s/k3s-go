package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
    // MAIN SECTION HTML CODE
    fmt.Fprintf(w, "<h1>Whoa, Go is neat!!</h1>")
    fmt.Fprintf(w, "<title>Go</title>")
    fmt.Fprintf(w, "<img src='images/APMC.jpg' alt='apmc' style='width:235px;height:320px;'>")
}

func main() {
    http.HandleFunc("/", index_handler)
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Welcome to our demo. Hope it's fun for you!!!!")
//	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health OK!")
	})

	http.ListenAndServe(":8080", nil)
}
