package main

import (
	"github.com/joho/godotenv"
	"github.com/julianh99/goexpenses/store"
	"os"
)

func main() {

	godotenv.Load()

	storeOptions := store.PostgresOptions{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
	}

	store := store.CreatePostgresStore(storeOptions)
	server := NewApiServer("localhost", 8080, store)

	server.Run()

}
