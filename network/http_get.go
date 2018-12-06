package main

import (
	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://studygolang.com/rank/view?limit=10&objtype=1&rank_type=today")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	//body, _ := ioutil.ReadAll(res.Body)
	var Response struct {
		Data struct {
			List []interface{} `json:"list"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&Response); err != nil {
		log.Fatal(err)
	}
	fmt.Println(Response.Data.List[0])
}
