package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestTaxTypesGet(t *testing.T) {
	req := client.NewTaxTypesGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

