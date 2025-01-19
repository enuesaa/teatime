package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) OnStartup(arg interface{}, resp *error) error {
	*resp = s.Impl.OnStartup()
	return nil
}

func (s *ConnectServer) OnShutdown(arg interface{}, resp *error) error {
	*resp = s.Impl.OnShutdown()
	return nil
}

func (s *ConnectServer) Info(arg interface{}, resp *ConnectResult[Info]) error {
	if err := s.Impl.OnStartup(); err != nil {
		return err
	}

	info, err := s.Impl.Info()
	*resp = ConnectResult[Info]{info, err}

	return s.Impl.OnShutdown()
}

func (s *ConnectServer) List(props ListProps, resp *ConnectResult[[]db.Tea]) error {
	if err := s.Impl.OnStartup(); err != nil {
		return err
	}

	list, err := s.Impl.List(props)
	*resp = ConnectResult[[]db.Tea]{list, err}

	return s.Impl.OnShutdown()
}

func (s *ConnectServer) Get(props GetProps, resp *ConnectResult[db.Tea]) error {
	if err := s.Impl.OnStartup(); err != nil {
		return err
	}

	data, err := s.Impl.Get(props)
	*resp = ConnectResult[db.Tea]{data, err}

	return s.Impl.OnShutdown()
}

func (s *ConnectServer) Create(props CreateProps, resp *ConnectResult[string]) error {
	if err := s.Impl.OnStartup(); err != nil {
		return err
	}

	data, err := s.Impl.Create(props)
	*resp = ConnectResult[string]{data, err}

	return s.Impl.OnShutdown()
}

func (s *ConnectServer) Update(props UpdateProps, resp *ConnectResult[string]) error {
	if err := s.Impl.OnStartup(); err != nil {
		return err
	}

	data, err := s.Impl.Update(props)
	*resp = ConnectResult[string]{data, err}

	return s.Impl.OnShutdown()
}

func (s *ConnectServer) Delete(props DeleteProps, resp *ConnectResult[bool]) error {
	if err := s.Impl.OnStartup(); err != nil {
		return err
	}

	data, err := s.Impl.Delete(props)
	*resp = ConnectResult[bool]{data, err}

	return s.Impl.OnShutdown()
}
