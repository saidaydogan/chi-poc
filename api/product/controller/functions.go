package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func getUrlParamInt(r *http.Request, param string) int {
	urlParam := chi.URLParam(r, param)
	res, _ := strconv.Atoi(urlParam)
	return res
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return errors message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

func respondWithErrors(w http.ResponseWriter, code int, errors []string) {
	respondwithJSON(w, code, map[string][]string{"errors": errors})
}
