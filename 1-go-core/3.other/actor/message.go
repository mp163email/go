package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Content string
}

func (m *TextMessage) Type() string {
	return "text"
}
