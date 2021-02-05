package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	resp, err := http.Get("https://api.github.com/users/" + login)
	if err != nil {
		log.Fatalf("can not call")
		return nil, err
	}

	defer resp.Body.Close()

	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error : %s", err)
	}

	fmt.Printf("%+v\n", user)
}
