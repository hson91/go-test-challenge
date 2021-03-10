package utilities

import (
	"encoding/json"
)

// Switch2Struct : switch 2 struct
func Switch2Struct(in interface{}, out interface{}) error {
	slcW, err := json.Marshal(in)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(slcW, &out); err != nil {
		return err
	}

	return nil
}

// PrintKeysOfObject :
func PrintKeysOfObject(obj interface{}) (string, error) {
	var (
		dict map[string]interface{}
	)

	if err := Switch2Struct(&obj, &dict); err != nil {
		return "", err
	}

	result := "{ \n"

	for k := range dict {
		result += "\t" + k + " \n"
	}

	result += "\n }"

	return result, nil
}
