package netsuite

import (
	"encoding/json"
	"strings"
)

func UnmarshalCustomFields(data []byte, v any) ([]byte, error) {
	withCustomFields := map[string]json.RawMessage{}
	withoutCustomFields := map[string]json.RawMessage{}

	err := json.Unmarshal(data, &withoutCustomFields)
	if err != nil {
		return nil, err
	}

	// only keep items that start with custcol_
	for k := range withoutCustomFields {
		if strings.HasPrefix(k, "custcol_") || strings.HasPrefix(k, "custbody_") {
			delete(withoutCustomFields, k)
		} else {
			withCustomFields[k] = withoutCustomFields[k]
		}
	}

	if len(withCustomFields) == 0 {
		return data, nil
	}

	colBytes, err := json.Marshal(withCustomFields)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(colBytes, &v)
	if err != nil {
		return nil, err
	}

	// return with cols removed
	return json.Marshal(withoutCustomFields)
}
