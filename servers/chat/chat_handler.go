package chat

import "github.com/chrislonng/starx"

type ChatHandler struct {

}

func (c *ChatHandler) Init() {}
func (c *ChatHandler) AfterInit() {}
func (c *ChatHandler) BeforeShutdown() {}
func (c *ChatHandler) Shutdown() {}

func (c *ChatHandler) Chating(session *starx.Session, data []byte) error {
	println(string(data))
	session.Response([]byte(`get user chating message`))
	return nil
}

func (c *ChatHandler) UserIsInputing(session *starx.Session, data []byte) error {
	println(string(data))
	return nil
}

