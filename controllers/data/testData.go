package data

import (
	"encoding/csv"
	"fmt"
	"graphApp/db"
	"graphApp/global"
	"log"
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

type customerRelation struct {
	cId1  string
	cId2  string
	rType string
	date  string
}
type customerProductRelation struct {
	cId   string
	pId   string
	rType string
	date  string
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
	err = addTestProducts()
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}
	err = addCustomerRelationTest()
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}
	err = addCustomerProductRelation()
	if err != nil {
		global.NewWebError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added Test Data!"))

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

		if err != nil {
			return err
		}
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
		valQ := fmt.Sprintf("(%s:Area{pincode:'%s'}) ,", val.internalId, val.pincode)
		q = q + valQ
	}
	q = q[:len(q)-1]
	q = q + "CREATE "

	for _, val := range data {
		valQ := fmt.Sprintf("(:Customer{name:'%s',email:'%s',internal_id:'%s'})-[:In]->(%s) ,", val.name, val.email, val.internalId, val.internalId)
		q = q + valQ
	}
	q = q[:len(q)-1]
	log.Println(q)

	err = db.RunQuery(q)

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
			c[line[6]] = true
		}

	}

	err = db.AddBrandsTest(b)
	if err != nil {
		return err
	}
	err = db.AddCategoriesTest(c)
	if err != nil {
		return err
	}

	q := ""
	q = q + "MATCH "

	for _, val := range data {
		valQ := fmt.Sprintf("(%s:Brand{name:'%s'}) ,", fmt.Sprintf("%s%d", val.internalId, 1), val.brand)
		q = q + valQ
	}
	for _, val := range data {
		valQ := fmt.Sprintf("(%s:Category{name:'%s'}) ,", fmt.Sprintf("%s%d", val.internalId, 2), val.category)
		q = q + valQ
	}

	q = q[:len(q)-1]
	q = q + "CREATE "

	for _, val := range data {
		valQ := fmt.Sprintf("(%s)<-[:By]-(:Product{name:'%s',email:'%s',internal_id:'%s',price:%v,description:'%s'})-[:Belongs]->(%s) ,", fmt.Sprintf("%s%d", val.internalId, 1), val.name, val.sku, val.internalId, val.price, val.description, fmt.Sprintf("%s%d", val.internalId, 2))
		q = q + valQ
	}
	q = q[:len(q)-1]

	err = db.RunQuery(q)
	if err != nil {
		return err
	}

	return nil
}

func addCustomerRelationTest() error {

	csvFile, err := os.Open("./cRelation.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return err
	}

	var data []customerRelation
	q := ""

	for _, line := range csvLines {
		cr := customerRelation{
			cId1:  line[0],
			cId2:  line[1],
			rType: line[2],
			date:  line[3],
		}
		data = append(data, cr)
	}
	q = q + "MATCH "
	for _, val := range data {
		q = q + fmt.Sprintf("(%s:Customer{internal_id:'%s'}) ,(%s:Customer{internal_id:'%s'}) ,", val.cId1, val.cId1, val.cId2, val.cId2)
	}
	q = q[:len(q)-1]
	q = q + "CREATE "

	for _, val := range data {
		q = q + fmt.Sprintf("(%s)-[:Relation{date:'%s',type:'%s'}]->(%s) ,", val.cId1, val.date, val.rType, val.cId2)
	}
	q = q[:len(q)-1]

	err = db.RunQuery(q)
	if err != nil {
		return err
	}

	return nil
}

func addCustomerProductRelation() error {
	csvFile, err := os.Open("./cpRelation.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return err
	}

	var data []customerProductRelation
	q := ""

	for _, line := range csvLines {
		cr := customerProductRelation{
			cId:   line[0],
			pId:   line[1],
			rType: line[2],
			date:  line[3],
		}
		data = append(data, cr)
	}
	q = q + "MATCH "
	for _, val := range data {
		q = q + fmt.Sprintf("(%s:Customer{internal_id:'%s'}) ,(%s:Product{internal_id:'%s'}) ,", val.cId, val.cId, val.pId, val.pId)
	}
	q = q[:len(q)-1]
	q = q + "CREATE "

	for _, val := range data {
		q = q + fmt.Sprintf("(%s)-[:%s{date:'%s'}]->(%s) ,", val.cId, val.rType, val.date, val.pId)
	}
	q = q[:len(q)-1]

	err = db.RunQuery(q)
	if err != nil {
		return err
	}

	return nil
}
