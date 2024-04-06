package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *InfoResult) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ConnectServer) List(arg interface{}, resp *ListResult) error {
	*resp = s.Impl.List()
	return nil
}
func (s *ConnectServer) Get(teaid string, resp *GetResult) error {
	*resp = s.Impl.Get(teaid)
	return nil
}
func (s *ConnectServer) Set(tea Tea, resp *SetResult) error {
	*resp = s.Impl.Set(tea)
	return nil
}
func (s *ConnectServer) Del(teaid string, resp *DelResult) error {
	*resp = s.Impl.Del(teaid)
	return nil
}
func (s *ConnectServer) GetCard(name string, resp *GetCardResult) error {
	*resp = s.Impl.GetCard(name)
	return nil
}
