package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "Scrooge McDuck",
	"type": "desposit",
	"amount": 100000.3
}
`

// Request is a bank transactions
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data) //simulate a file/socket

	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("cant decode %s", err)
	}

	fmt.Printf("got: %+v\n", req)

	//Create response
	prevBalance := 850000.0         // Loeaded from database
	resp := map[string]interface{}{ //key is string, value can be of different type.Empty interface means any type.
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	//Encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error : cant encode %s", err)
	}

}
