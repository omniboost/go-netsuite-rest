package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestSalesTaxItemsGet(t *testing.T) {
	req := client.NewSalesTaxItemsGetRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
