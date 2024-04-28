package app

import (
	"net/http"

	v1 "github.com/brequet/groquette/internal/controller/http/v1"
	"github.com/brequet/groquette/internal/usecase"
	"github.com/brequet/groquette/internal/usecase/webapi"
	"github.com/go-chi/chi"
)

var (
	groqApiKey = "TODO"
)

func Run() {
	httpClient := &http.Client{}

	chatUseCase := usecase.NewChatUseCase(
		webapi.NewGroqClient(httpClient, groqApiKey),
	)

	r := chi.NewRouter()
	v1.NewRouter(r, chatUseCase)

	http.ListenAndServe("localhost:8080", r)
}
