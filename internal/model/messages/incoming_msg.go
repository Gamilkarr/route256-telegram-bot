package messages

import (
	"fmt"
	"github.com/Gamilkarr/route256-telegram-bot/internal/repository"
	"strconv"
	"strings"
	"time"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		s.tgClient.SendMessage("Привет", msg.UserID)
		return nil
	} else if strings.HasPrefix(msg.Text, "/spent") {
		atrs := strings.Split(msg.Text, " ")
		if len(atrs) != 3 {
			s.tgClient.SendMessage("Неверное количество аргументов", msg.UserID)
			return nil
		}
		date := time.Now().UTC().Round(time.Hour * 24)
		category := atrs[1]
		sum, err := strconv.Atoi(atrs[2])
		if err != nil {
			s.tgClient.SendMessage("Неверно указана сумма", msg.UserID)
			return nil
		}
		repository.AddCost(date, category, sum)
		s.tgClient.SendMessage("Готово", msg.UserID)
		return nil
	} else if strings.HasPrefix(msg.Text, "/get_spends") {
		atrs := strings.Split(msg.Text, " ")
		if len(atrs) != 2 {
			s.tgClient.SendMessage("Неверное количество аргументов", msg.UserID)
			return nil
		}
		end := time.Now()
		var begin time.Time
		switch atrs[1] {
		case "неделя":
			begin = end.AddDate(0, 0, -7)
		case "месяц":
			begin = end.AddDate(0, -1, 0)
		case "год":
			begin = end.AddDate(-1, 0, 0)
		default:
			s.tgClient.SendMessage("Некорректный период", msg.UserID)
			return nil
		}
		spending, _ := repository.GetCostsForPeriod(begin, end)
		s.tgClient.SendMessage(fmt.Sprintf("Траты за %s:", atrs[1]), msg.UserID)
		for category, sum := range spending {
			s.tgClient.SendMessage(fmt.Sprintf("%s: %d", category, sum), msg.UserID)
		}
		return nil
	}
	s.tgClient.SendMessage("Прости, я не знаю такой команды", msg.UserID)
	return nil
}
