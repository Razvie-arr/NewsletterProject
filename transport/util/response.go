package util

import (
	"encoding/json"
	"net/http"
)

func WriteErrResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(body)
	w.Write(b)
}

func WriteResponseWithJsonBody(w http.ResponseWriter, statusCode int, body any) {
	responseBytes, err := json.Marshal(body)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, "Error marshalling tokens to JSON while writing response: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseBytes)
}
