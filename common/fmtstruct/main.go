package fmtstruct

import (
	"encoding/json"
	"fmt"
	"bytes"
)

func String(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}
	return out.String()
}