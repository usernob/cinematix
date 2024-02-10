package logjson

import (
	"encoding/json"
	"fmt"
)

func ToJSON(data interface{}) {
  b, _ := json.MarshalIndent(data, "", "  ")
  fmt.Println(string(b))
}
