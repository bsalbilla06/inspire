package main

import (
  "fmt"
  "encoding/json"
)

type Quote struct {
	Words string	`json:"q"`
	Author string	`json:"a"`
}

func main() {
	var quote1 Quote

	Data := []byte(`{
		"q": "A wise man once said nothing at all",
		"a": "Drake"
	}`)

	err := json.Unmarshal(Data, &quote1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v said \"%v\"\n", quote1.Author, quote1.Words)
}
