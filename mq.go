package connector

type MessageHandler func(IMessage)

type IBroker interface {
	Connect(url string) error
	Close() error
	Publish(topic string, v any) error
	Subscribe(topic string, handler MessageHandler) error
	OpenChannel(name string) (IChannel, error)
}

type IChannel interface {
	Close() error
	Publish(topic string, v any) error
	Subscribe(topic string, handler MessageHandler) error
}

type IMessage interface {
	//Acknowledge consumed message	/ 确认消费该消息
	Ack() error
	//Requeued to be consumed again / 重新排队并再次消费
	Requeue() error
	//Reject to consume again / 拒绝再次消费
	Reject() error
	//Unmarshal To target
	Unmarshal(v any) error
	//MIME content type
	ContentType() string
	// Body raw data
	Body() []byte
	//Message ID
	ID() string
}
