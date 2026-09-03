package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	// Set the Content-Type header to indicate that the response body contains JSON data.
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code for the response using w.WriteHeader(statusCode). This allows the client to understand the outcome of the request (e.g., success, error, etc.).
	w.WriteHeader(statusCode)

	// Use json.NewEncoder(w).Encode(data) to encode the provided data as JSON and write it to the response body.
	return json.NewEncoder(w).Encode(data)
}

func GetErrorResponse(err error, statusCode int) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
		Code:    statusCode,
	}
}

// ValidationErrorResponse takes a validator.ValidationErrors object and constructs an ErrorResponse containing the validation error messages and the corresponding HTTP status code (400 Bad Request).
func ValidationErrorResponse(errs validator.ValidationErrors) ErrorResponse {
	var errorMessages []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("Field %s is required", err.Field()))
		}
	}

	return ErrorResponse{
		Message: strings.Join(errorMessages, ","),
		Code:    http.StatusBadRequest,
	}

}