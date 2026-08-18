package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestStatisticalJournalEntriesGet(t *testing.T) {
	req := client.NewStatisticalJournalEntriesGetRequest()
	req.QueryParams().Limit = 100
	req.QueryParams().Offset = 100
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
