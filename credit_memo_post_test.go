package netsuite_test

import (
	"testing"
)

func TestCreditMemoPost(t *testing.T) {
	// req := client.NewCreditMemoPostRequest()
	// req.RequestBody().Subsidiary.ID = "49"
	// // req.RequestBody().Entity.ID = 145
	// // req.RequestBody().Location.ID = 5
	// req.RequestBody().Item = netsuite.CreditMemo{
	// 	Item: netsuite.CreditMemoItemCollection{
	// 		{
	// 			Account: netsuite.Account{
	// 				ID: "635",
	// 			},
	// 			Amount: 80000,
	// 			Item: netsuite.CreditMemoItemItemItem{
	// 				ID: 131,
	// 			},
	// 			ItemSubType: "Resale",
	// 			ItemType:    "NonInvtPart",
	// 		},
	// 	},
	// }
	// resp, err := req.Do()
	// if err != nil {
	// 	t.Error(err)
	// }
	//
	// b, _ := json.MarshalIndent(resp, "", "  ")
	// log.Println(string(b))
}

