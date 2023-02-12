package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	sliceD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(sliceD)
	fmt.Println(string(slcB))
}
