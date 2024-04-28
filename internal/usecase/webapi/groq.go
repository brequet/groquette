package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/brequet/groquette/internal/entity"
)

type GroqClient struct {
	url        string
	httpClient *http.Client
	apiKey     string
}

func NewGroqClient(httpClient *http.Client, apiKey string) *GroqClient {
	return &GroqClient{
		url:        "https://api.groq.com/openai",
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

// Request types
type (
	RequestBody struct {
		Messages []Message `json:"messages"`
		Model    string    `json:"model"`
	}

	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
)

// Reponse types
type (
	GroqChatResponse struct {
		ID                string   `json:"id"`
		Object            string   `json:"object"`
		Created           int64    `json:"created"`
		Model             string   `json:"model"`
		Choices           []Choice `json:"choices"`
		Usage             Usage    `json:"usage"`
		SystemFingerprint string   `json:"system_fingerprint"`
		XGroq             XGroq    `json:"x_groq"`
	}

	Choice struct {
		Index        int         `json:"index"`
		Message      Message     `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	}

	Usage struct {
		PromptTokens     int     `json:"prompt_tokens"`
		PromptTime       float64 `json:"prompt_time"`
		CompletionTokens int     `json:"completion_tokens"`
		CompletionTime   float64 `json:"completion_time"`
		TotalTokens      int     `json:"total_tokens"`
		TotalTime        float64 `json:"total_time"`
	}

	XGroq struct {
		ID string `json:"id"`
	}
)

func (client *GroqClient) SendRequest(userMessage string, chatModel entity.ChatModel) (*GroqChatResponse, error) {
	requestBody := RequestBody{
		Messages: []Message{
			{
				Role:    "user",
				Content: userMessage,
			},
		},
		Model: string(chatModel),
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", client.url+"/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var response GroqChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &response, nil
}
