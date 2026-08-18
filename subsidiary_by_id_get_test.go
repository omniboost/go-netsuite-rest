package netsuite_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestSubsidiaryByIDGet(t *testing.T) {
	req := client.NewSubsidiaryByIDGetRequest()
	req.PathParams().SubsidiaryID = "710" // Set the subsidiary ID here
	req.QueryParams().ExpandSubResources = true
	client.SetDisallowUnknownFields(false)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
