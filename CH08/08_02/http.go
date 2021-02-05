package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	resp, err := http.Get("https://api.github.com/users/tebeka")
	if err != nil {
		log.Fatalf("can not call")
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	fmt.Println("--------")

	//Post request
	job := &Job{
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer                    // Post function will cause the request body to be io.writter.
	enc := json.NewEncoder(&buf)            // use bytes, which is an in memory writer or reader
	if err := enc.Encode(job); err != nil { //encode the job inside the in memory buffer
		log.Fatalf("error: cant encdoe %s", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error : cant call")
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
