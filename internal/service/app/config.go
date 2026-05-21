package app

import "flag"

// Config holds runtime configuration for the app service.
type Config struct {
	BindAddr string
}

// ReadConfig parses configuration from CLI flags.
func ReadConfig() Config {
	cfg := Config{BindAddr: ""}

	flag.StringVar(&cfg.BindAddr, "addr", ":8080", "an address for HTTP server listener")
	flag.Parse()

	return cfg
}
