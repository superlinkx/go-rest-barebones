package encoder

import (
	"encoding/json"
	"net/http"
)

type JSONError struct {
	Error string `json:"error"`
}

// WriteJSONResponse writes a JSON response on the response writer
func WriteJSONResponse(response http.ResponseWriter, statusCode int, data any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(data)
}

func WriteJSONError(response http.ResponseWriter, statusCode int, err error) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(JSONError{Error: err.Error()})
}
