package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/twpayne/go-greet"
)

var (
	name = flag.String("name", "", "name to greet")
)

func run() error {
	flag.Parse()
	_, err := fmt.Println(greet.Greet(*name))
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
