package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

const answerCallback = "https://slack.com/api/chat.postMessage"

var (
	answers = []string{
		"Bello",
		"Poopaye",
		"Tank yu",
		"Bi-do",
		"Pwede na?",
		"Para tu",
		"Sa la ka!",
		"Tatata bala tu",
		"Underwear...",
		"Me want banana",
		"Luk at tu",
		"Tulaliloo ti amo",
		"Hana",
		"Dul",
		"Sae",
		"Bottom - Buttom",
		"Butt",
		"Chasy",
		"Kampai",
		"BEE DO BEE DO BEE DO",
		"Muak Muak Muak",
		"La boda",
		"Stupa",
		"Baboi",
		"Bananonina",
		"Po ka",
	}
)

// Answer is the response structure
type Answer struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// AnswerMessage is the handler of /bananas
func AnswerMessage(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
	index := rand.Int() % len(answers)
	channel, ok := ctx.GetPostForm("channel_id")
	if !ok {
		log.Error(errors.New("no channel provided"))
		return
	}

	answer := Answer{
		Channel: channel,
		Text:    answers[index],
	}

	io, err := json.Marshal(answer)
	if err != nil {
		log.Error(err)
		return
	}

	req, err := http.NewRequest(answerCallback, "application/json", bytes.NewReader(io))
	if err != nil {
		log.Error(err)
		return
	}

	req.Header.Set("Authorization", viper.GetString("token"))
	client := new(http.Client)
	_, err = client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
}
