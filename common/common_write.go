package common

import (
	"encoding/json"
	"log"
	"os"
)

func WriteJSON(filename string, data interface{}) {
	byt, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Panicln(err)
	}

	err = os.WriteFile(filename, byt, 0755)
	if err != nil {
		log.Panicln(err)
	}
}
