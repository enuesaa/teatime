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
			Data:   info,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) List(props ListProps, resp *Result[[]Tea]) error {
	list, err := s.Impl.List(props)
	if err != nil {
		*resp = Result[[]Tea]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[[]Tea]{
			Data:   list,
			HasErr: false,
		}
	}
	return nil
}

func (s *ConnectServer) Act(props ActProps, resp *Result[string]) error {
	message, err := s.Impl.Act(props)
	if err != nil {
		*resp = Result[string]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	} else {
		*resp = Result[string]{
			Data:   message,
			HasErr: false,
		}
	}
	return nil
}
