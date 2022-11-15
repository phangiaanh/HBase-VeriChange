package main

import (
	"fmt"
	"hbase-verichange/generate"
	"hbase-verichange/verify"
)

func main() {
	verify.VerifySalesPerson()
	generate.GenerateSalesPerson()
	fmt.Println("GenerateSalesPerson")

	verify.VerifyStore()
	generate.GenerateStore()
	fmt.Println("GenerateStore")

	verify.VerifyCustomer()
	generate.GenerateCustomer()
	fmt.Println("GenerateCustomer")

	verify.VerifySalesOrderHeader()
	generate.GenerateSalesOrderHeader()
	fmt.Println("GenerateSalesOrderHeader")

	generate.GenerateSalesOrderHeaderSalesReason()
	fmt.Println("GenerateSalesOrderHeaderSalesReason")

	verify.VerifySalesPersonQuotaHistory()
	generate.GenerateSalesPersonQuotaHistory()
	fmt.Println("GenerateSalesPersonQuotaHistory")
}
