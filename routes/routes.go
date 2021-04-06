package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/quintsys/ga-exporter/ga"
)

type Routes struct {
}

func (f *Routes) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ok := struct {
		Api string `json:"api,omitempty"`
	}{"ok"}
	json.NewEncoder(w).Encode(ok)
}

func (f *Routes) Ad(w http.ResponseWriter, r *http.Request) {
	data, err := ga.AdData()
	if err != nil {
		log.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (f *Routes) Origin(w http.ResponseWriter, r *http.Request) {
	data, err := ga.OriginData()
	if err != nil {
		log.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
