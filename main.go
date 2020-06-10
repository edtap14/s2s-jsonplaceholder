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
	Prettify bool   `json:"prettify,omitempty"`
}
type toDoList struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
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

			//Obtención de tareas
			res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos?userId=%d", usuarios[u.UserName]))
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			var toDo []toDoList
			err2 := json.NewDecoder(res.Body).Decode(&toDo)
			if err != nil {
				log.Fatal(err2)
			}

			if u.Prettify == false {
				js, err := json.Marshal(toDo) // send json compressed
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			} else {
				js, err := json.MarshalIndent(toDo, "", "\t")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			}

		}

	})
	fmt.Printf("Server is running at %s 🚀 \n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
