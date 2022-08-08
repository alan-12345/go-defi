package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	start_bot()
}
