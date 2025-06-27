package msgs

type (
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	NewUserMessage struct {
		Content string `json:"content"`
	}

	JoinResponseMessage struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)
