package namefinder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/health"
	"log"
)

type Server struct {
	config *ServerConfig
}

func NewServer(config *ServerConfig) *Server {
	return &Server{config: config}
}

type ServerConfig struct {
	DNSHost   string
	DNSPort   string
	AdminHost string
	AdminPort string
}

func (s *Server) Start() {
	router := gin.Default()

	health.NewCheckHandler(router).Register()

	err := router.Run(fmt.Sprintf("%s:%s", s.config.AdminHost, s.config.AdminPort))

	if err != nil {
		log.Fatal("error starting admin server: ", err)
	}
}
