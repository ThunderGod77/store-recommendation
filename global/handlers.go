package global

import (
	"encoding/json"
	"log"
	"net/http"
)

type WebError struct {
	Msg        string `json:"msg"`
	Err        bool   `json:"err"`
	statusCode int
}

func (we WebError) ReturnError(w http.ResponseWriter) {
	log.Println(we.Msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(we.statusCode)
	jE, _ := json.Marshal(we)
	_, err := w.Write(jE)
	if err != nil {
		log.Println(err)
	}

}

func NewWebError(w http.ResponseWriter, err error, statusCode int) {
	we := WebError{
		Msg:        err.Error(),
		Err:        false,
		statusCode: statusCode,
	}
	we.ReturnError(w)
}
