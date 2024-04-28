package usecase

import (
	"fmt"

	"github.com/brequet/groquette/internal/entity"
)

type ChatUseCase struct {
	groqClient GroqClient
}

func NewChatUseCase(groqClient GroqClient) *ChatUseCase {
	return &ChatUseCase{
		groqClient: groqClient,
	}
}

func (uc *ChatUseCase) SendMessage(message string, chatModel entity.ChatModel) (string, error) {
	response, err := uc.groqClient.SendRequest(message, chatModel)
	if err != nil {
		return "", fmt.Errorf("send groq request: %w", err)
	}
	return response.Choices[0].Message.Content, nil
}
