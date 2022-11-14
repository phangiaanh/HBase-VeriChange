package verify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

var (
	MaximumStore int64
)

type SalesPersonDB struct {
	SalesPersonDB []SalesPerson `json:"SalesPerson"`
}

type SalesPerson struct {
	BusinessEntityID int64  `json:"BusinessEntityID"`
	TerritoryID      int64  `json:"TerritoryID"`
	SalesQuota       string `json:"SalesQuota"`
	Bonus            string `json:"Bonus"`
	CommissionPct    string `json:"CommissionPct"`
	SalesYTD         string `json:"SalesYTD"`
	SalesLastYear    string `json:"SalesLastYear"`
	rowguid          string `json:"rowguid"`
}

type StoreDB struct {
	StoreDB []Store `json:"Store"`
}

type Store struct {
	BusinessEntityID int64  `json:"BusinessEntityID"`
	Name             string `json:"Name"`
	SalesPersonID    int64  `json:"SalesPersonID"`
	Demographics     string `json:"Demographics"`
	rowguid          string `json:"rowguid"`
}

type CustomerDB struct {
	CustomerDB []Customer `json:"Customer"`
}

type Customer struct {
	CustomerID    int64  `json:"CustomerID"`
	PersonID      int64  `json:"PersonID"`
	StoreID       int64  `json:"StoreID"`
	TerritoryID   int64  `json:"TerritoryID"`
	AccountNumber string `json:"AccountNumber"`
	rowguid       string `json:"rowguid"`
	ModifiedDate  string `json:"ModifiedDate"`
}

func VerifySalesPerson() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesPerson.json")
	var res SalesPersonDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for j, item := range res.SalesPersonDB {
		// fmt.Println(item.TerritoryID)
		if item.TerritoryID == 0 {
			item.TerritoryID = int64(rand.Intn(10))
		}
		// fmt.Println(item.SalesQuota)
		if item.SalesQuota == "" {
			i := int64(rand.Intn(5) + 25)
			var f float64 = float64(i * 10000)
			item.SalesQuota = fmt.Sprintf("%.4f", f)
		}

		res.SalesPersonDB[j] = item
	}

	// fmt.Println(res)

	content, err := json.MarshalIndent(res, "", "\t")
	// _ = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", file, 0777)
	// content, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func VerifyStore() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/Store.json")
	var res StoreDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for j, item := range res.StoreDB {

		if item.Name == "" {
			item.Name = "Random Name"
		}

		if item.SalesPersonID == 0 {
			i := int64(rand.Intn(282) + 1)
			item.SalesPersonID = i
		}

		res.StoreDB[j] = item
	}

	// fmt.Println(res)

	content, err := json.MarshalIndent(res, "", "\t")
	// _ = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", file, 0777)
	// content, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../HBase-Migration/export_json/Store.json", content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func VerifyCustomer() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/Customer.json")
	var res CustomerDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for j, item := range res.CustomerDB {

		if item.StoreID == 0 {
			fmt.Println(item.CustomerID)
			i := rand.Intn(int(MaximumStore)-3000) + 3000
			item.StoreID = int64(i)
		}

		if item.TerritoryID == 0 {
			i := rand.Intn(10)
			item.TerritoryID = int64(i)
		}

		res.CustomerDB[j] = item
	}

	// fmt.Println(res)

	content, err := json.MarshalIndent(res, "", "\t")
	// _ = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", file, 0777)
	// content, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../HBase-Migration/export_json/Customer.json", content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
