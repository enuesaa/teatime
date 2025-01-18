package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Result[Info]) error {
	info, err := s.Impl.Info()
	*resp = NewResult(info, err)
	return nil
}

func (s *ConnectServer) List(props ListProps, resp *Result[[]db.Tea]) error {
	list, err := s.Impl.List(props)
	*resp = NewResult(list, err)
	return nil
}

func (s *ConnectServer) Get(props GetProps, resp *Result[db.Tea]) error {
	data, err := s.Impl.Get(props)
	*resp = NewResult(data, err)
	return nil
}

func (s *ConnectServer) Create(props CreateProps, resp *Result[string]) error {
	data, err := s.Impl.Create(props)
	*resp = NewResult(data, err)
	return nil
}

func (s *ConnectServer) Update(props UpdateProps, resp *Result[string]) error {
	data, err := s.Impl.Update(props)
	*resp = NewResult(data, err)
	return nil
}

func (s *ConnectServer) Delete(props DeleteProps, resp *Result[bool]) error {
	data, err := s.Impl.Delete(props)
	*resp = NewResult(data, err)
	return nil
}
