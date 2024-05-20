package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandMessge(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	botMsg := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("接收到文本: %s\n", msg.Text))
	bot.Send(botMsg)
}
func main() {
	bot, err := tgbotapi.NewBotAPI("7059574115:AAGCtTbU8ngXRFUrUrE2P6JtjrhqLpXn4mA")
	if err != nil {
		log.Printf("创建机器人程序失败[ERROR = %s]", err)
		return
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		HandMessge(bot, update.Message)
	}
}
