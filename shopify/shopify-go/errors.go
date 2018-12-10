package shopifygo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	Errors string `json:"errors"`
}

func (err ApiError) Error() string {
	if err == (ApiError{}) {
		return ""
	}
	return fmt.Sprintf("Shopify: %s", err.Errors)
}

// relevantError returns any non-nil http-related error (creating the request, getting
// the response, decoding) if any. If the response was a success, it returns nil.
// If the decoded apiError is non-nil the apiError is returned. Otherwise, an attempt
// is made to decode the response in to all other possible error types. Finally,
// a generic error is returned if needed.
func relevantError(resp *http.Response, httpError error, errorDetails *json.RawMessage, errorTypes ...error) error {
	if httpError != nil {
		return httpError
	}

	if 200 <= resp.StatusCode && resp.StatusCode <= 299 {
		return nil
	}

	apiError := new(ApiError)
	if json.Unmarshal(*errorDetails, apiError) == nil {
		return apiError
	}

	for _, errorType := range errorTypes {
		if json.Unmarshal(*errorDetails, errorType) == nil {
			return errorType
		}
	}

	return fmt.Errorf("Unknown API error encountered. Status Code: %d", resp.StatusCode)
}
