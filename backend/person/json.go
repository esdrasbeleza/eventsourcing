package person

import "encoding/json"

func outputJSON(something interface{}) []byte {
	json, _ := json.Marshal(something)
	return json
}
