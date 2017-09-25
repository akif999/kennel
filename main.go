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

	app := NewApp()

	err := app.Init()
	if err != nil {
		app.End()
		log.Fatal(err)
	}
	err = app.Run()
	if err != nil {
		app.End()
		log.Fatal(err)
	}
	app.End()
}
