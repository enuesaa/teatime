package main

import (
	"fmt"
	"strings"

	"github.com/enuesaa/teatime/pkg/plug"
)

func main() {
	handler := Handler{
		redis: NewRedis(),
	}
	plug.Serve(&handler)
}

type Handler struct {
	redis Redis
}

func (h *Handler) Init() error {
	return h.redis.Run()
}

func (s *Handler) Info() plug.InfoResult {
	info := plug.Info{
		Name: "coredata",
		Description: "coredata provider",
	}
	return plug.NewInfoResult(info)
}

func (h *Handler) List() plug.ListResult {
	list, err := h.redis.Keys("*")
	if err != nil {
		return plug.NewListErrResult(err)
	}
	return plug.NewListResult(list)
}

func (h *Handler) Get(id string) plug.GetResult {
	keys, err := h.redis.Keys(h.fmtPattern(id))
	if err != nil {
		return plug.NewGetErrResult(err)
	}
	row := plug.Row{
		Id: id,
		Values: make(plug.Values, 0),
	}
	for _, key := range keys {
		val, err := h.redis.Get(key)
		if err != nil {
			return plug.NewGetErrResult(err)
		}
		_, name := h.splitKey(key)
		row.Values[name] = val
	}

	return plug.NewGetResult(row)
}

func (h *Handler) Set(row plug.Row) error {
	if err := h.Del(row.Id); err != nil {
		return err
	}
	for name, value := range row.Values {
		if err := h.redis.Set(h.fmtKey(row.Id, name), value); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) Del(id string) error {
	keys, err := h.redis.Keys(h.fmtPattern(id))
	if err != nil {
		return err
	}
	for _, key := range keys {
		if err := h.redis.Delete(key); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) fmtKey(id string, name string) string {
	return fmt.Sprintf("%s:%s", id, name)
}

func (h *Handler) fmtPattern(id string) string {
	return fmt.Sprintf("%s:*", id)
}

func (h *Handler) splitKey(key string) (string, string) {
	splitted := strings.Split(key, ":")
	if len(splitted) == 2 {
		return splitted[0], splitted[1]
	}
	return "", ""
}
