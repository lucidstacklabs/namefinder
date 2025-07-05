package main

import (
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder"
	"github.com/lucidstacklabs/namefinder/internal/pkg/env"
)

func main() {
	namefinder.NewServer(&namefinder.ServerConfig{
		DNSHost:   env.GetOrDefault("DNS_HOST", "0.0.0.0"),
		DNSPort:   env.GetOrDefault("DNS_PORT", "5353"),
		AdminHost: env.GetOrDefault("ADMIN_HOST", "0.0.0.0"),
		AdminPort: env.GetOrDefault("ADMIN_PORT", "5354"),
	}).Start()
}
