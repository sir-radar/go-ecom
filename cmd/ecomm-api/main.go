package main

import (
	"log"

	"github.com/sir-radar/go-ecom/db"
)

func main(){
	db, err := db.NewDatabase()
	if err != nil{
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")
}
