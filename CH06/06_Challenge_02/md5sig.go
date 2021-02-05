package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parseSignaturesFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	sigs := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", path, lnum)

		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}

type result struct {
	path  string
	match bool
	err   error
}

func fileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	fmt.Printf("%x %s\n", hash.Sum(nil), path)
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func md5Worker(path string, sig string, out chan *result) {
	r := &result{path: path}
	s, err := fileMD5(path)
	if err != nil {
		r.err = err
		out <- r
		return
	}

	r.match = (s == sig)
	out <- r
}
func main() {
	p := "D:\\99_private\\gotest\\Exercise\\Ex_Files_Go_EssT\\Ex_Files_Go_EssT\\Exercise Files\\Ch06\\06_07\\nasa-logs\\"
	sigs, err := parseSignaturesFile(p + "md5sum.txt")
	if err != nil {
		log.Fatal("%s", err)
	}

	out := make(chan *result)
	for path, sig := range sigs {
		go md5Worker(p+path, sig, out)
	}

	ok := true
	for range sigs {
		r := <-out
		switch {
		case r.err != nil:
			fmt.Printf("%s\n", r.path, r.err)
			ok = false
		case !r.match:
			fmt.Printf("mismatch %s\n", r.path)
			ok = false
		}
	}

	if !ok {
		os.Exit(1)
	}
}
