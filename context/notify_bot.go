package context

import (
	"gopkg.in/tucnak/telebot.v2"
)

type NotifyBot struct {
	*telebot.Bot
	users []int
}

func initNotifyBot(token string, users []int) (*NotifyBot, error) {
	n, err := telebot.NewBot(telebot.Settings{Token: token})
	if err != nil {
		return nil, err
	}

	return &NotifyBot{n, users}, nil
}

func (n *NotifyBot) Notify(content string) error {
	for _, v := range n.users {
		_, err := n.Send(&telebot.Chat{ID: int64(v)}, content)
		if err != nil {
			return err
		}
	}

	return nil
}
