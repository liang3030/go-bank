package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(s Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("account number =>", acc.Number)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "seedFN3", "seedLN3", "seedPW3")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	fmt.Println("hello, go bank start")
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", store)

	// TODO: why need a * in seed - because the return value is address
	if *seed {
		fmt.Println("seed the db")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
