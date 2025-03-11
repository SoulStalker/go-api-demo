package main

import (
	"fmt"
	"go/api-demo/configs"
	"go/api-demo/internal/auth"
	"go/api-demo/internal/auth/link"
	"go/api-demo/pkg/db"
	"go/api-demo/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})


	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}
