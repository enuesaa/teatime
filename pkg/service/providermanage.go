package service

import (
	"encoding/json"
	"fmt"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/google/uuid"
)

type Record[T interface{}] struct {
	Id string
	Data T
}

type ProviderConf struct {
	Name string `json:"name"`
	Command string `json:"command"`
}

func NewProviderManageService() *ProviderManageService {
	return &ProviderManageService{}
}

type ProviderManageService struct {}

func (srv *ProviderManageService) List() []Record[ProviderConf] {
	redis := repository.RedisRepository{}
	keys := redis.Keys("provider:")
	list := make([]Record[ProviderConf], 0)
	for _, key := range keys {
		record, err := srv.Describe(key)
		if err != nil {
			return list
		}
		list = append(list, record)
	}
	return list
}

func (srv *ProviderManageService) Create(conf ProviderConf) (string, error) {
	redis := repository.RedisRepository{}
	id, err := srv.GenId()
	if err != nil {
		return "", err
	}
	if err := redis.JsonSet(id, conf); err != nil {
		return "", err
	}
	return id, nil
}

func (srv *ProviderManageService) Describe(id string) (Record[ProviderConf], error) {
	redis := repository.RedisRepository{}
	data, err := redis.JsonGet(id)
	if err != nil {
		return *new(Record[ProviderConf]), err
	}
	var conf ProviderConf
	if err := json.Unmarshal(data, &conf); err != nil {
		return *new(Record[ProviderConf]), err
	}
	record := Record[ProviderConf]{
		Id: id,
		Data: conf,
	}
	return record, nil
}

func (srv *ProviderManageService) Update(id string, conf ProviderConf) (string, error) {
	redis := repository.RedisRepository{}
	if err := redis.JsonSet(id, conf); err != nil {
		return "", err
	}
	return id, nil
}

func (srv *ProviderManageService) Delete(id string) error {
	redis := repository.RedisRepository{}
	if err := redis.JsonDel(id); err != nil {
		return err
	}
	return nil
}

func (srv *ProviderManageService) GenId() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("provider:%s", uid.String()), nil
}
