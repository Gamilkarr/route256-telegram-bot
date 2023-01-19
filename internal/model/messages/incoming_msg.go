package messages

// Интерфейс отправителя сообщений
type MessageSender interface {
	SendMessenge(text string, userID int64) error
}

// Модель для обработки входящих сообщений
type Model struct {
	tgClient MessageSender
}

// Эта функция создает новый объект модели и возвращает ссылку на него
func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

// Тип сообщение
type Message struct {
	Text   string
	UserID int64
}

// Метод для обработки входящего сообщения
// !Непонятно почему тут ресивер не m
func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		s.tgClient.SendMessenge("Hello!", msg.UserID)
		return nil
	}
	s.tgClient.SendMessenge("Sorry, i don't know this command", msg.UserID)
	return nil
}
