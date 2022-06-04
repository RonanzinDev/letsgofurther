package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Api version , later this will be automaticalyy
const version = "1.0.0"

//This struct will hold on all the config settings for our application
type config struct {
	port int
	env  string
}

// Define an application struct to hold the dependencies for our HTTP handlers, helpers,
// and middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	// with this we can use flags like -port=3000 or any port that you want
	flag.IntVar(&cfg.port, "port", 4000, "Api server port")
	// with this we can use flags like -env=production
	flag.StringVar(&cfg.env, "env", "development", "Enviroment(development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//Start http server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
