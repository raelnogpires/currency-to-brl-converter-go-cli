package main

import (
	"fmt"
	"log"

	"github.com/raelnogpires/currency-to-brl-converter-go-cli/fetch"
	"github.com/raelnogpires/currency-to-brl-converter-go-cli/format_data"
)

func main() {
	fmt.Print("currency -> ")

	var s string

	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		log.Fatal(err)
	}

	exchange := fetch.Fetch(s)
	if err != nil {
		log.Fatal(err)
	}

	ask, time := format_data.FormatData(exchange)

	response := fmt.Sprintf("\n%s\ncurrency for %s in brl now is: R$%.2f", time, s, ask)

	fmt.Println(response)
}
