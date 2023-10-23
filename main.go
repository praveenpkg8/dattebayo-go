package main

import (
	"log"

	"github.com/praveenpkg8/dattebayo-go/database"
	"github.com/praveenpkg8/dattebayo-go/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	database.Init()
	server.Init()
	// scripts.PopulateDB()
}
