package Message

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"strings"
)

func HandleMessage(update telego.Update, bot *telego.Bot) {
	fmt.Printf(update.Message.Text)
	message := update.Message.Text
	chatId := update.Message.Chat.ID

	//    fmt.Println(strings.Contains("GeeksforGeeks", "for"))
	if strings.Contains(message, "hello") {
		err := bot.SendChatAction(
			tu.ChatAction(
				tu.ID(chatId),
				"typing",
			),
		)
		if err != nil {
			return
		}
		sentMess, _ := bot.SendMessage(
			tu.Message(
				tu.ID(chatId),
				"Lô con mẹ mày",
			),
		)
		fmt.Printf("Sent Message: %v\n", sentMess)

	}
}
