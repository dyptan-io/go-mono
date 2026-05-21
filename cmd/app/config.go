package main

import (
	"flag"
)

// Config contains server configuration for the service.
type Config struct {
	BindAddr string
}

func readConfig() Config {
	config := Config{}

	flag.StringVar(&config.BindAddr, "addr", ":8080", "an address for HTTP server listener")
	flag.Parse()

	return config
}
