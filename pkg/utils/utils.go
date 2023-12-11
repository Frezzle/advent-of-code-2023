package utils

import "encoding/json"

// Source: https://stackoverflow.com/a/51270134
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
