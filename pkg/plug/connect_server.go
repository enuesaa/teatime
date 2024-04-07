package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Result[Info]) error {
	info, err := s.Impl.Info()
	if err != nil {
		*resp = Result[Info]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[Info]{
			Data: info,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) List(arg interface{}, resp *Result[[]string]) error {
	list, err := s.Impl.List()
	if err != nil {
		*resp = Result[[]string]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[[]string]{
			Data: list,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) Get(teaid string, resp *Result[Tea]) error {
	tea, err := s.Impl.Get(teaid)
	if err != nil {
		*resp = Result[Tea]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[Tea]{
			Data: tea,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) Set(tea Tea, resp *Result[bool]) error {
	if err := s.Impl.Set(tea); err != nil {
		*resp = Result[bool]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[bool]{
			Data: true,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) Del(teaid string, resp *Result[bool]) error {
	if err := s.Impl.Del(teaid); err != nil {
		*resp = Result[bool]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[bool]{
			Data: true,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) GetCard(name string, resp *Result[Card]) error {
	card, err := s.Impl.GetCard(name)
	if err != nil {
		*resp = Result[Card]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[Card]{
			Data: card,
			HasErr: false,
		}
	}
	return nil
}
