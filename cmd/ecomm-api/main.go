package main

import (
	"log"

	"github.com/sir-radar/go-ecom/db"
	"github.com/sir-radar/go-ecom/ecomm-api/server"
	"github.com/sir-radar/go-ecom/ecomm-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	_ = server.NewServer(st)
	// hdl := handler.NewHandler(srv)
}
