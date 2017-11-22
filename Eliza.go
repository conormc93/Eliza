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

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ello|[Hh]ey|[Hh]i.*`, input); matched{
		responses := []string {
			"Hello. Nice to meet you.",
			"Hi there.",
			"What's up. What do you want to talk about.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hy\b|why.*`, input); matched{
		responses1 := []string {
			"Why do you ask that?",
			"I can't tell you why, it's a secret.",
		}
		randIndex := rand.Intn(len(responses1))
		fmt.Fprintf(w, responses1[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]ho.*`, input); matched{
		responses2 := []string {
			"Im not sure who you are on about. Tell me more about yourself though!",
			"I dont know who that is. Tell me more about you instead!",
		}
		randIndex := rand.Intn(len(responses2))
		fmt.Fprintf(w, responses2[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hat.*`, input); matched{
		responses3 := []string {
			"I'm not really sure to be honest.",
		}
		randIndex := rand.Intn(len(responses3))
		fmt.Fprintf(w, responses3[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hen.*`, input); matched{
		responses3 := []string {
			"I'm not really sure to be honest.",
		}
		randIndex := rand.Intn(len(responses3))
		fmt.Fprintf(w, responses3[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]here.*`, input); matched{
		responses3 := []string {
			"I'm not really sure to be honest.",
		}
		randIndex := rand.Intn(len(responses3))
		fmt.Fprintf(w, responses3[randIndex])
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
