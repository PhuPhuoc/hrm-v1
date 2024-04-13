package common

import (
	"encoding/json"
	"net/http"
)

// handle parse from json to model(payload) when recieve data from body of request(http)
func ParseJson(r *http.Request, payload any) error {
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return err
	}
	return nil
}

// handle write json when response to client
func WriteJSON(w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json")
	if res, ok := response.(error_response); ok {
		w.WriteHeader(res.StatusCode)
	}

	if _, ok := response.(success_response); ok {
		w.WriteHeader(http.StatusOK)
	}
	return json.NewEncoder(w).Encode(response)
}
