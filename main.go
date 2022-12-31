package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abibby/buzzer/controllers"
	"github.com/abibby/buzzer/ui"
	"github.com/abibby/fileserver"
	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/room", controllers.JoinRoom).Methods("GET")

	r.PathPrefix("/").
		Handler(fileserver.WithFallback(ui.Content, "dist", "index.html", nil)).
		Methods("GET")

	log.Print("Listening on http://localhost:3335")
	http.ListenAndServe(":3335", r)
}
