package handler

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

/*
Entry point handler for Location information
*/
func InfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		infohandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET or POST are supported.", http.StatusNotImplemented)
		return
	}

}

/*
Dedicated handler for GET requests
*/
func infohandleGetRequest(w http.ResponseWriter, r *http.Request) {
	uniName := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")

	inURL := uniURL + "search"
	//check if it is valid url input
	if uniName != "uniinfo" {
		inURL += "?name=" + uniName

		// Write content type header (best practice)
		w.Header().Add("content-type", "application/json")

		Unis := GetUniInfo(inURL) //gets the output structs

		// Instantiate encoder
		encoder := json.NewEncoder(w)

		// Encode the univesity structs
		err := encoder.Encode(Unis)
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
			return
		}
	} else {
		//url is missing mandatory components
		http.Error(w, "No functionality without search. Please use "+UNIINFO_PATH+"{partial_or_complete_university_name}.", http.StatusOK)
	}
}
