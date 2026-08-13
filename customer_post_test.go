package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerPost(t *testing.T) {
	req := client.NewCustomerPostRequest()
	req.RequestBody().Subsidiary.ID = "46"
	req.RequestBody().IsPerson = true
	req.RequestBody().FirstName = "Kees"
	req.RequestBody().LastName = "Zorge"
	req.RequestBody().Email = "kees@Omniboost.io"
	req.RequestBody().Phone = "1335132342"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
