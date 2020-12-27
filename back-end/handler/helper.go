package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/Bharathstark/lourde/back-end/logger"
)

type malformedRequest struct {
	status int
	msg    string
}

//ResponseMessage is
type ResponseMessage struct {
	Status int    `json:"response_code"`
	Msg    string `json:"message"`
}

//Response is
type Response struct {
	Status interface{} `json:"response_status"`
	Data   interface{} `json:"data"`
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) interface{} {

	if r.Header.Get("Content-Type") != "" {
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			msg := "Content-Type header is not application/json"
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}
		}
	}
	body := http.MaxBytesReader(w, r.Body, 1*MB)
	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &ResponseMessage{Status: http.StatusUnsupportedMediaType, Msg: msg}
	}

	return nil
}

//ConstructResponse is used to Construct a Response
func ConstructResponse(data interface{}, messageType string) *Response {
	if (messageType == "failed" || messageType == "custom") && data != nil {
		response := Response{
			Status: data,
			Data:   nil,
		}
		return &response
	}

	if messageType == "failed" {
		response := Response{
			Status: ResponseMessage{
				Status: http.StatusBadRequest,
				Msg:    "Failed",
			},
			Data: data,
		}
		return &response
	}
	response := Response{
		Status: ResponseMessage{
			Status: http.StatusOK,
			Msg:    "success",
		},
		Data: data,
	}

	return &response

}

func uploadFile(entity *string, multipart *multipart.FileHeader) (string, *error) {

	file, err := multipart.Open()
	defer file.Close()
	if err != nil {
		return "", &err
	}
	//create destination file making sure the path is writeable.
	entityPath := filepath.Join("img", *entity, multipart.Filename)
	path := filepath.Join("dist", entityPath)
	dir := filepath.Dir(path)
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	dst, err := os.Create(path)

	defer dst.Close()
	if err != nil {
		return path, &err
	}

	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		return "", &err
	}
	logger.Log.Printf("file has been created and stored -" + path)
	return filepath.ToSlash(entityPath), nil

}

func deleteFile(entitypath *string) {
	path := filepath.Join("dist", *entitypath)
	_, err := os.Stat(path)
	if err == nil {
		e := os.Remove(path)
		if e != nil {
			logger.Log.Panicf("Error Deleting the file" + path)
		}
	}
}

func validateURL(w http.ResponseWriter, r *http.Request, dst interface{}, id string) bool {
	rows, error := database.GetEntity(dst, &id)
	if *rows <= 0 || *error != nil {
		response := ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
		return false
	}
	return true

}
