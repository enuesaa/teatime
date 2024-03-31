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
func (s *ConnectServer) Get(id string, resp *GetResult) error {
	*resp = s.Impl.Get(id)
	return nil
}
func (s *ConnectServer) Set(row Row, resp *error) error {
	*resp = s.Impl.Set(row)
	return nil
}
func (s *ConnectServer) Del(id string, resp *error) error {
	*resp = s.Impl.Del(id)
	return nil
}
func (s *ConnectServer) GetCard(name string, resp *GetCardResult) error {
	*resp = s.Impl.GetCard(name)
	return nil
}
