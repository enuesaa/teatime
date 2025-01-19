package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ConnectServer struct {
	impl ProviderInterface
	logger Logger
}

func (s *ConnectServer) startup() bool {
	if err := s.impl.OnStartup(); err != nil {
		s.logger.Err(err)
		return false
	}
	return true
}

func (s *ConnectServer) shutdown() bool {
	if err := s.impl.OnShutdown(); err != nil {
		s.logger.Err(err)
		return false
	}
	return true
}

func (s *ConnectServer) OnStartup(arg interface{}, resp *error) error {
	*resp = s.impl.OnStartup()
	return nil
}

func (s *ConnectServer) OnShutdown(arg interface{}, resp *error) error {
	*resp = s.impl.OnShutdown()
	return nil
}

func (s *ConnectServer) Info(arg interface{}, resp *ConnectResult[Info]) error {
	s.logger.Info("info")
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	info, err := s.impl.Info()
	*resp = ConnectResult[Info]{info, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) List(args ListArgs, resp *ConnectResult[[]db.Tea]) error {
	s.logger.Info("list: %+v", args)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	list, err := s.impl.List(args)
	*resp = ConnectResult[[]db.Tea]{list, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Get(args GetArgs, resp *ConnectResult[db.Tea]) error {
	s.logger.Info("get: %+v", args)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.impl.Get(args)
	*resp = ConnectResult[db.Tea]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Create(args CreateArgs, resp *ConnectResult[string]) error {
	s.logger.Info("create: %+v", args)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.impl.Create(args)
	*resp = ConnectResult[string]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Update(args UpdateArgs, resp *ConnectResult[string]) error {
	s.logger.Info("update: %+v", args)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	data, err := s.impl.Update(args)
	*resp = ConnectResult[string]{data, err}
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}

func (s *ConnectServer) Delete(args DeleteArgs, resp *error) error {
	s.logger.Info("delete: %+v", args)
	if !s.startup() {
		return nil
	}
	defer s.shutdown()

	err := s.impl.Delete(args)
	*resp = err
	if err != nil {
		s.logger.Err(err)
	}
	return nil
}
