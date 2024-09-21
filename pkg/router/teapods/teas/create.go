package teas

// type CreateTeaReq struct {
// 	Teabox string            `json:"teabox" validate:"min=1"`
// 	Vals   map[string]string `json:"vals" validate:"required"`
// }

// func CreateTea(c echo.Context) error {
// 	teapodName := c.Param("teapod")

// 	var req CreateTeaReq
// 	if err := Validate(c, &req); err != nil {
// 		return err
// 	}

// 	teapodSrv := service.NewTeapodSrv(teapodName)
// 	teabox, err := teapodSrv.GetTeabox(req.Teabox)
// 	if err != nil {
// 		return err
// 	}

// 	if err := teapodSrv.ValidateTeaboxVals(teabox, req.Vals); err != nil {
// 		apperr := NewAppErr()
// 		apperr.Errors = append(apperr.Errors, AppErrItem{
// 			Path:    "",
// 			Message: err.Error(),
// 		})
// 		return apperr
// 	}

// 	teaid, err := teapodSrv.CreateTea(req.Teabox, req.Vals)
// 	if err != nil {
// 		return err
// 	}
// 	return WithData(c, NewIdSchema(teaid))
// }
