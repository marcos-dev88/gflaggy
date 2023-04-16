package main

import (
	"log"

	"github.com/marcos-dev88/gflaggy"
)

func main() {

	flag := gflaggy.NewFlag("-name")
	flagInt := gflaggy.NewFlag("-age")

	name, err := flag.String()

	if err != nil {
		log.Fatalf("err: %v", err)
	}

	age, err := flagInt.Bool()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	log.Printf("data -> %s | %v", name, age)

}
