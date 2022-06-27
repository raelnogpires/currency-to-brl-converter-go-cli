package format_data

import (
	"log"
	"strconv"

	"github.com/spyzhov/ajson"
)

func FormatData(d []byte) (interface{}, interface{}) {
	root, _ := ajson.Unmarshal(d)

	ask, _ := root.JSONPath("$..ask")
	time, _ := root.JSONPath("$..create_date")

	if len(ask) == 0 {
		log.Fatalf("\nthe currency searched does not exists :(")
	}

	askInterface, _ := ask[0].Value()
	timeString, _ := time[0].Value()

	stringAsk := askInterface.(string)
	floatAsk, _ := strconv.ParseFloat(stringAsk, 64)

	return floatAsk, timeString
}
