package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestJournalEntryLineGet(t *testing.T) {
	req := client.NewJournalEntryLineGetRequest()
	req.PathParams().JournalEntryID = 3856

	for _, i := range []int{0, 1} {
		req.PathParams().JournalEntryLineNo = i
		// req.QueryParams().Fields = netsuite.Fields{"line"}
		resp, err := req.Do(context.Background())
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(resp, "", "  ")
		log.Println(string(b))
	}
}
