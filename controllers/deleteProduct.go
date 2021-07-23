package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"graphApp/db"
	"graphApp/global"
	"net/http"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		global.NewWebError(w, errors.New("invalid id of product"), http.StatusBadRequest)
		return
	}

	err := db.DeleteP(id)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	respJs, err := json.Marshal(resp{
		Err: false,
		Msg: "Deleted product successfully!",
	})

	sendResp(w, http.StatusOK, respJs)

}
