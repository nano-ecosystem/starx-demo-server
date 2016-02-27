package main

import (
	"github.com/chrislonng/starx"
	"starx-demo-server/servers/gate"
	"starx-demo-server/servers/chat"
)

func main() {
	starx.Set("gate", func(){
		starx.Handler(new(gate.LoginHandler))
	});
	starx.Set("chat", func(){
		starx.Handler(new(chat.ChatHandler))
		starx.Remote(new(chat.ChatRemote))
	});
	starx.Start()
}
