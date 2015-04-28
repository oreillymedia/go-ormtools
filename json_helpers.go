package ormtools

import "encoding/json"

func JsonToMap(body string) (map[string]interface{}, error) {
	var data map[string]interface{}
	jsonErr := json.Unmarshal([]byte(body), &data)
	return data, jsonErr
}
