package main

import (
	"fmt"

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

type ProviderService struct {}

func (srv *ProviderService) List() []Record[ProviderConf] {
	redis := NewRedis()
	keys := redis.Keys("provider:*")
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

func (srv *ProviderService) Create(conf ProviderConf) (string, error) {
	id, err := srv.GenId()
	if err != nil {
		return "", err
	}
	fmt.Println(id)
	return id, nil
}

func (srv *ProviderService) Describe(id string) (Record[ProviderConf], error) {
	record := Record[ProviderConf]{
		Id: id,
	}
	return record, nil
}

func (srv *ProviderService) Update(id string, conf ProviderConf) (string, error) {
	return id, nil
}

func (srv *ProviderService) Delete(id string) error {
	return nil
}

func (srv *ProviderService) GenId() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("provider:%s", uid.String()), nil
}
