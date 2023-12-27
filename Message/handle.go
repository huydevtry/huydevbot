package Message

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func HandleMessage(update telego.Update, bot *telego.Bot) {
	fmt.Printf(update.Message.Text)
	message := update.Message.Text
	message = strings.TrimSpace(message)
	chatId := update.Message.Chat.ID

	bot.SendChatAction(
		tu.ChatAction(
			tu.ID(chatId),
			"typing",
		),
	)

	openApiToken := os.Getenv("OPENAPI_TOKEN")
	client := openai.NewClient(openApiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4TurboPreview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "Bạn là một bot của Huy đẹp trai",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	respContent := resp.Choices[0].Message.Content

	//Process message response
	backtickStr := "```"
	if strings.Contains(respContent, backtickStr) {
		respContent = strings.Replace(respContent, backtickStr, "\n"+backtickStr+"\n", -1)
	}
	_, _ = bot.SendMessage(
		tu.Message(
			tu.ID(chatId),
			respContent,
		),
	)
}
