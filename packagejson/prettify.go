package packagejson

import (
	"encoding/json"
	"log"
)

func prettifyJSON(data []byte) []byte {
	var rawData interface{}
	err := json.Unmarshal(data, &rawData)
	if err != nil {
		log.Fatal(err)
	}

	pretty, err := json.MarshalIndent(rawData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return append(pretty, "\n"...)
}
