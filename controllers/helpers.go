package controllers

import "net/http"

func sendResp(w http.ResponseWriter, statusCode int, respJs []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respJs)
}
