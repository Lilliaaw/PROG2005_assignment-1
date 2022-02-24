package handler

import "net/http"

/*
Empty handler as default handler
*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No functionality on root level. Please use paths "+UNIINFO_PATH+", "+NEIGHBOURUNIS_PATH+" or "+DIAG_PATH+".", http.StatusNotFound)
}
