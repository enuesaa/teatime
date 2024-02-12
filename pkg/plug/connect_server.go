package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Info) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ConnectServer) List(arg ListArg, resp *[]string) error {
	*resp = s.Impl.List(arg)
	return nil
}
func (s *ConnectServer) Get(arg GetArg, resp *Row) error {
	*resp = s.Impl.Get(arg)
	return nil
}
func (s *ConnectServer) Set(arg SetArg, resp *error) error {
	*resp = s.Impl.Set(arg)
	return nil
}
func (s *ConnectServer) Del(arg DelArg, resp *error) error {
	*resp = s.Impl.Del(arg)
	return nil
}
