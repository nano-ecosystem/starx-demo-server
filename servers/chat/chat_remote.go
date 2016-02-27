package chat

type ChatRemote struct {

}

func (c *ChatRemote) Setup() {

}

func (c *ChatRemote) GetUserList() ([]byte, error) {
	return []byte(`user list`), nil
}
