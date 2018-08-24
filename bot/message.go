package bot

// PostMessageURL send message url
var PostMessageURL = "https://slack.com/api/chat.postMessage"

// Message slack message
// https://api.slack.com/methods/chat.postMessage
type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// NewMessage New send message
func NewMessage(channel, text string) *Message {
	return &Message{
		Channel: channel,
		Text:    text,
	}
}
