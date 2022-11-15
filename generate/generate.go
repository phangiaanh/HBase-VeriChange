package generate

import (
	"encoding/json"
	"fmt"
	"hbase-verichange/verify"
	"io/ioutil"
	"math/rand"
	"strconv"
)

var (
	MaximumStore    int64
	MaximumCustomer int64
	MaximumSales    int64
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

type SalesOrderHeaderSalesReasonDB struct {
	SalesOrderHeaderSalesReason []SalesOrderHeaderSalesReason `json:"SalesOrderHeaderSalesReason"`
}

type SalesOrderHeaderSalesReason struct {
	SalesOrderID  int64 `json:"SalesOrderID"`
	SalesReasonID int64 `json:"SalesReasonID"`
}

type SalesPersonQuotaHistoryDB struct {
	SalesPersonQuotaHistoryDB []SalesPersonQuotaHistory `json:"SalesPersonQuotaHistory"`
}

type SalesPersonQuotaHistory struct {
	BusinessEntityID int64  `json:"BusinessEntityID"`
	QuotaDate        string `json:"QuotaDate"`
	SalesQuota       string `json:"SalesQuota"`
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
	fmt.Println(MaximumStore)

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

	var j int64 = 40000
	for a := 0; a <= int(MaximumStore)-3000; a++ {
		k := rand.Intn(15) + 5
		for b := 1; b <= k; b++ {
			var newItem Customer
			newItem.CustomerID = j
			j = j + 1
			newItem.StoreID = int64(a + 3000)
			newItem.PersonID = newItem.StoreID - 1
			newItem.TerritoryID = int64(rand.Intn(10))
			newItem.AccountNumber = "Random Account"

			res.CustomerDB = append(res.CustomerDB, newItem)
		}

	}

	verify.MaximumCustomer = j - 1
	MaximumCustomer = j - 1

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

func GenerateSalesOrderHeader() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesOrderHeader.json")
	var res SalesOrderHeaderDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	var j int64 = 80000
	for a := 1; a <= 273; a++ {
		k := rand.Intn(300) + 1800
		for b := 1; b <= k; b++ {
			var newItem SalesOrderHeader
			newItem.SalesOrderID = j
			j = j + 1
			i := rand.Intn(int(MaximumCustomer)-40000) + 40000
			newItem.CustomerID = int64(i)
			newItem.SalesPersonID = int64(a)

			res.SalesOrderHeaderDB = append(res.SalesOrderHeaderDB, newItem)
		}

	}
	MaximumSales = j - 1

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

func GenerateSalesOrderHeaderSalesReason() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesOrderHeaderSalesReason.json")
	var res SalesOrderHeaderSalesReasonDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for a := 80000; a <= int(MaximumSales); a++ {
		var newItem SalesOrderHeaderSalesReason
		newItem.SalesOrderID = int64(a)
		newItem.SalesReasonID = int64(rand.Intn(10))

		res.SalesOrderHeaderSalesReason = append(res.SalesOrderHeaderSalesReason, newItem)

	}

	content, err := json.MarshalIndent(res, "", "\t")
	// _ = ioutil.WriteFile("../HBase-Migration/export_json/SalesPerson.json", file, 0777)
	// content, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("../HBase-Migration/export_json/SalesOrderHeaderSalesReason.json", content, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateSalesPersonQuotaHistory() {
	data, _ := ioutil.ReadFile("../HBase-Migration/export_json/SalesPersonQuotaHistory.json")
	var res SalesPersonQuotaHistoryDB
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
	}

	for a := 1; a <= 273; a++ {
		k := rand.Intn(18) + 1
		for b := 1; b <= k; b++ {
			var newItem SalesPersonQuotaHistory
			newItem.BusinessEntityID = int64(a)
			if b > 12 {
				newItem.QuotaDate = fmt.Sprintf("2022/%s", strconv.FormatInt(int64(b)-12, 10))
			} else {
				newItem.QuotaDate = fmt.Sprintf("2021/%s", strconv.FormatInt(int64(b), 10))
			}
			k := float64((rand.Intn(900) + 100) * 1000)
			newItem.SalesQuota = fmt.Sprintf("%.4f", k)

			res.SalesPersonQuotaHistoryDB = append(res.SalesPersonQuotaHistoryDB, newItem)
		}

	}

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
