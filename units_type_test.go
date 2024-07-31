package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUnitsTypeGet(t *testing.T) {
	req := client.NewUnitsTypeGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
