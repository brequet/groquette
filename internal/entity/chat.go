package entity

type ChatModel string

const (
	Llama38b8192     ChatModel = "llama3-8b-8192"
	Llama270b4096    ChatModel = "llama2-70b-4096"
	Mixtral8x7b32768 ChatModel = "mixtral-8x7b-32768"
	Gemma7bit        ChatModel = "gemma-7b-it"
	Llama370b8192    ChatModel = "llama3-70b-8192"
)

func ValidateChatModel(value string) bool {
	switch ChatModel(value) {
	case Llama38b8192, Llama270b4096, Mixtral8x7b32768, Gemma7bit, Llama370b8192:
		return true
	default:
		return false
	}
}
