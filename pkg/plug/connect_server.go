package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Result[Info]) error {
	info, err := s.Impl.Info()
	*resp = NewResult(info, err)
	return nil
}

func (s *ConnectServer) On(event Event, resp *Result[string]) error {
	message, err := s.Impl.On(event)
	*resp = NewResult(message, err)
	return nil
}
