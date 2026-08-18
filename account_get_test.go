package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestAccountGet(t *testing.T) {
	req := client.NewAccountGetRequest()
	// req.QueryParams().Account = "FLD"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
