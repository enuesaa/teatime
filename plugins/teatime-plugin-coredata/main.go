package main

import (
	"fmt"
	"strings"

	"github.com/enuesaa/teatime/pkg/plug"
)

func main() {
	plug.Serve(&Handler{})
}

type Handler struct {}

func (h *Handler) Init() error {
	redis := NewRedis()
	return redis.Run()
}

func (s *Handler) Info() plug.Result[plug.Info] {
	info := plug.Info{
		Name: "coredata",
		Description: "coredata provider",
	}
	return plug.NewResult[plug.Info](info)
}

func (h *Handler) List() plug.Result[[]string] {
	redis := NewRedis()
	list, err := redis.Keys("*")
	if err != nil {
		return plug.NewErrResult[[]string](err)
	}
	return plug.NewResult[[]string](list)
}

func (h *Handler) Get(id string) plug.Result[plug.Row] {
	redis := NewRedis()
	pattern := fmt.Sprintf("%s:", id)
	keys, err := redis.Keys(pattern)
	if err != nil {
		return plug.NewErrResult[plug.Row](err)
	}
	row := plug.Row{
		Id: id,
		Values: make(plug.Values, 0),
	}
	for _, key := range keys {
		val, err := redis.Get(key)
		if err != nil {
			return plug.NewErrResult[plug.Row](err)
		}
		name := strings.TrimPrefix(key, fmt.Sprintf("%s:", id))
		row.Values[name] = val
	}

	return plug.NewResult[plug.Row](row)
}

func (h *Handler) Set(row plug.Row) error {
	if err := h.Del(row.Id); err != nil {
		return err
	}
	redis := NewRedis()
	for name, value := range row.Values {
		key := fmt.Sprintf("%s:%s", row.Id, name)
		if err := redis.Set(key, value); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) Del(id string) error {
	redis := NewRedis()
	pattern := fmt.Sprintf("%s:", id)
	keys, err := redis.Keys(pattern)
	if err != nil {
		return err
	}
	for _, key := range keys {
		if err := redis.Delete(key); err != nil {
			return err
		}
	}
	return nil
}
