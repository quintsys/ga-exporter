package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/quintsys/ga-exporter/env"
	m "github.com/quintsys/ga-exporter/middleware"
	"github.com/quintsys/ga-exporter/routes"
)

// App contains application data
type App struct {
	Port string
	routes.Routes
}

// Start sets the handlers and starts the server
func (a *App) Start() {
	http.HandleFunc("/",
		m.ChainMiddlewares(a.Index, m.Logger, m.ContentTypeJSON))

	middlewares := []m.Middleware{
		m.Logger,
		m.PostOnly,
		m.ContentTypeJSON,
	}
	http.HandleFunc("/ad",
		m.ChainMiddlewares(a.Ad, middlewares...))
	http.HandleFunc("/origin",
		m.ChainMiddlewares(a.Origin, middlewares...))

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
