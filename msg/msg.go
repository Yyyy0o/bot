package msg

type MessageProducer interface {
	GetMessage() []Message
}

type Message struct {
	Type    string
	Content string
}
