package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID   int      `json:"id"`
	Nome string   `json:"nome"`
	Tags []string `json:"tags"`
}

func main() {
	//struct to json
	p1 := produto{1, "Note", []string{"a", "b"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json to string
	var p2 produto
	jsonString := `{"id":1, "nome": "alisson", "tags":["a", "b"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
