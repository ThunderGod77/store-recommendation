package controllers

import (
	"encoding/json"
	"graphApp/db"
	"graphApp/global"
	"io/ioutil"
	"net/http"
)

func WishlistProduct(w http.ResponseWriter, r *http.Request) {
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

	err = db.WishlistP(vp.CustomerId, vp.ProductId, vp.Date)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	respJs, err := json.Marshal(resp{
		Err: false,
		Msg: "Added relation!(wishlist)",
	})
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	sendResp(w, http.StatusCreated, respJs)

}
