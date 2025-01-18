package srvlog

func (srv *Srv) Create(message LogMessage) (string, error) {
	query := srv.repos.DB.Logs()

	return query.Create(message)
}
