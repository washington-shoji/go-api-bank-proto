package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", store)

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
