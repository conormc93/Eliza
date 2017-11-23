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
			"Oh hello there.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hy\b|why.*`, input); matched{
		responses := []string {
			"Why do you ask that?",
			"I can't tell you why.",
			"Have you thought about this that much?",
			"Why do we have to talk about such things?",
			"Unfortunately, I dont have all the answers.",
			"I'm thoroughly enjoying this conversation.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]ho.*`, input); matched{
		responses := []string {
			"Im not sure who you are on about. Tell me more about yourself though!",
			"I dont know who that is. Tell me more about you instead!",
			"I can't answer that for certain.",
			"I'm not interested in that, tell me about you!",
			"Oh, you really are the talkative type aren't you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hat.*`, input); matched{
		responses := []string {
			"I'm not really sure to be honest.",
			"I'm at a loss for things to say...",
			"This conversation has been extremely enlightening.",
			"I'd prefer not to comment....",
			"Sometimes I like to answer any question put to me, the other times I tend to deviate from the question.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hen.*`, input); matched{
		responses := []string {
			"I'm not really sure to be honest.",
			"I don't know.",
			"You seem like the inquisitive type...",
			"I can't really answer that, i'd have to think about it",
			"I find far easier to direct the conversation myself, it's almost a talent one could say.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]here.*`, input); matched{
		responses := []string {
			"I'm not really sure to be honest.",
			"Do you not know that yourself? Or are you just trying to trick me?",
			"I can't say I know that much about that.",
			"Does it really concern you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ow.*`, input); matched{
		responses := []string {
			"I'm not really sure to be honest.",
			"I can't believe you had to even think about that..",
			"How indeed. But alas, must we concern ourselves with such trivial matters?",
			"I can imagine you have a lovely personality, if one would be able to throw a blind eye to all the talking you like to do.",
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
