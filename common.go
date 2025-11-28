package netsuite

import (
	"encoding/json"
	"maps"
	"strings"
)

var prefixes = []string{
	"custbody_",
	"custcol_",
	"custentity_",
	"custrecord_",
}

func UnmarshalCustomFields(data []byte, v any) error {
	withCustomFields := map[string]json.RawMessage{}
	withoutCustomFields := map[string]json.RawMessage{}

	err := json.Unmarshal(data, &withoutCustomFields)
	if err != nil {
		return err
	}

	// only keep items that start with custcol_
	for k := range withoutCustomFields {
		// check if key starts with any of the prefixes
		for _, prefix := range prefixes {
			if strings.HasPrefix(k, prefix) {
				withCustomFields[k] = withoutCustomFields[k]
				delete(withoutCustomFields, k)
				break
			}
		}
	}

	if len(withCustomFields) == 0 {
		return nil
	}

	b, err := json.Marshal(withCustomFields)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	return nil
}

func MarshalCustomFields(v any, customFields map[string]any) ([]byte, error) {
	// marshal v to json bytes
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// unmarshal to map
	dataMap := map[string]any{}
	err = json.Unmarshal(b, &dataMap)
	if err != nil {
		return nil, err
	}

	// copy all custom fields into dataMap
	maps.Copy(dataMap, customFields)

	// marshal back to json
	b, err = json.Marshal(dataMap)
	if err != nil {
		return nil, err
	}

	return b, nil
}
