package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/crocodiles128/go-api/server/routes"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes.ConfigRoutes(s.server)

	log.Println("Servidor rodando na porta: " + s.port)
	log.Fatal(s.server.Run(":" + s.port))
}