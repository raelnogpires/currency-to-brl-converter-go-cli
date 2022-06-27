package fetch

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Fetch(curr string) []byte {
	search := strings.ToUpper(curr)

	url := fmt.Sprintf("https://economia.awesomeapi.com.br/last/%v-BRL", search)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
