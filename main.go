package main

import (
	"log"

	"github.com/narendrayogi/init-go/a"
	"github.com/narendrayogi/init-go/z"
)

func init() {
	log.Println("Main initialisation")
}

func main() {
	log.Println("Executing main")
	log.Println(a.SomeGlobalVariable + " " + z.SomeGlobalVariable)
}
