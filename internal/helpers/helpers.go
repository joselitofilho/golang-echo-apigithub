package helpers

import "encoding/json"

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	objMap := make(map[string]interface{})

	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &objMap); err != nil {
		return nil, err
	}

	return objMap, nil
}
