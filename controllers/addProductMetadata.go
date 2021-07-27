package controllers

import (
	"encoding/json"
	"graphApp/db"
	"graphApp/global"
	"io/ioutil"
	"net/http"
)

type mD struct {
	Name string `json:"name"`
}

func addProductMetadata(key string, action func(string) error) func(w http.ResponseWriter, r *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}

		var md mD

		err = json.Unmarshal(body, &md)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}

		err = action(md.Name)
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}

		respJs, err := json.Marshal(resp{
			Err: false,
			Msg: key + " added successfully!",
		})
		if err != nil {
			global.NewWebError(w, err, http.StatusInternalServerError)
			return
		}

		sendResp(w, http.StatusCreated, respJs)

	}
	return f
}

var AddBrand = addProductMetadata("Brand", db.AddBrand)
var AddCategory = addProductMetadata("Category", db.AddCategory)
