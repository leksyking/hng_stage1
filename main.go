package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type Desription struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           uint   `json:"age"`
	Bio           string `json:"bio"`
}

func main() {
	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //all
		AllowedMethods:   []string{http.MethodPost, http.MethodGet},
		AllowedHeaders:   []string{"*"}, //all
		AllowCredentials: false,         //none
	})

	mux.HandleFunc("/", greetings)
	handler := c.Handler(mux)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server started on port", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
func greetings(w http.ResponseWriter, r *http.Request) {
	description := Desription{
		SlackUsername: "leksyking",
		Backend:       true,
		Age:           22,
		Bio:           "I am a very hardworking and cooperative engineer always ready to learn from people around me and learn new technologies to help develop myself and softwares for people.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(description)
}
