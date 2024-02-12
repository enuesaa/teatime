package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Init(arg interface{}, resp *error) error {
	*resp = s.Impl.Init()
	return nil
}
func (s *ConnectServer) Info(arg interface{}, resp *Info) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ConnectServer) List(arg interface{}, resp *[]string) error {
	*resp = s.Impl.List()
	return nil
}
func (s *ConnectServer) Get(id string, resp *Row) error {
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
