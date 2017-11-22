package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"regexp"
	"time"
	
)

func elizaResponse(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("value")

	if matched, _ := regexp.MatchString(`(?i).*\b[Ff]ather\b.*`, input); matched{
		responses := []string {
			"\nI’m not sure what you’re trying to say. Could you explain it to me?",
			"\nHow does that make you feel?",
			"\nWhy do you say that?",
			"\nAre you sure",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
}

func main() {
	rand.Seed(time.Now().Unix()) 

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/elizaResponse", elizaResponse)
	http.ListenAndServe(":8080", nil)
}
