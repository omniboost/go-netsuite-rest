package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustom(t *testing.T) {
	req := client.NewCustomRequest()
	client.SetBaseURL("https://5325833-sb2.restlets.api.netsuite.com/app/site/hosting/restlet.nl")
	req.SetMethod("POST")
	req.QueryParams().Script = 1357
	req.QueryParams().Deploy = 1

	type custom struct {
		NSCreditMemoID int   `json:"nsCreditMemoId"`
		NSInvoiceID    []int `json:"nsInvoiceId"`
	}

	c := custom{
		NSCreditMemoID: 1647254,
		NSInvoiceID: []int{
			1644147,
		},
	}
	req.SetRequestBody(c)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
