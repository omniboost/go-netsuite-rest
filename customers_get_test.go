package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestCustomersGet(t *testing.T) {
	req := client.NewCustomersGetRequest()
	req.QueryParams().Limit = 10
	req.QueryParams().Offset = 10
	// req.QueryParams().Q = "id BETWEEN_NOT [1, 42]"
	// req.QueryParams().Q = "email START_WITH kees@omniboost"
	req.QueryParams().Q = "subsidiary = 46"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
