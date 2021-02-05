package main

import (
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	Login struct {
		User     string
		Passward string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatal("cant open %s", err)
	}
	defer file.Close()
	cfg := &Config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		log.Fatal("error donfig decode %s", err)
	}
	fmt.Printf("%+v\n", cfg)
}
