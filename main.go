package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, firstName string, lastName string, password string) *Account {
	acc, err := NewAccount(firstName, lastName, password)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("new account => ", acc.Number)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Bob", "G", "BlueBerryMan")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", store)

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")

		// Seed account
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
