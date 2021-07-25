package data

import (
	"encoding/csv"
	"fmt"
	"graphApp/db"
	"graphApp/global"
	"net/http"
	"os"
	"strconv"
)

type customerData struct {
	name       string
	internalId string
	email      string
	pincode    string
}

type productData struct {
	name        string
	sku         string
	internalId  string
	price       float64
	description string
	brand       string
	category    string
}

func LoadTestData(w http.ResponseWriter, r *http.Request) {
	err := db.DeleteAll()
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}
	err = addTestCustomers()
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}
}

func addTestCustomers() error {

	csvFile, err := os.Open("./customer.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}

	s := map[string]bool{}
	var data []customerData

	for _, line := range csvLines {
		cd := customerData{
			name:       line[0],
			internalId: line[1],
			email:      line[2],
			pincode:    line[3],
		}
		data = append(data, cd)
		_, ok := s[line[3]]
		if !ok {
			s[line[3]] = true
		}

	}

	err = db.AddPincodesTest(s)
	if err != nil {
		return err
	}

	q := "MATCH "

	for _, val := range data {
		valQ := fmt.Sprintf("(%s:Area{pincode:%s}) ,", val.internalId, val.pincode)
		q = q + valQ
	}
	q = q[:len(q)-1]
	q = q + "CREATE "

	for _, val := range data {
		valQ := fmt.Sprintf("(:Customer{name:'%s',email:'%s',internal_id:'%s'})-[:In]->(%s) ,", val.name, val.email, val.internalId, val.internalId)
		q = q + valQ
	}
	q = q[:len(q)-1]

	err = db.AddCustomersTest(q)
	if err != nil {
		return err
	}

	return nil
}

func addTestProducts() error {

	csvFile, err := os.Open("./product.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}

	var data []productData

	b := map[string]bool{}
	c := map[string]bool{}
	for _, line := range csvLines {

		prcInt, err := strconv.ParseFloat(line[3], 32)
		if err != nil {
			return err
		}

		cd := productData{
			name:        line[0],
			sku:         line[1],
			internalId:  line[2],
			price:       prcInt,
			description: line[4],
			brand:       line[5],
			category:    line[6],
		}
		data = append(data, cd)
		_, ok := b[line[5]]
		if !ok {
			b[line[5]] = true
		}
		_, ok = c[line[6]]
		if !ok {
			b[line[6]] = true
		}

	}

	return nil
}

func addCustomerRelation() error {

	return nil
}

func addCustomerProductRelation() error {

	return nil
}
