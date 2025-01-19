package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ConnectServer struct {
	Impl ProviderInterface
	logger Logger
}

func (s *ConnectServer) startup() bool {
	if err := s.Impl.OnStartup(); err != nil {
		s.logger.Err(err)
		return false
	}
	return true
}

func (s *ConnectServer) shutdown() bool {
	if err := s.Impl.OnShutdown(); err != nil {
		s.logger.Err(err)
		return false
	}
	return true
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
	s.logger.Info("info")
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	info, err := s.Impl.Info()
	*resp = ConnectResult[Info]{info, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) List(props ListProps, resp *ConnectResult[[]db.Tea]) error {
	s.logger.Info("list: %+v", props)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	list, err := s.Impl.List(props)
	*resp = ConnectResult[[]db.Tea]{list, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Get(props GetProps, resp *ConnectResult[db.Tea]) error {
	s.logger.Info("get: %+v", props)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.Impl.Get(props)
	*resp = ConnectResult[db.Tea]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Create(props CreateProps, resp *ConnectResult[string]) error {
	s.logger.Info("create: %+v", props)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.Impl.Create(props)
	*resp = ConnectResult[string]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Update(props UpdateProps, resp *ConnectResult[string]) error {
	s.logger.Info("update: %+v", props)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.Impl.Update(props)
	*resp = ConnectResult[string]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Delete(props DeleteProps, resp *ConnectResult[bool]) error {
	s.logger.Info("delete: %+v", props)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.Impl.Delete(props)
	*resp = ConnectResult[bool]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}
