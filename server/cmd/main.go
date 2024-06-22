package main

import (
	"log"

	"github.com/e-lua/go-nextjs-ts-chat/db"
	"github.com/e-lua/go-nextjs-ts-chat/internal/user"
	"github.com/e-lua/go-nextjs-ts-chat/internal/ws"
	"github.com/e-lua/go-nextjs-ts-chat/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initiliaze database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
