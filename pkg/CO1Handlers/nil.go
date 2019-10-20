package CO1Handlers

import "net/http"

//This is the default handler for invalid requests
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid Request", http.StatusBadRequest)
}