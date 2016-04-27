package gate

import "github.com/chrislonng/starx"

type LoginHandler struct {

}

func(h *LoginHandler) Init() {
	println("login handler inited")
}
func(h *LoginHandler) AfterInit() {}
func(h *LoginHandler) BeforeShutdown() {}
func(h *LoginHandler) Shutdown() {}

func(h *LoginHandler) Login(session *starx.Session, data []byte) error {
	println(string(data))
	session.Push("OnServerNotify2", []byte(`push message test, sent by gate.LoginHandler.Login`))
	session.Response([]byte(`get login request, this is response message`))
	return nil
}

func(h *LoginHandler) NotifyTest(session *starx.Session, data []byte) error {
	println(string(data))
	session.Push("OnServerNotify", []byte(`push message test, sent by gate.LoginHandler.NotifyTest`))
	return nil
}

func(h *LoginHandler) GetUserInfomation(session *starx.Session, data []byte) error {
	println(string(data))
	return nil
}

func(h *LoginHandler) GetLevel(session *starx.Session, data []byte) error {
	println(string(data))
	return nil
}

func(h *LoginHandler) GetLastestRank(session *starx.Session, data []byte) error {
	println(string(data))
	return nil
}

func(h *LoginHandler) TestRpc(session *starx.Session, data []byte) error {
	println(string(data))
	res, err := session.RPC("chat.ChatRemote.GetUserList")
	if err != nil {
		return err
	}
	session.Response(res)
	return nil
}
