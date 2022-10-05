package decoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeJSONRequest(request *http.Request, data any) error {
	if request.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("content type is not application/json")
	} else if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		return fmt.Errorf("error decoding JSON body: %w", err)
	} else {
		return nil
	}
}
