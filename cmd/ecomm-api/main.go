package main

import (
	"log"

	"github.com/ianschenck/envflag"
	"github.com/sir-radar/go-ecom/db"
	"github.com/sir-radar/go-ecom/ecomm-api/handler"
	"github.com/sir-radar/go-ecom/ecomm-api/server"
	"github.com/sir-radar/go-ecom/ecomm-api/storer"
)

func main() {
	const minSecretKeySize = 32
	var secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")
	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("secret key must be at least %d characters long", minSecretKeySize)
	}
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
