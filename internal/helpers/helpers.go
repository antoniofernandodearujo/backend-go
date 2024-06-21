package helpers

import (
	"challenge-api/internal/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Envelop map[string]interface{}

type Message struct {
	InfoLog *log.Logger
	ErrorLog  *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var MessageLogs = Message{
	InfoLog: infoLog,
	ErrorLog:  errorLog,
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxByte := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != nil {
		return errors.New("body must only have a single JSON value")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.WriteHeader(status)
	if data == nil {
		return nil
	}
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func ErrorJSON ( w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload services.JsonResponseModel
	payload.Error = true
	payload.Message = err.Error()
	WriteJSON(w, statusCode, payload, nil)
}