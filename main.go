package main

import (
	"log"

	"gopkg.in/alecthomas/kingpin.v2"
)

const ()

var (
	debug = kingpin.Flag("debug", "Set debug mode").Short('d').Default("false").Bool()
)

func main() {

	kingpin.Parse()

	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	err = Run()
	if err != nil {
		log.Fatal(err)
	}
	End()
}
