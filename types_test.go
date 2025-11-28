package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"

	netsuite "github.com/omniboost/go-netsuite-rest"
)

func TestJournalEntryCustomFieldsUnmarshalling(t *testing.T) {
	jsonData := `{
		"custbody_int": 42,
		"custbody_float": 3.14,
		"custbody_string": "TEST",
		"custbody_bool": true,
		"custcol_int": 42,
		"custcol_float": 3.14,
		"custcol_string": "TEST",
		"custcol_bool": true,
		"int": 42,
		"amountPaid": 3.14,
		"billAddress": "TEST",
		"toBeEmailed": true
	}`

	var result netsuite.Invoice
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(result.CustomFields) != 8 {
		t.Fatalf("Expected 8 custom fields, got %d", len(result.CustomFields))
	}

	// check if all custom fields have the correct values
	for key, expectedValue := range map[string]any{
		"custbody_int":    float64(42),
		"custbody_float":  3.14,
		"custbody_string": "TEST",
		"custbody_bool":   true,
		"custcol_int":     float64(42),
		"custcol_float":   3.14,
		"custcol_string":  "TEST",
		"custcol_bool":    true,
	} {
		if value, exists := result.CustomFields[key]; !exists {
			t.Errorf("Custom field %s not found", key)
		} else if value != expectedValue {
			t.Errorf("Custom field %s: expected %v, got %v", key, expectedValue, value)
		}
	}

	// check if all the standard fields have the correct values
	// int is not a field of Invoice struct so we skip that test
	// if result.Int != 42 {
	// 	t.Errorf("Expected Int to be 42, got %d", result.Int)
	// }
	if result.AmountPaid != 3.14 {
		t.Errorf("Expected AmountPaid to be 3.14, got %f", result.AmountPaid)
	}
	if result.BillAddress != "TEST" {
		t.Errorf("Expected BillAddress to be TEST, got %s", result.BillAddress)
	}

	b, _ := json.MarshalIndent(result, "", "  ")
	log.Println(string(b))

}

func TestJournalEntryCustomFieldsMarshalling(t *testing.T) {
	result := netsuite.Invoice{
		CustomFields: map[string]any{
			"custbody_int":    42,
			"custbody_float":  3.14,
			"custbody_string": "TEST",
			"custbody_bool":   true,
			"custcol_int":     42,
			"custcol_float":   3.14,
			"custcol_string":  "TEST",
			"custcol_bool":    true,
		},
		AmountPaid:  3.14,
		BillAddress: "TEST",
	}

	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// log the marshalled JSON
	log.Println(string(b))

	var unmarshalledResult netsuite.Invoice
	err = json.Unmarshal(b, &unmarshalledResult)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(unmarshalledResult.CustomFields) != 8 {
		t.Fatalf("Expected 8 custom fields, got %d", len(unmarshalledResult.CustomFields))
	}

	// check if all custom fields have the correct values
	for key, expectedValue := range result.CustomFields {
		value, exists := unmarshalledResult.CustomFields[key]
		if !exists {
			t.Errorf("Custom field %s not found", key)
		}

		_, expectedIsInt := expectedValue.(int)
		_, valueIsFloat := value.(float64)
		if expectedIsInt && valueIsFloat {
			value = int(value.(float64))
		}

		if value != expectedValue {
			// if expectedValue is float64 and value is int, convert value to
			// float64 for comparison
			// this is needed because json.Unmarshal converts all numbers to
			// float64
			t.Errorf("Custom field %s: expected %v, got %v", key, expectedValue, value)
		}
	}

	if unmarshalledResult.AmountPaid != result.AmountPaid {
		t.Errorf("Expected AmountPaid to be %f, got %f", result.AmountPaid, unmarshalledResult.AmountPaid)
	}
	if unmarshalledResult.BillAddress != result.BillAddress {
		t.Errorf("Expected BillAddress to be %s, got %s", result.BillAddress, unmarshalledResult.BillAddress)
	}
}
