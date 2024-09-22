package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://ajax.gogocdn.net/ajax/load-list-episode?ep_start=0&ep_end=11&id=15058&default_ep=0&alias=boku-no-hero-academia-7th-season"

func main() {
	res, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	defer res.Body.Close()

	bytesData, err := ioutil.ReadAll(res.Body)

	contain := string(bytesData)

	fmt.Printf("Data: %v\n", contain)
}
