package students

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Farindra968/go_project1/internal/storage"
	"github.com/Farindra968/go_project1/internal/types"
	"github.com/Farindra968/go_project1/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func StudentHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating student.")
		var student types.Student

		err :=json.NewDecoder(r.Body).Decode(&student)
		// Check if the request body is empty
		// errors.is() is used to check if the error returned by json.NewDecoder(r.Body).Decode(&student) is of type io.EOF, which indicates that the request body is empty.
		// io.EOF is a predefined error in the io package that represents the end of a file or stream. In this context, it is used to check if the request body has no content.
		// json.EOF is a predefined error in the encoding/json package that represents the end of a JSON input stream. It is used to check if the JSON decoder has reached the end of the input without finding any valid JSON data.

		if errors.Is(err, io.EOF) { 
			response.WriteJSONResponse(w, http.StatusBadRequest, response.GetErrorResponse(err, http.StatusBadRequest))

		}

		if err != nil {
			response.WriteJSONResponse(w, http.StatusBadRequest, response.GetErrorResponse(err, http.StatusBadRequest))
		}

		// request validation using the validator package
		// validator.New() creates a new instance of the validator, which is used to perform validation on the struct fields based on the validation tags defined in the Student struct.
		// Struct(student) validates the student struct against the validation rules specified in the struct tags (e.g., `validate:"required"`). If any validation rule is violated, an error is returned.
		
		err = validator.New().Struct(student)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			response.WriteJSONResponse(w, http.StatusBadRequest, response.ValidationErrorResponse(validationErrors))
			return
		}

		result, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
			student.Password,
		)

		if err != nil {
			response.WriteJSONResponse(w, http.StatusInternalServerError, response.GetErrorResponse(err, http.StatusInternalServerError))
			return
		}

		response.WriteJSONResponse(w, http.StatusCreated, response.GetSuccessResponse("Student created successfully", http.StatusCreated, result))
		slog.Info("Student created successfully with ID:", "id", result)

	}
}

func GetStudentByID(storage storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Fetching student by ID:", "id", id)

		student, err := storage.GetStudentByID(id)
		if err != nil {
			response.WriteJSONResponse(w, http.StatusInternalServerError, response.GetErrorResponse(err, http.StatusInternalServerError))
			return
		}

		response.WriteJSONResponse(w, http.StatusOK, response.GetSuccessResponse("Student fetched successfully", http.StatusOK, student))
		slog.Info("Student fetched successfully with ID:", "id", id)

	}
}