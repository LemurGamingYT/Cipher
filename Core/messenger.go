package core

type Messenger struct {
	storeMessage string
	storeValue any
}

func NewMessenger(msg string, storeValue any) *Messenger {
	return &Messenger{msg, storeValue}
}
