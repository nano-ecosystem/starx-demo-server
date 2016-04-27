package chat

type ChatRemote struct {

}

func (c *ChatRemote) Init() {}
func (c *ChatRemote) AfterInit() {}
func (c *ChatRemote) BeforeShutdown() {}
func (c *ChatRemote) Shutdown() {}

func (c *ChatRemote) GetUserList() ([]byte, error) {
	return []byte(`user list`), nil
}
