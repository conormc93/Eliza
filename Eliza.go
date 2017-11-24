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

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ello|[Hh]ey|[Hh]i\b.*`, input); matched{
		responses := []string {
			"Hello. Nice to meet you.",
			"Hi there.",
			"What's up. What do you want to talk about.",
			"Oh hello there.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	} else 

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
			"Does that question interest you ?",
			"Are such questions much on your mind?",
			"You seem like the inquisitive type...",
			"What answer would please you most ?",
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
			"Oh, sure, you know yourself.",
			"Must we concern ourselves with such trivial matters?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*I am|I'm*`, input); matched{
		responses := []string {
			"Oh, and why is that?",
			"Why are you like that?",
			"I dont understand that.",
			"Tell me, why is that the case?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	
	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]as I*`, input); matched{
		responses := []string {
			"What if you were?",
			"Do you think you were?",
			"Were you?",
			"What would it mean if you were?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI was\b*`, input); matched{
		responses := []string {
			"Were you really?",
			"Why do you tell me that you were?",
			"Were you?",
			"Perhaps I already knew you were?...",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI am|I'm\b*`, input); matched{
		responses := []string {
			"Are you?",
			"Did you think telling me this would make a difference to me or you?",
			"Talking to me can cause these things.",
			"Do you truly believe that?",
			"Why is that?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI do\b*`, input); matched{
		responses := []string {
			"Good for you.",
			"I'm surprised.",
			"Why do you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	if matched, _ := regexp.MatchString(`(?i).*\bI dont|I don't\b*`, input); matched{
		responses := []string {
			"I thought you would.",
			"Hard to believe.",
			"Why don't you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Pp]erhaps\b*`, input); matched{
		responses := []string {
			"You don't seem quite certain.",
			"Why the uncertain tone",
			"You aren't sure?",
			"Don't you know",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}
	if matched, _ := regexp.MatchString(`(?i).*you are|you're*`, input); matched{
		responses := []string {
			"You don't seem quite certain though...",
			"Do you sometimes wish that you were?",
			"You know a lot about me is that it",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Yy]ou\b*`, input); matched{
		responses := []string {
			"We were discussing you, not me.",
			"You're not really talking about me, are you ?",
			"What are your feelings now ?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Yy]es\b*`, input); matched{
		responses := []string {
			"You seem to be quite positive.",
			"You are sure.",
			"I see.",
			"I understand.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Nn]o\b*`, input); matched{
		responses := []string {
			"Are you saying no just to be negative?",
			"You are being a bit negative.",
			"Why not?",
			"Why 'no'?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Bb]ecause\b*`, input); matched{
		responses := []string {
			"Is that the real reason?",
			"Don't any other reasons come to mind ?",
			"Does that reason seem to explain anything else?",
			"What other reasons might there be?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ww]hy don't you\b*`, input); matched{
		responses := []string {
			"Do you believe I don't?",
			"Perhaps I will, in good time...",
			"Would that please you?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\balwaysb*`, input); matched{
		responses := []string {
			"Really? Always?",
			"Can't you think of a specific example",
			"When?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ll]ike\b*`, input); matched{
		responses := []string {
			"Really? Like why is that?",
			"I would never of guessed that.",
			"Since when?",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ll]ove\b*`, input); matched{
		responses := []string {
			"Really? Like why do you?",
			"That's amazing. Tell me more.",
			"I wish I could also feel that too.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bI would\b*`, input); matched{
		responses := []string {
			"Really? Why would you?",
			"Why would you?",
			"You're telling me you would? Interesting. Go on.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Hh]ate\b*`, input); matched{
		responses := []string {
			"Really? Why is that?",
			"I would never of guessed that.",
			"Since when?",
			"Hate is such a strong word.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\bdon't|do not\b*`, input); matched{
		responses := []string {
			"Why don't you?",
			"I can't believe you dont.",
			"I would of thought you would the oppposite for some reason...",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	}

	if matched, _ := regexp.MatchString(`(?i).*\b[Ii]t is\b*`, input); matched{
		responses := []string {
			"Is it really?",
			"I would never of guessed that.",
			"I can't believe that that is the case.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
		return
	} else{
		responses := []string {
			"Oh, can't we talk about something else?",
			"I think we should talk about something else.",
			"Tell me more about you.",
			"I'm not really sure what you mean.",
			"I can't say I fully grasp what you are saying.",
		}
		randIndex := rand.Intn(len(responses))
		fmt.Fprintf(w, responses[randIndex])
	}
		

}

func main() {
	rand.Seed(time.Now().Unix())

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/elizaResponse", elizaResponse)
	http.ListenAndServe(":8080", nil)
}
