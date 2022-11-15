package verify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

var (
	MaximumStore    int64
	MaximumCustomer int64
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

type SalesOrderHeaderDB struct {
	SalesOrderHeaderDB []SalesOrderHeader `json:"SalesOrderHeader"`
}

type SalesOrderHeader struct {
	SalesOrderID int64 `json:"SalesOrderID"`

	OrderDate string `json:"OrderDate"`
	DueDate   string `json:"DueDate"`
	ShipDate  string `json:"ShipDate"`

	Status          int64 `json:"Status"`
	OnlineOrderFlag int64 `json:"OnlineOrderFlag"`

	RevisionNumber      int64  `json:"RevisionNumber"`
	SalesOrderNumber    string `json:"SalesOrderNumber"`
	PurchaseOrderNumber string `json:"PurchaseOrderNumber"`
	AccountNumber       string `json:"AccountNumber"`

	CustomerID             int64  `json:"CustomerID"`
	SalesPersonID          int64  `json:"SalesPersonID"`
	TerritoryID            int64  `json:"TerritoryID"`
	BillToAddressID        int64  `json:"BillToAddressID"`
	ShipToAddressID        int64  `json:"ShipToAddressID"`
	ShipMethodID           int64  `json:"ShipMethodID"`
	CreditCardID           int64  `json:"CreditCardID"`
	CreditCardApprovalCode string `json:"CreditCardApprovalCode"`
	CurrencyRateID         int64  `json:"CurrencyRateID"`

	SubTotal string `json:"SubTotal"`
	TaxAmt   string `json:"TaxAmt"`
	Freight  string `json:"Freight"`
	TotalDue string `json:"TotalDue"`
	Comment  string `json:"Comment"`
	rowguid  string `json:"rowguid"`
}

type SalesPersonQuotaHistoryDB struct {
	SalesPersonQuotaHistoryDB []SalesPersonQuotaHistory `json:"SalesPersonQuotaHistory"`
}

type SalesPersonQuotaHistory struct {
	BusinessEntityID int64  `json:"BusinessEntityID"`
	QuotaDate        string `json:"QuotaDate"`
	SalesQuota       string `json:"SalesQuota"`
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
			// fmt.Println(item.CustomerID)
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

func VerifySalesOrderHeader() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesOrderHeader.json")
	var res SalesOrderHeaderDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for j, item := range res.SalesOrderHeaderDB {

		if item.CustomerID == 0 {
			// fmt.Println(item.CustomerID)
			i := rand.Intn(int(MaximumCustomer)-40000) + 40000
			item.CustomerID = int64(i)
		}

		if item.TerritoryID == 0 {
			i := rand.Intn(10)
			item.TerritoryID = int64(i)
		}

		if item.SalesPersonID == 0 {
			item.SalesPersonID = int64(rand.Intn(273) + 1)
		}

		res.SalesOrderHeaderDB[j] = item
	}

	// fmt.Println(res)

	content, err := json.MarshalIndent(res, "", "\t")
	// _ = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", file, 0777)
	// content, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../HBase-Migration/export_json/SalesOrderHeader.json", content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func VerifySalesPersonQuotaHistory() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesPersonQuotaHistory.json")
	var res SalesPersonQuotaHistoryDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	var saleID int64 = 1234567
	var month int64 = 1
	for j, item := range res.SalesPersonQuotaHistoryDB {

		if item.BusinessEntityID != saleID {
			fmt.Println(item.BusinessEntityID)
			saleID = item.BusinessEntityID
			month = 1
			item.QuotaDate = fmt.Sprintf("2021/%s", strconv.FormatInt(month, 10))
		} else {
			month += 1
			if month > 12 {
				item.QuotaDate = fmt.Sprintf("2022/%s", strconv.FormatInt(month-12, 10))
			} else {
				item.QuotaDate = fmt.Sprintf("2021/%s", strconv.FormatInt(month, 10))
			}
		}

		res.SalesPersonQuotaHistoryDB[j] = item
	}

	// fmt.Println(res)

	content, err := json.MarshalIndent(res, "", "\t")
	// _ = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", file, 0777)
	// content, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../HBase-Migration/export_json/SalesPersonQuotaHistory.json", content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
