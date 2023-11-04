package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

// should implement ProviderInterface
type Handler struct {}
func (s *Handler) Info() plug.Info {
	return plug.Info{
		Name: "aaa",
		Description: "pinit provider",
		Cards: []string{
			"main",
		},
		PanelMap: map[string]string{
			"main": "main",
		},
	}
}

func (h *Handler) DescribeCard(arg plug.DescribeCardArg) plug.Card {
	if arg.Name == "main" {
		return plug.Card{
			Enable: true,
			Layout: "textCard",
			TextCard: plug.TextCardConfig{
				Heading: "pinit status",
				Center: true,
			},
		}
	}
	return plug.Card{Enable: false}
}
func (h *Handler) DescribePanel(arg plug.DescribePanelArg) plug.Panel {
	return plug.Panel{
		Enable: true,
		Layout: "main",
		TablePanel: plug.TablePanelConfig{
			Title: "title",
			Description: "sample table panel",
			Head: []string{"a", "b", "c"},
			Records: []plug.TablePanelRecord{
				{
					Model: "memos",
					Name: "first-memo",
					Values: []plug.TablePanelRecordValue {
						{
							Readonly: false,
							Key: "name",
						},
					},
				},
			},
		},
	}
}
func (h *Handler) Register(arg plug.RegisterArg) error {
	return nil
}
func (h *Handler) Get(arg plug.GetArg) plug.Record {
	return plug.Record{}
}
func (h *Handler) Set(arg plug.SetArg) error {
	return nil
}
func (h *Handler) Del(arg plug.DelArg) error {
	return nil
}
