package main

import (
	"fmt"
	"go/api-demo/configs"
	"go/api-demo/internal/auth"
	"go/api-demo/internal/link"
	"go/api-demo/internal/stat"
	"go/api-demo/internal/user"
	"go/api-demo/pkg/db"
	"go/api-demo/pkg/event"
	"go/api-demo/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus: eventBus,
		StatRepository: statRepository,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus: eventBus,
		StatService: statService,
		// StatRepository: statRepository,
		Config:         conf,
	})
	stat.NewStatHandler(router, &stat.StatHandlerDeps{
		StatRepository: statRepository,
	})
	
	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}
