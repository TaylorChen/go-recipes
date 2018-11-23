package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Radar struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	res, err := http.Get("https://my.oschina.net/didispace/radar/getUserPortraitRadarMap")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	isV := json.Valid(body)
	var r Radar
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Println(isV, err)
	}
	fmt.Println(r.Code, r.Message, r.Time)
}
