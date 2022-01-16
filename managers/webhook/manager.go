package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/Vladimir-Urik/AutoVote/logger"
	"net/http"
)

func SendDataToWebhook(content interface{}, embed []Embed, url string) {
	bytesRepresentation, _ := json.Marshal(Payload{
		Content: content,
		Embeds:  embed,
	})
	_, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		logger.Error("Error sending webhook: " + err.Error())
	}
}

type Payload struct {
	Content interface{} `json:"content"`
	Embeds  []Embed     `json:"embeds"`
}

type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
}
