package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestInvoicesGet(t *testing.T) {
	req := client.NewInvoicesGetRequest()
	req.QueryParams().Limit = 10
	req.QueryParams().Q = `SELECT * FROM invoice WHERE createdDate ON_OR_AFTER "01.04.24"`
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
