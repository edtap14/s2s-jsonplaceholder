package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "🌱\n")
}

func main() {
	port:= ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/borrego", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Borrego Viudo 🐏\n")
	})

	fmt.Printf("Server is running at %s 🚀",port)
	err := http.ListenAndServe(port, mux)
	if err != nil{
		log.Fatal(err)
	}
}