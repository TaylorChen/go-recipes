package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email,omitempty"`
}

func main() {
	u := User{ID: 1, Name: "gopher"}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))

	var v User
	_ = json.Unmarshal([]byte(`{"id":2,"name":"go"}`), &v)
	fmt.Printf("unmarshal: %+v\n", v)
}
