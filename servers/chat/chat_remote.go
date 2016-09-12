package chat

type ChatRemote struct {

}

func (c *ChatRemote) Init() {}
func (c *ChatRemote) AfterInit() {}
func (c *ChatRemote) BeforeShutdown() {}
func (c *ChatRemote) Shutdown() {}

func (c *ChatRemote) GetUserList() (string, error) {
	return `user list`, nil
}
