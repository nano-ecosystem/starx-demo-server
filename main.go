package main

import (
	"github.com/chrislonng/starx"
	"starx-demo-server/servers/gate"
	"starx-demo-server/servers/chat"
	"os"
)

func main() {
	starx.Set("gate", func(){
		starx.Register(new(gate.LoginHandler))
	});
	starx.Set("chat", func(){
		starx.Register(new(chat.ChatHandler))
		starx.Register(new(chat.ChatRemote))
	});

	starx.SetServerID(os.Args[1])
	starx.Run()
}
