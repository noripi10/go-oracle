package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/noripi10/go-oracle/libs"
	_ "github.com/noripi10/go-oracle/libs"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error read .env")
	}

	libs.Query()

	// fmt.Println("Press Enter Key to exit ...")
	// fmt.Scanln()
}
