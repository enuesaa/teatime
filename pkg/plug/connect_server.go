package plug

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *InfoResult) error {
	info, err := s.Impl.Info()
	if err != nil {
		*resp = NewInfoErr(err)
	} else {
		*resp = NewInfoResult(info)
	}
	return nil
}

func (s *ConnectServer) List(arg interface{}, resp *ListResult) error {
	list, err := s.Impl.List()
	if err != nil {
		*resp = NewListErr(err)
	} else {
		*resp = NewListResult(list)
	}
	return nil
}

func (s *ConnectServer) Get(teaid string, resp *GetResult) error {
	tea, err := s.Impl.Get(teaid)
	if err != nil {
		*resp = NewGetErr(err)
	} else {
		*resp = NewGetResult(tea)
	}
	return nil
}

func (s *ConnectServer) Set(tea Tea, resp *SetResult) error {
	if err := s.Impl.Set(tea); err != nil {
		*resp = NewSetErr(err)
	} else {
		*resp = NewSetResult()
	}
	return nil
}

func (s *ConnectServer) Del(teaid string, resp *DelResult) error {
	if err := s.Impl.Del(teaid); err != nil {
		*resp = NewDelErr(err)
	} else {
		*resp = NewDelResult()
	}
	return nil
}

func (s *ConnectServer) GetCard(name string, resp *GetCardResult) error {
	card, err := s.Impl.GetCard(name)
	if err != nil {
		*resp = NewGetCardErr(err)
	} else {
		*resp = NewGetCardResult(card)
	}
	return nil
}
