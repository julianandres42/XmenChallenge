package main

import (
	"XmenChallenge/app"
	"os"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	a.Run(":8080")
}
