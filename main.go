package main

import (
	"flag"
	"log"
)

func main() {
	var veces = flag.Int("n", 0, "Insert your name")
	flag.Parse()

	if *nFlag != "" {
		log.Default().Println("Hello, " + *nFlag + "!")
	} else {
		log.Fatal("no name given")
	}
}
