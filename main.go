package main

import (
	"XmenChallenge/app"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_MACHINE"))

	//a.Run(":8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/mutant", a.IsMutant)
	http.HandleFunc("/stats", a.Stats)
	port := "8080"
	log.Printf("Defaulting to port %s", port)
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Print("Index call success")
	fmt.Fprint(w, "Hello, Magneto!")
}
