package controllers

import (
	"encoding/json"
	"fmt"
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

func cpRelation(actionString string, actionFunc func(string, string, string) error) func(http.ResponseWriter, *http.Request) {

	cpr := func(w http.ResponseWriter, r *http.Request) {
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

		err = actionFunc(vp.CustomerId, vp.ProductId, vp.Date)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}

		respJs, err := json.Marshal(resp{
			Err: false,
			Msg: fmt.Sprintf("Added relation!(%s)", actionString),
		})
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}

		sendResp(w, http.StatusCreated, respJs)

	}
	return cpr
}

var ViewProduct = cpRelation("view", db.ViewP)
var OrderProduct = cpRelation("order", db.OrderP)
var WishlistProduct = cpRelation("wishlist", db.WishlistP)
