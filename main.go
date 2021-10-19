package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	channelID := os.Getenv("TELEGRAM_CHATID")

	b, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Verbose: func() bool { if os.Getenv("DEBUG") == "true" { return true } else { return false }}(),
	})

	if err != nil {
		log.Fatal(err)

		return
	}

	b.Handle("/nick", func(m *tb.Message) {

		// Only if we work on a desired channel
		if m.Chat.Recipient() == channelID {
			log.Printf("Command: /nick, From: [%s], Message: [%s]", m.Sender.Username, m.Payload)
			nickname := strings.TrimSpace(m.Payload)

			if nickname != "" {
				// Check if nickname is already taken
				data, _ := b.AdminsOf(m.Chat)
				
				for _, item := range data {
					s := strings.Split(item.Title, ",")

					for _, v := range s {
						if strings.EqualFold(strings.TrimSpace(v), nickname) {
							log.Printf("Nickname [%s] is already taken", nickname)
							b.Send(m.Sender, fmt.Sprintf("Ник %s уже занят, попробуйте другой", nickname))

							return
						}
					}
				}

				chatMember, err := b.ChatMemberOf(m.Chat, m.Sender)
				chatMember.Rights.CanManageChat = true

				if err != nil {
					log.Printf("Could not get user info for [%s]", m.Sender.Username)

					return
				}

				err = b.Promote(m.Chat, chatMember)

				if err != nil {
					log.Printf("Failed to promote the user [%s]", m.Sender.Username)
				} else {
					// Sleep for a while to give telegram some time to update admins
					time.Sleep(3 * time.Second)

					b.SetAdminTitle(m.Chat, m.Sender, nickname)
				}

				b.Send(m.Sender, fmt.Sprintf("Ваш никнейм установлен: %s", nickname))
			}
		}
	})

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		log.Printf("A new user [%s] joined to the channel", m.Sender.Username)

		b.Send(m.Chat, fmt.Sprintf("Привет @%s 👋 Ты можешь установить свой старый никнейм из mIRC, используя команду `/nick СтарыйНик` прямо в канале", m.Sender.Username), &tb.SendOptions{
			ParseMode: tb.ModeMarkdownV2,
		})
	})

	b.Start()
}