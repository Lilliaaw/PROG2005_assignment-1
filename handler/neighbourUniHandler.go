package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

/*
Entry point handler for collection information
*/
func NeighbourHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		neighhandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only POST and GET are supported.", http.StatusNotImplemented)
		return
	}

}

/*
Dedicated handler for GET requests
*/
func neighhandleGetRequest(w http.ResponseWriter, r *http.Request) {
	url := strings.ReplaceAll(r.URL.Path, " ", "%20")
	urlParts := strings.Split(url, "/")

	//checks if valid url
	if len(urlParts) > 5 {
		inURL := countryURL + "name/" + urlParts[4]
		nameUni := urlParts[5] //saves uniName search

		//checks if ?limit is in the url and saves the value
		limitString := strings.Split(r.URL.RawQuery, "limit=")
		limit := 0
		if len(limitString) > 1 {
			limit, _ = strconv.Atoi(limitString[1])
		}

		// Write content type header (best practice)
		w.Header().Add("content-type", "application/json")

		neighbourUnis := GetNeighbourUnis(inURL, nameUni, limit) //gets the output structs

		// Instantiate encoder
		encoder := json.NewEncoder(w)

		// Encode university structs
		err := encoder.Encode(neighbourUnis)
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
			return
		}

		// Explicit specification of return status code --> will default to 200 if not provided.
		http.Error(w, "", http.StatusOK)
	} else {
		//url is missing mandatory components
		http.Error(w, "No functionality without search. Please use "+NEIGHBOURUNIS_PATH+"{country}/{partial_or_complete_university_name}{?limit={number}}.", http.StatusNotFound)
	}
}
