package controller

import (
	"golang.org/x/exp/maps"
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
)

type ListProviderResponseItem struct {
	Name string `json:"name"`
}
type ListProviderResponse struct {
	Items []ListProviderResponseItem `json:"items"`
}

func ListProviders(c *gin.Context) {
	res := ListProviderResponse{
		Items: make([]ListProviderResponseItem, 0),
	}
	for _, name := range maps.Keys(service.ProviderCommandMap) {
		res.Items = append(res.Items, ListProviderResponseItem{
			Name: name,
		})
	}
	c.JSON(200, res)
}

type DescribeProviderResponse struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cards       []string `json:"cards"`
	PanelMap    map[string]string
}
func DescribeProvider(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)
	info, err := providerSrv.GetInfo()
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := DescribeProviderResponse{
		Name:        info.Name,
		Description: info.Description,
		Cards:       info.Cards,
		PanelMap:    info.PanelMap,
	}
	c.JSON(200, res)
}

type DescribeCardResponse struct {
	Enable   bool                `json:"enable"`
	Layout   string              `json:"layout"`
	TextCard plug.TextCardConfig `json:"textCardConfig"`
}
func DescribeCard(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	cardName := c.Param("cardName")
	card, err := providerSrv.DescribeCard(cardName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := DescribeCardResponse{
		Enable:   card.Enable,
		Layout:   card.Layout,
		TextCard: card.TextCard,
	}
	c.JSON(200, res)
}

type DescribePanelResponse struct {
	Enable     bool                  `json:"enable"`
	Layout     string                `json:"layout"`
	TablePanel plug.TablePanelConfig `json:"tablePanelConfig"`
}
func DescribePanel(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	panelName := c.Param("panelName")
	panel, err := providerSrv.DescribePanel(panelName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := DescribePanelResponse{
		Enable:     panel.Enable,
		Layout:     panel.Layout,
		TablePanel: panel.TablePanel,
	}
	c.JSON(200, res)
}

// func (srv *ProviderService) Register(model string, name string) error {
// 	provider, err := srv.GetProvider()
// 	if err != nil {
// 		return err
// 	}
// 	return provider.Register(plug.RegisterArg{Model: model, Name: name})
// }

// func (srv *ProviderService) Get(model string, name string) (plug.Record, error) {
// 	provider, err := srv.GetProvider()
// 	if err != nil {
// 		return plug.Record{}, err
// 	}
// 	return provider.Get(plug.GetArg{Model: model, Name: name}), nil
// }

// func (srv *ProviderService) Set(model string, name string, record plug.Record) error {
// 	provider, err := srv.GetProvider()
// 	if err != nil {
// 		return err
// 	}
// 	return provider.Set(plug.SetArg{Model: model, Name: name, Record: record})
// }

// func (srv *ProviderService) Del(model string, name string) error {
// 	provider, err := srv.GetProvider()
// 	if err != nil {
// 		return err
// 	}
// 	return provider.Del(plug.DelArg{Model: model, Name: name})
// }

