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
	data := struct {
		Api string `json:"api"`
	}{"ok"}

	returnEncodedData(w, data)
}

func (f *Routes) Ad(w http.ResponseWriter, r *http.Request) {
	data, err := ga.AdData()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnEncodedData(w, data)
}

func (f *Routes) Origin(w http.ResponseWriter, r *http.Request) {
	data, err := ga.OriginData()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnEncodedData(w, data)
}

// return JSON encoded data to client
func returnEncodedData(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
