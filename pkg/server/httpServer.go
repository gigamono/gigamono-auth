package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-auth/internal/rest"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gin-gonic/gin"
)

func (server *AuthServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.Types.Auth.PublicPort),
	)
	if err != nil {
		return err
	}

	server.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: server,
	}

	return httpServer.Serve(listener)
}

func (server *AuthServer) setRoutes() {
	// Add middlewares.
	server.Use(gin.CustomRecovery(middleware.PanicHandler))

	v1Group := server.Group("/rest/v1")
	rest.V1Delegate(v1Group, &server.App)
}
