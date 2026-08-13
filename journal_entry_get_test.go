package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestJournalEntryGet(t *testing.T) {
	req := client.NewJournalEntryGetRequest()
	req.PathParams().ID = 1299002
	// req.QueryParams().Fields = netsuite.Fields{"line"}
	req.QueryParams().ExpandSubResources = true
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
