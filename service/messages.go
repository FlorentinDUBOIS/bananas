package service

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

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
		"Buttom",
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
	ResponseType string `json:"response_type,omitempty"`
	Channel      string `json:"channel,omitempty"`
	Text         string `json:"text,omitempty"`
}

// AnswerMessage is the handler of /bananas
func AnswerMessage(ctx *gin.Context) {
	index := rand.Int() % len(answers)
	channel, ok := ctx.GetPostForm("channel_id")
	if !ok {
		log.Error(errors.New("no channel provided"))
		return
	}

	answer := Answer{
		ResponseType: "in_channel",
		Channel:      channel,
		Text:         answers[index],
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("token")))
	ctx.JSON(http.StatusOK, answer)
}
