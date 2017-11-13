package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"regexp"
	"time"
	
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w ,elizaResponse(r.URL.Query().Get("value")))
	
}

func elizaResponse(input string )string {

	if matched, _ := regexp.MatchString(`(?i).*\b[Ff]ather\b.*`, input); matched{
		return "\nWhy dont you tell me more about your father?"
	}

	responses := []string {
		"\nI’m not sure what you’re trying to say. Could you explain it to me?",
		"\nHow does that make you feel?",
		"\nWhy do you say that?",
		"\nAre you sure",
	}

	randIndex := rand.Intn(len(responses))
	return responses[randIndex]
}

func main() {
	rand.Seed(time.Now().Unix()) 

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
