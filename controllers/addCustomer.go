package controllers

import (
	"encoding/json"
	"graphApp/db"
	"graphApp/global"
	"io/ioutil"
	"net/http"
)

type customer struct {
	Name       string `json:"name"`
	InternalId string `json:"internalId"`
	Pincode    int    `json:"pincode"`
	Email      string `json:"email"`
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	var c customer

	err = json.Unmarshal(body, &c)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = db.AddC(c.Name, c.InternalId, c.Pincode, c.Email)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}


	respJs, err := json.Marshal(resp{
		Err: false,
		Msg: "Customer added successfully!",
	})
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	sendResp(w, http.StatusCreated, respJs)

}
