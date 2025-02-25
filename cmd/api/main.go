package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	// Import the pq driver so that it can register itself with the database/sql package
	_ "github.com/lib/pq"
	"greenlight.wolfheros.com/internal/data"
)

const version = "1.0.0"

// 
// Define a config strcut holding all the application configuration, 
// Add a DB struct field to hold the configuration settings
//
type config struct {
	port int
	env  string
	db struct{
		dsn string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime	time.Duration	
	}
}

// Define an application struct which will hold
// dependencies for HTTPhandlers, helpers, and middleware
type application struct {
	config config
	logger *slog.Logger
	models data.Models
}

func main() {
	var cfg config

	// Read value from command line while starting the application.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment(development|staging|production)")

	// Read [DSN] value from the db-dsn command-line flag
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	// Reas the Database connection pool settings from command-line flags
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15 * time.Minute, "PostgreSQL max connction idle time")
	// Reading all the input value frome commander line
	flag.Parse()

	//Initial a structed logger which write log entries to the standard out steam.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


	// create db connection pool
	db, err := openDB(cfg)
	if err!= nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Defer a call to db.Close() so that the connection pool is closed before the main() function exits
	defer db.Close()
	
	// Logging a message to say the connection pool has been successfully establised
	logger.Info("database connection pool established")

	//Use [data.NewModels()] function to initialize a Models struct.
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	//Declare a new servermux
	// add a /vi/healthcheck route

	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Declare a Http server listen on the port provide in the config
	// usse servermux as the handler
	// contain, time out, and log message
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	//start the HTTP server
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	// because the err already declared, so here change from := to =
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

// [openDB()] funtion return a [sql.DB] connection pool
func openDB(cfg config)(*sql.DB, error){

	// Use [sql.Open()] to create an empty connection pool,
	// using the DSN from the config struct

	db, err:=sql.Open("postgres", cfg.db.dsn)
	if err!=nil {
		return nil, err
	}

	// Set maximum number of open connections
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	// Set maximum time of idle connections in the pool
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)
	// Set maximum number of idle connections in the pool;
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Create a context with a 5 seconds timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Use [PingContext[]] to establish a new connection to the database,
	// If the connection couldn't be established successfully within the 5 second deadline.
	// it will return an error.
	err = db.PingContext(ctx)
	if err!=nil {
		db.Close()
		return nil, err
	}

	// return the db connection pool
	return db, nil
}
