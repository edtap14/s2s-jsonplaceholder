package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	UserName string `json:"name"`
	Email    string `json:"email"`
}

var usuarios = map[string]int{
	"Youstark":  1,
	"Alejandro": 2,
	"Ethien":    3,
	"Edgar":     4,
}

func main() {
	port := ":3000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var u user
			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				log.Fatal(err)
			}
			result := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos?userId=%d", usuarios[u.UserName])
			fmt.Fprintf(w, "%s\n", result)
		}

	})

	fmt.Printf("Server is running at %s 🚀", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
