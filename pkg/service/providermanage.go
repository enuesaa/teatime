package service

type Record[T interface{}] struct {
	Id string
	Data T
}

type ProviderConf struct {
	Name string
	Command string
}

func NewProviderManageService() *ProviderManageService {
	return &ProviderManageService{}
}

type ProviderManageService struct {}

func (srv *ProviderManageService) List() []Record[ProviderConf] {
	return make([]Record[ProviderConf], 0)
}

func (srv *ProviderManageService) Create(conf ProviderConf) (string, error) {
	return "", nil
}

func (srv *ProviderManageService) Describe(id string) (Record[ProviderConf], error) {
	return *new(Record[ProviderConf]), nil
}

func (srv *ProviderManageService) Update(Id string, conf ProviderConf) (string, error) {
	return "", nil
}

func (srv *ProviderManageService) Delete(Id string) error {
	return nil
}
