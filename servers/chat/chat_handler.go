package chat

import "github.com/chrislonng/starx"

type ChatHandler struct {

}

func (c *ChatHandler) Setup() {

}

func (c *ChatHandler) Chating(session *starx.Session, data []byte) error {
	println(string(data))
	session.Response([]byte(`get user chating message`))
	return nil
}

func (c *ChatHandler) UserIsInputing(session *starx.Session, data []byte) error {
	println(string(data))
	return nil
}

