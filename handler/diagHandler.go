package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
Entry point handler for Location information
*/
func DiagHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		diaghandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET or POST are supported.", http.StatusNotImplemented)
		return
	}

}

/*
Dedicated handler for GET requests
*/
func diaghandleGetRequest(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")

	//Instantiate the client
	client := &http.Client{}

	//Issue request
	uniRes, err := client.Get(uniURL)
	if err != nil {
		log.Fatal("Error in response:", err)
	}
	//Issue request
	countryRes, err := client.Get(countryURL + "/all")
	if err != nil {
		log.Fatal("Error in response:", err)
	}

	//saves the diagnostics in output struct
	diagnostics := Diagnostic{
		Statusuniapi:     uniRes.StatusCode,
		Statuscountryapi: countryRes.StatusCode,
		Version:          urlParts[2],
		Duration:         fmt.Sprint(GetDuration()),
	}

	// Write content type header (best practice)
	w.Header().Add("content-type", "application/json")

	// Instantiate encoder
	encoder := json.NewEncoder(w)

	// Encode diagnostics struct
	err = encoder.Encode(diagnostics)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}
}
