package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	
	res, _ := http.Get("http://google.com.br")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	
	fmt.Printf("%s", body)
}