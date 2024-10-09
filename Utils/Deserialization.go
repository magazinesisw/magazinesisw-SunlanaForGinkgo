package Utils

import (
	"encoding/json"
	"fmt"
)

func JsonDeserialization(data string, obj interface{}) error {

	err := json.Unmarshal([]byte(data), obj)
	if err != nil {
		return fmt.Errorf("反序列化失败: %v", err)
	}

	return nil
}
