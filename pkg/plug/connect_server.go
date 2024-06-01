package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Result[Info]) error {
	info, err := s.Impl.Info()
	*resp = NewResult(info, err)
	return nil
}

func (s *ConnectServer) List(props ListProps, resp *Result[[]Tea]) error {
	list, err := s.Impl.List(props)
	*resp = NewResult(list, err)
	return nil
}

func (s *ConnectServer) Act(props ActProps, resp *Result[string]) error {
	message, err := s.Impl.Act(props)
	*resp = NewResult(message, err)
	return nil
}
