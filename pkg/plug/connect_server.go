package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Result[Info]) error {
	info, err := s.Impl.Info()
	*resp = NewResult(info, err)
	return nil
}

func (s *ConnectServer) Create(event CreateEvent, resp *Result[Logs]) error {
	message, err := s.Impl.Create(event)
	*resp = NewResult(message, err)
	return nil
}

func (s *ConnectServer) Update(event UpdateEvent, resp *Result[Logs]) error {
	message, err := s.Impl.Update(event)
	*resp = NewResult(message, err)
	return nil
}

func (s *ConnectServer) Delete(event DeleteEvent, resp *Result[Logs]) error {
	message, err := s.Impl.Delete(event)
	*resp = NewResult(message, err)
	return nil
}

func (s *ConnectServer) On(event Event, resp *Result[Logs]) error {
	message, err := s.Impl.On(event)
	*resp = NewResult(message, err)
	return nil
}
