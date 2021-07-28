package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"graphApp/db"
	"graphApp/global"
	"net/http"
)

func Recommend(w http.ResponseWriter, r *http.Request) {
	pId := mux.Vars(r)["pId"]
	rType := r.URL.Query().Get("type")
	cId := r.URL.Query().Get("cid")
	if pId == "" || rType == "" {
		global.NewWebError(w, errors.New("invalid id of product"), http.StatusBadRequest)
		return
	}
	if rType == "ab" {
		resp, err := db.AlsoBought(pId, cId)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}
		respJs, err := json.Marshal(resp)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}
		sendResp(w, http.StatusOK, respJs)
		return
	} else if rType == "sc" {
		resp, err := db.SameCategory(pId, cId)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}
		respJs, err := json.Marshal(resp)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}
		sendResp(w, http.StatusOK, respJs)
		return
	} else if rType == "sb" {
		resp, err := db.SameBrand(pId, cId)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}
		respJs, err := json.Marshal(resp)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}
		sendResp(w, http.StatusOK, respJs)
		return
	}

}
