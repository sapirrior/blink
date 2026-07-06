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
	}

	var url = os.Args[1]
	err := engine.Open(url)
	fmt.Println("Opening Link...")
	
	if err != nil {
		log.Fatal("blink:", err)
	}
}
