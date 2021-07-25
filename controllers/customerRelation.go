package controllers

import (
	"encoding/json"
	"graphApp/db"
	"graphApp/global"
	"io/ioutil"
	"net/http"
)

type cR struct {
	CId1         string `json:"cid1"`
	CId2         string `json:"cid2"`
	RelationType string `json:"relation_type"`
	Date         string `json:"date"`
}

func AddCustomerRelation(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	var cr cR

	err = json.Unmarshal(body, &cr)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	err = db.AddRelation(cr.CId1, cr.CId2, cr.RelationType, cr.Date)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	respJs, err := json.Marshal(resp{
		Err: false,
		Msg: "Added relation!(customer)",
	})
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	sendResp(w, http.StatusCreated, respJs)
}
