package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sapirrior/blink/internal/engine"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please include an url to open")
		os.Exit(1)
	}

	var url = os.Args[1]
	err := engine.Open(url)
	if err != nil {
		log.Fatal("blink:", err)
		os.Exit(1)
	}

	fmt.Println("Opening Link...")
}
