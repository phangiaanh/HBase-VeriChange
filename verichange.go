package main

import (
	"hbase-verichange/generate"
	"hbase-verichange/verify"
)

func main() {
	verify.VerifySalesPerson()
	generate.GenerateSalesPerson()

	verify.VerifyStore()
	generate.GenerateStore()

	verify.VerifyCustomer()
	generate.GenerateCustomer()

	verify.VerifySalesOrderHeader()
	generate.GenerateSalesOrderHeader()
}
