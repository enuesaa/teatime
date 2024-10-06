package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Result[Info]) error {
	info, err := s.Impl.Info()
	*resp = NewResult(info, err)
	return nil
}

func (s *ConnectServer) On(props OnProps, resp *Result[bool]) error {
	is, err := s.Impl.On(props)
	*resp = NewResult(is, err)
	return nil
}

func (s *ConnectServer) Logs(arg interface{}, resp *Result[string]) error {
	logs, err := s.Impl.Logs()
	*resp = NewResult(logs, err)
	return nil
}
