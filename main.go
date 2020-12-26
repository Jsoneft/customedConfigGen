package main

import (
	"github.com/Jsoneft/customedConfigGen/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Exec err = %v", err)
	}
}
