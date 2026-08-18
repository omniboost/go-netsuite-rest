package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestNexusesGet(t *testing.T) {
	req := client.NewNexusesGetRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
