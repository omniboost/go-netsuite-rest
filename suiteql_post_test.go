package netsuite_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSuiteqlPost(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	//req.RequestBody().Q = "Select id, email from customer where email = 'leon@omniboost.io'"
	//req.RequestBody().Q = "SELECT * FROM customer where customer.firstName = 'Pieter' and lastName = 'Baecke' and customer.subsidiary.id = 46"
	//req.RequestBody().Q = "SELECT * FROM department where name like '%196%'"
	//req.RequestBody().Q = "SELECT * FROM department"
	//req.RequestBody().Q = "SELECT * FROM account"
	//req.RequestBody().Q = "SELECT * FROM classification"
	req.RequestBody().Q = "SELECT TOP 10 transaction.* FROM transaction INNER JOIN customlist_nch_source ON transaction.custbody_nch_source = customlist_nch_source.internalid and customlist_nch_source.name not in ('Mews', 'EventTemple') WHERE transaction.createdDate >= to_date('2025-10-01 00:00:00', 'yyyy-mm-dd hh24:mi:ss') AND transaction.recordtype = 'invoice'"
	req.RequestBody().Q = "SELECT TOP 10 transaction.* FROM transaction WHERE transaction.createdDate >= to_date('2025-10-01 00:00:00', 'yyyy-mm-dd hh24:mi:ss') AND transaction.recordtype = 'invoice' AND BUILTIN.DF(custbody_nch_source) not in ('Mews', 'EventTemple')"
	//req.RequestBody().Q = "SELECT TOP 10 * FROM transaction WHERE createdDate >= to_date('2025-10-01 00:00:00', 'yyyy-mm-dd hh24:mi:ss') AND recordtype = 'creditmemo' AND custbody_nch_source NOT IN ('Mews', 'EventTemple')"

	// creditmemo
	// invoice
	// req.RequestBody().Q = "SELECT * FROM classification where name like '%196%'"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
	//creditmemos, err := resp.ToCreditMemoTransactions(client)
	//if err != nil {
	//	t.Error(err)
	//}
	//log.Fatalf("%+v", creditmemos)
	//
	//os.Exit(12)

	invoices, err := resp.ToInvoiceTransactions(client)
	if err != nil {
		t.Error(err)
	}
	log.Fatalf("%+v", invoices)
	os.Exit(12)

	customers, err := resp.ToDepartments(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers)
}

// func TestSuiteqlPostCustomers(t *testing.T) {
// 	req := client.NewSuiteqlPostRequest()
// 	// req.QueryParams().Limit = 12
// 	req.RequestBody().Q = "SELECT * FROM customer where companyName = 'Agate Utvikling AS_170122'"
// 	req.RequestBody().Q = "SELECT * FROM customer where lastmodifieddate > '20.03.2020' and lastmodifieddate < '20.03.2022'"
// 	req.RequestBody().Q = "SELECT * FROM customer where parent = 13299"
// 	req.RequestBody().Q = "SELECT * FROM customer where id IN(2607)"
//
// 	// req.RequestBody().Q = "SELECT * FROM classification where name like '%196%'"
// 	resp, err := req.Do()
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	// b, _ := json.MarshalIndent(resp, "", "  ")
// 	// fmt.Println(string(b))
//
// 	customers, err := resp.ToCustomers(client)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	log.Fatalf("%+v", customers[1])
// }

func TestSuiteqlPostAddresses(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.QueryParams().Limit = 12
	req.RequestBody().Q = "SELECT * FROM EntityAddress where nkey IN (458404, 458304)"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	customers, err := resp.ToAddresses(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers[1])
}

func TestSuiteqlPostLocations(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.QueryParams().Limit = 100
	req.RequestBody().Q = "SELECT * FROM location"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	customers, err := resp.ToLocations(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers[1])
}

func TestSuiteqlPostCustomers(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.QueryParams().Limit = 100
	req.RequestBody().Q = "SELECT * FROM customer"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	customers, err := resp.ToCustomers(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", customers[1])
}

func TestSuiteqlPostSalesTaxItems(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.QueryParams().Limit = 100
	req.RequestBody().Q = "SELECT * FROM salestaxitem"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	salesTaxItems, err := resp.ToSalesTaxItems(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", salesTaxItems)
}

func TestSuiteqlPostNexus(t *testing.T) {
	req := client.NewSuiteqlPostRequest()
	req.QueryParams().Limit = 100
	req.RequestBody().Q = "SELECT * FROM nexus where id IN('Netherlands')"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(b))

	nexuses, err := resp.ToNexuses(client)
	if err != nil {
		t.Error(err)
	}

	log.Fatalf("%+v", nexuses)
}
