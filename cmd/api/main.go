package main

import (
	"flag"
	"log/slog"
)

const version = "1.0.0"

// Define a config strcut holding all the application configuration
type config struct {
	port int
	env  string
}

// Define an application struct which will hold
// dependencies for HTTP handlers, helpers, and middleware
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production)")
	flag.Parse()

	//Initial logger

}
