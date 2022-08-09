package messages

type ReplyAddress struct {
	URL string
}

// RequestReply defines a request reply message.
type RequestReply struct {
	Data         interface{}
	ReplyAddress ReplyAddress
}
