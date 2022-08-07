package messages

type ReplyAddress struct {
	URL string
}

// RequestReply
type RequestReply struct {
	Data         interface{}
	ReplyAddress ReplyAddress
}
