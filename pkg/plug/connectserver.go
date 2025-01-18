package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *ConnectResult[Info]) error {
	info, err := s.Impl.Info()
	*resp = ConnectResult[Info]{info, err}
	return nil
}

func (s *ConnectServer) List(props ListProps, resp *ConnectResult[[]db.Tea]) error {
	list, err := s.Impl.List(props)
	*resp = ConnectResult[[]db.Tea]{list, err}
	return nil
}

func (s *ConnectServer) Get(props GetProps, resp *ConnectResult[db.Tea]) error {
	data, err := s.Impl.Get(props)
	*resp = ConnectResult[db.Tea]{data, err}
	return nil
}

func (s *ConnectServer) Create(props CreateProps, resp *ConnectResult[string]) error {
	data, err := s.Impl.Create(props)
	*resp = ConnectResult[string]{data, err}
	return nil
}

func (s *ConnectServer) Update(props UpdateProps, resp *ConnectResult[string]) error {
	data, err := s.Impl.Update(props)
	*resp = ConnectResult[string]{data, err}
	return nil
}

func (s *ConnectServer) Delete(props DeleteProps, resp *ConnectResult[bool]) error {
	data, err := s.Impl.Delete(props)
	*resp = ConnectResult[bool]{data, err}
	return nil
}
