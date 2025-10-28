package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSalesTaxItemGet(t *testing.T) {
	req := client.NewSalesTaxItemGetRequest()
	req.PathParams().ID = 57
	// req.QueryParams().Fields = netsuite.Fields{"line"}
	req.QueryParams().ExpandSubResources = true
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
