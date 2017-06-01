package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.andreas-grohmann.com")
	if err != nil {
		log.Fatalf("fetch : %v", err)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}