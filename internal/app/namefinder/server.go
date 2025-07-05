package namefinder

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/admin"
	dnsLib "github.com/lucidstacklabs/namefinder/internal/app/namefinder/dns"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/health"
	"github.com/miekg/dns"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Server struct {
	config *ServerConfig
}

func NewServer(config *ServerConfig) *Server {
	return &Server{config: config}
}

type ServerConfig struct {
	DNSHost       string
	DNSPort       string
	AdminHost     string
	AdminPort     string
	MongoEndpoint string
	MongoDatabase string
}

func (s *Server) Start() {
	// Database setup
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(s.config.MongoEndpoint))

	if err != nil {
		log.Fatal("error while connecting to mongo database: ", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("error while pinging mongo database: ", err)
	}

	mongoDatabase := client.Database(s.config.MongoDatabase)

	// Services setup
	admin.NewService(mongoDatabase.Collection("admins"))

	// DNS server setup

	dnsHandler := dnsLib.NewHandler()
	dns.HandleFunc(".", dnsHandler.Handle)

	dnsServer := &dns.Server{
		Addr: fmt.Sprintf("%s:%s", s.config.DNSHost, s.config.DNSPort),
		Net:  "udp",
	}

	go func() {
		log.Printf("starting DNS server on %s:%s", s.config.DNSHost, s.config.DNSPort)
		err = dnsServer.ListenAndServe()

		if err != nil {
			log.Fatal("error while starting DNS server: ", err)
		}
	}()

	// Admin server setup
	router := gin.Default()
	health.NewCheckHandler(router).Register()

	log.Printf("starting admin server on %s:%s", s.config.AdminHost, s.config.AdminPort)

	err = router.Run(fmt.Sprintf("%s:%s", s.config.AdminHost, s.config.AdminPort))

	if err != nil {
		log.Fatal("error starting admin server: ", err)
	}
}
