package chatbot

import (
	"bytes"
	"encoding/json"
	"io"
	"mini-project/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatAI struct{}

var AiPayload = map[string]interface{}{
	"model": "gpt-3.5-turbo",
	"messages": []map[string]string{
		{"role": "system", "content": "Anda seorang ahli dalam bidang hasil pertanian seperti buah, sayuran, dan biji bijian"},
	},
}

func NewChatAI() *ChatAI {
	return &ChatAI{}
}

func (c *ChatAI) HandleChatAi(ctx echo.Context) error {
	chatbotMessages := AiPayload["messages"].([]map[string]string)
	if len(chatbotMessages) == 1 {
		chatbotMessages = append(chatbotMessages, map[string]string{"role": "system", "content": "Anda seorang ahli dalam bidang hasil pertanian seperti buah, sayuran, dan biji bijian"})
	}

	var request AIRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Failed to parse request body", err.Error()))
	}

	chatbotMessages = append(chatbotMessages, map[string]string{"role": "user", "content": request.Messages[0].Content})

	payload := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": []map[string]string{chatbotMessages[len(chatbotMessages)-2], chatbotMessages[len(chatbotMessages)-1]},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to marshal JSON payload", err.Error()))
	}

	resp, err := http.Post("https://wgpt-production.up.railway.app/v1/chat/completions", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to send request", err.Error()))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to read response body", err.Error()))
	}

	return ctx.JSONBlob(resp.StatusCode, body)

}
