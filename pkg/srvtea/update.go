package srvtea

func (srv *Srv) Update(teaId string, tea Tea) (string, error) {
	if err := srv.Delete(teaId); err != nil {
		return teaId, err
	}
	tea.Id = teaId
	if _, err := srv.repos.DB.Create(srv.CollectionName(), tea); err != nil {
		return teaId, err
	}
	return teaId, nil
}
