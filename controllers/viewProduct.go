package controllers

import (
	"encoding/json"
	"graphApp/db"
	"graphApp/global"
	"io/ioutil"
	"net/http"
)

type vP struct {
	ProductId  string `json:"productId"`
	CustomerId string `json:"customerId"`
	Date       string `json:"date"`
}

func ViewProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	var vp vP

	err = json.Unmarshal(body, &vp)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	err = db.ViewP(vp.CustomerId, vp.ProductId, vp.Date)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	respJs, err := json.Marshal(resp{
		Err: false,
		Msg: "Added relation!(view)",
	})
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	sendResp(w, http.StatusCreated, respJs)

}
