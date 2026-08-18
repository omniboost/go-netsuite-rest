package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestCreditMemosGet(t *testing.T) {
	req := client.NewCreditMemosGetRequest()
	req.QueryParams().Limit = 10
	req.QueryParams().Q = `SELECT * FROM creditMemo WHERE createdDate ON_OR_AFTER "01.04.24"`
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
