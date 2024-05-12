package chatbot

import (
	"bytes"
	"encoding/json"
	"io"
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

func (c *ChatAI) HandleChatCompletion(ctx echo.Context) error {
	var request AIRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to parse request body"})
	}

	payload := AIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: request.Messages,
		Stream:   true,
	}

	respBody, statusCode, err := sendRequestToChatbot(payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSONBlob(statusCode, respBody)
}

func sendRequestToChatbot(payload AIRequest) ([]byte, int, error) {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resp, err := http.Post("https://wgpt-production.up.railway.app/v1/chat/completions", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return body, resp.StatusCode, nil
}
