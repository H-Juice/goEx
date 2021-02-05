package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { //ResponseWriter can write to Response
	fmt.Fprintf(w, "Hello Gophers!")
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) { //handler receives http.ReponseWriter and pointer to a request
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body) //Request body r pointer allocated to dec json decoder
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil { //
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "division by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("unknown operation : %s")
	}

	//Enconde and return result
	w.Header().Set("Content-Typ", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("cant encode %v %s", resp, err)
	}
}
func main() {
	http.HandleFunc("/hello", helloHandler) //Mount Endpoint to helloHandler
	http.HandleFunc("/math", mathHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil { //Potentially endless loop
		log.Fatal(err)
	}

}
