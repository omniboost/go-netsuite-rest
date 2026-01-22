package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLocationGet(t *testing.T) {
	req := client.NewLocationGetRequest()
	req.PathParams().ID = 126
	// req.QueryParams().Fields = netsuite.Fields{"line"}
	// req.QueryParams().ExpandSubResources = true
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
