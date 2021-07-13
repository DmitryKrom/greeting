package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name given")
	}
	message := fmt.Sprintf(sentenceToChoose(), name)

	return message, nil
}

func init (){
	rand.Seed(time.Now().UnixNano())
	rand.Seed(1)
}

func sentenceToChoose()string{
	sentences := []string{
		"Hi %v, great to see you!",
		"%v, hi! how you doing?",
		"I am glad to see you, %v!",
		"Hello, %v, nice to meet you!"
	}
	sentence := sentences[rand.Intn(len(sentences))]
	return sentence
}
