package main

import (
	"log"
	"os"
	"secret-santa/instances"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	config := instances.GetConfig()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = config.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Обработка входящих сообщений
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyMarkup = createKeyboard()

		switch update.Message.Text {
		case "/start":
			msg.Text = "Привет! Я бот для игры в тайного санту :)\n Нажми на кнопку:"
		case "Создать комнату":
			msg.Text = "Ты нажал на кнопку 1!"
		case "Присоединится к комнате":
			msg.Text = "Ты нажал на кнопку 2!"
		default:
			msg.Text = "Я не знаю, что ответить на это сообщение 😕"
		}

		bot.Send(msg)
	}
}

func createKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Создать комнату"),
			tgbotapi.NewKeyboardButton("Присоединится к комнате"),
		),
	)

	return keyboard
}
