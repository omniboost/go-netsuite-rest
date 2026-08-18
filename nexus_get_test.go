package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestNexusGet(t *testing.T) {
	req := client.NewNexusGetRequest()
	req.PathParams().ID = "1"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
