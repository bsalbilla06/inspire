package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io"
)

type Quote struct {
	Words string	`json:"q"`
	Author string	`json:"a"`
}

func main() {
	quote := make([]Quote, 1)

	url := "https://zenquotes.io/api/random"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	
	err = json.Unmarshal(body, &quote)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\"%v\" - %v\n", quote[0].Words, quote[0].Author)
}
