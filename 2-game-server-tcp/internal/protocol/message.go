package protocol

// 定义了一个类-消息结构体
type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}
