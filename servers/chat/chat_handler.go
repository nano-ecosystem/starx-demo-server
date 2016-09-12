package chat

import (
	"github.com/chrislonng/starx/component"
	"github.com/chrislonng/starx/session"
)

type ChatHandler struct {
	component.Base
}

func (c *ChatHandler) Chating(session *session.Session, data []byte) error {
	println(string(data))
	session.Response([]byte(`get user chating message`))
	return nil
}

func (c *ChatHandler) UserIsInputing(session *session.Session, data []byte) error {
	println(string(data))
	return nil
}
