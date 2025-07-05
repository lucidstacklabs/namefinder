package main

import (
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder"
	"github.com/lucidstacklabs/namefinder/internal/pkg/env"
)

func main() {
	namefinder.NewServer(&namefinder.ServerConfig{
		DNSHost:       env.GetOrDefault("DNS_HOST", "0.0.0.0"),
		DNSPort:       env.GetOrDefault("DNS_PORT", "5300"),
		AdminHost:     env.GetOrDefault("ADMIN_HOST", "0.0.0.0"),
		AdminPort:     env.GetOrDefault("ADMIN_PORT", "5301"),
		MongoEndpoint: env.GetOrDefault("MONGO_ENDPOINT", "mongodb://localhost:27017"),
		MongoDatabase: env.GetOrDefault("MONGO_DB", "namefinder"),
	}).Start()
}
