package main

import (
	"log"
	"vestibulum/blog/bootstrap"
)

func main() {

	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(":5000"))

}
