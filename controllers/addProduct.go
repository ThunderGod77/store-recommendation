package controllers

import (
	"encoding/json"
	"graphApp/db"
	"graphApp/global"
	"io/ioutil"
	"net/http"
)

type product struct {
	Name        string  `json:"name"`
	Sku         string  `json:"sku"`
	Id          string  `json:"id"`
	Brand       string  `json:"brand"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

type resp struct {
	Err bool   `json:"err"`
	Msg string `json:"msg"`
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	var p product

	err = json.Unmarshal(body, &p)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = db.AddP(p.Name, p.Sku, p.Id, p.Price, p.Description, p.Brand, p.Category)
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	respJs, err := json.Marshal(resp{
		Err: false,
		Msg: "Product added successfully!",
	})
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	sendResp(w, http.StatusCreated, respJs)

}
