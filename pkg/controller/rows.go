package controller

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Note struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
func ListRows(c echo.Context) error {
	res := NewListResponse[IdSchema]()

	providerSrv := service.NewProviderService("coredata")
	ids, err := providerSrv.ListRows()
	if err != nil {
		return err
	}
	for _, id := range ids {
		res.Items = append(res.Items, IdSchema{
			Id: id,
		})
	}

	return c.JSON(200, res)
}

func DescribeRow(c echo.Context) error {
    id := c.Param("id")

	providerSrv := service.NewProviderService("coredata")
	row, err := providerSrv.GetRow(id)
	if err != nil {
		return err
	}
	note := Note {
		Name: row.Values["name"],
		Description: row.Values["description"],
	}
	res := NewDescribeResponse[Note](note)
	return c.JSON(200, res)
}

func CreateRow(c echo.Context) error {
	var note Note
	if err := Validate(c, &note); err != nil {
		return err
	}
	values := plug.Values{
		"name": note.Name,
		"description": note.Description,
	}
	providerSrv := service.NewProviderService("coredata")
	id, err := providerSrv.CreateRow(values)
	if err != nil {
		return err
	}
	return c.JSON(200, NewIdSchema(id))
}

func UpdateRow(c echo.Context) error {
    id := c.Param("id")

	var note Note
	if err := Validate(c, &note); err != nil {
		return err
	}
	values := plug.Values{
		"name": note.Name,
		"description": note.Description,
	}
	providerSrv := service.NewProviderService("coredata")
	_, err := providerSrv.UpdateRow(id, values)
	if err != nil {
		return err
	}
	return c.JSON(200, NewIdSchema(id))
}

func DeleteRow(c echo.Context) error {
	id := c.Param("id")

	providerSrv := service.NewProviderService("coredata")
	if err := providerSrv.DeleteRow(id); err != nil {
		return err
	}

	return c.JSON(200, NewEmptySchema())
}
