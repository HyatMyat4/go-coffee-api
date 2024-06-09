package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/takumi/coffee-api/services"
)

type Envlop map[string]interface{}

type Message struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Ltime)

var MessageLogs = &Message{
	InfoLog:  infoLog,
	ErrorLog: errorLog,
}

func ReadJSON(w http.ResponseWriter, req *http.Request, data interface{}) error {
	maxByte := 104857

	req.Body = http.MaxBytesReader(w, req.Body, int64(maxByte))

	decode := json.NewDecoder(req.Body)

	err := decode.Decode(data)

	if err != nil {
		return err
	}

	err = decode.Decode(&struct{}{})

	if err != nil {
		return errors.New("body must have only a single json object")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil

}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload services.JsonResponse

	payload.Error = true

	payload.Message = err.Error()

	WriteJSON(w, statusCode, payload)

}
