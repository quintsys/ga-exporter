package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/quintsys/ga-exporter/env"
	"github.com/quintsys/ga-exporter/middleware"
	"github.com/quintsys/ga-exporter/routes"
)

// App contains application data
type App struct {
	Port string
	routes.Routes
}

// Start sets the handlers and starts the server
func (a *App) Start() {
	http.Handle("/", middleware.LogWrap(a.Index))
	http.Handle("/ad", middleware.LogWrap(a.Ad))
	http.Handle("/origin", middleware.LogWrap(a.Origin))

	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting ga-exporter on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	server := App{
		Port: env.Lookup("PORT", "8080"),
	}
	server.Start()
}
