package tools

import (
	"net/http"
	"encoding/json"
)

func JsonResponse(response interface{}, w http.ResponseWriter) {
	result, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
