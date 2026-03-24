package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCertificatesGet(t *testing.T) {
	req := client.NewCertificatesGetRequest()
	req.PathParams().ClientID = "3c7aa83bcd9cc4a1fec810027f3a5d5c050a318c0e8b7b7bc99df1573c80db03"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
