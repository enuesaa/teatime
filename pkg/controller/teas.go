package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func ListTeas(c echo.Context) error {
	teapodName := c.Param("teapod")

	teas, err := service.NewTeapodSrv(teapodName).ListTeas()
	if err != nil {
		return err
	}
	list := make([]IdSchema, 0)
	for _, tea := range teas {
		list = append(list, IdSchema{
			Id: tea.Teaid,
		})
	}
	return WithData(c, list)
}

func GetTea(c echo.Context) error {
	teapodName := c.Param("teapod")
	teaid := c.Param("teaid")

	tea, err := service.NewTeapodSrv(teapodName).GetTea(teaid)
	if err != nil {
		return err
	}
	return WithData(c, tea.Vals)
}

type CreateTeaReq struct {
	Teabox string `json:"teabox" validate:"min=1"`
	Vals map[string]string `json:"vals" validate:"required"`
}
func CreateTea(c echo.Context) error {
	teapodName := c.Param("teapod")

	var req CreateTeaReq
	if err := Validate(c, &req); err != nil {
		return err
	}

	teapodSrv := service.NewTeapodSrv(teapodName)
	teabox, err := teapodSrv.GetTeabox(req.Teabox)
	if err != nil {
		return err
	}

	if err := teapodSrv.ValidateTeaboxVals(teabox, req.Vals); err != nil {
		apperr := NewAppErr()
		apperr.Errors = append(apperr.Errors, AppErrItem{
			Path: "",
			Message: err.Error(),
		})
		return apperr
	}

	teaid, err := teapodSrv.CreateTea(req.Vals)
	if err != nil {
		return err
	}
	return WithData(c, NewIdSchema(teaid))
}

func UpdateTea(c echo.Context) error {
	teapodName := c.Param("teapod")
	teaid := c.Param("teaid")


	var req CreateTeaReq
	if err := Validate(c, &req); err != nil {
		return err
	}

	teapodSrv := service.NewTeapodSrv(teapodName)
	teabox, err := teapodSrv.GetTeabox(req.Teabox)
	if err != nil {
		return err
	}

	if err := teapodSrv.ValidateTeaboxVals(teabox, req.Vals); err != nil {
		apperr := NewAppErr()
		apperr.Errors = append(apperr.Errors, AppErrItem{
			Path: "",
			Message: err.Error(),
		})
		return apperr
	}
	if _, err := teapodSrv.UpdateTea(teaid, req.Vals); err != nil {
		return err
	}
	return WithData(c, NewIdSchema(teaid))
}

func DeleteTea(c echo.Context) error {
	teapodName := c.Param("teapod")
	teaid := c.Param("teaid")

	if err := service.NewTeapodSrv(teapodName).DeleteTea(teaid); err != nil {
		return err
	}

	return WithData(c, NewEmptySchema())
}
