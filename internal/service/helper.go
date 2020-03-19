package service

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func marshalToJSON(src interface{}) (string, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return "", fmt.Errorf("model couldn't be marshaled into JSON format. %v\n", err)
	}
	return string(data), nil
}

func unmarshalToStruct(data []byte, dst interface{}) error {
	err := json.Unmarshal(data, dst)
	if err != nil {
		return fmt.Errorf("data [%s] couldn't be unmarshaled into the given struct [%s]. %v\n", string(data), reflect.TypeOf(dst), err)
	}
	return nil
}
