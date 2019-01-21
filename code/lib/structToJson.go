package lib

import (
	"encoding/json"
	"fmt"
)

func StructToJson(obj interface{}) []byte {
	if jsonBytes, err := json.Marshal(obj); err == nil {
		return jsonBytes
	} else {
		fmt.Println(err)
		return nil
	}

}
