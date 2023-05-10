package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Application server port")
	flag.Parse()

	logger := log.New(os.Stdout, "[dev] ", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
		config: cfg,
	}

	appPort := fmt.Sprintf(":%d", cfg.port) // convert port (int) to string with colon for serving

	httpServer := &http.Server{
		Addr:    appPort,
		Handler: app.route(),
	}

	logger.Printf("starting http server on localhost%s", appPort)
	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Fatal("http server could not start:", err)
	}
}
