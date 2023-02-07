package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

/*
var botKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("/hi"),
		tu.KeyboardButton("/help"),
		tu.KeyboardButton("/status"),
	),
).WithResizeKeyboard()
*/

func main() {

	bot, err := telego.NewBot(os.Getenv("TGBOT_API"), telego.WithDefaultDebugLogger())
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	botUser, err := bot.GetMe()
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	log.Printf("Work with account %v\n", botUser)

	updates, err := bot.UpdatesViaLongPolling(nil)

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		var messageToUser string

		switch message.Text {
		case "/hi":
			messageToUser = fmt.Sprintf("Привет, %s! Я твой бот-компаньон :)", message.From.FirstName)
		case "/help":
			messageToUser = "Я понимаю команды /hi и /status) А смысл этого?...."
		case "/status":
			messageToUser = "Я в порядке. Работаю... А ты?)"
		}

		_, _ = bot.SendMessage(tu.Messagef(
			tu.ID(message.Chat.ID),
			messageToUser,
		))

	}, th.AnyCommand())

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		chatID := tu.ID(message.Chat.ID)
		_, _ = bot.CopyMessage(
			tu.CopyMessage(chatID, chatID, message.MessageID),
			//.WithReplyMarkup(botKeyboard),
		)
	})

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Start()
}
