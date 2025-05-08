package utils

import (
	"bytes"
	"encoding/json"
)

func JSONPrettyPrint(in interface{}) string {
	data, _ := json.Marshal(in)
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "\t")
	if err != nil {
		return string(data)
	}
	return out.String()
}
