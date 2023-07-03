package main

import (
	"log"

	"github.com/cloudbees-test/simple-go-api-app/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
