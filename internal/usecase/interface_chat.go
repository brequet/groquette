package usecase

import (
	"github.com/brequet/groquette/internal/entity"
	"github.com/brequet/groquette/internal/usecase/webapi"
)

type (
	GroqClient interface {
		SendRequest(userMessage string, chatModel entity.ChatModel) (*webapi.GroqChatResponse, error)
	}
)
