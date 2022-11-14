package generate

import (
	"encoding/json"
	"fmt"
	"hbase-verichange/verify"
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

func GenerateSalesPerson() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesPerson.json")
	var res SalesPersonDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for a := 1; a <= 273; a++ {
		var newItem SalesPerson

		i := int64(rand.Intn(5) + 25)
		var f float64 = float64(i * 10000)
		j := int64(rand.Intn(500) + 200)
		var ff float64 = float64(j * 1000)
		k := int64(rand.Intn(10) + 10)
		var fff float64 = float64(float64(k) / 1000)
		l := int64(rand.Intn(24000000000) + 10000000000)
		var ffff float64 = float64(l / 10000)
		m := int64(rand.Intn(24000000000) + 10000000000)
		var fffff float64 = float64(m / 10000)

		newItem.BusinessEntityID = int64(a)
		newItem.TerritoryID = int64(rand.Intn(10))
		newItem.SalesQuota = fmt.Sprintf("%.4f", f)
		newItem.Bonus = fmt.Sprintf("%.4f", ff)
		newItem.CommissionPct = fmt.Sprintf("%.4f", fff)
		newItem.SalesYTD = fmt.Sprintf("%.4f", ffff)
		newItem.SalesLastYear = fmt.Sprintf("%.4f", fffff)

		res.SalesPersonDB = append(res.SalesPersonDB, newItem)
	}

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

func GenerateStore() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/Store.json")
	var res StoreDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	var j int64 = 3000
	for a := 1; a <= 273; a++ {
		k := rand.Intn(20) + 70
		for b := 1; b <= k; b++ {
			var newItem Store
			newItem.BusinessEntityID = j
			j = j + 1
			newItem.Name = "Random Name"
			newItem.SalesPersonID = int64(a)

			res.StoreDB = append(res.StoreDB, newItem)
		}

	}

	verify.MaximumStore = j - 1
	MaximumStore = j - 1

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

func GenerateCustomer() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/Customer.json")
	var res CustomerDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

}
