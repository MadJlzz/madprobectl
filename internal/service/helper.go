package service

import (
	"encoding/json"
	"fmt"
)

func marshalToJSON(src interface{}) (string, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return "", fmt.Errorf("model couldn't be marshaled into JSON format. %v\n", err)
	}
	return string(data), nil
}
