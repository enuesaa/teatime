package controller

import (
	"fmt"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
)

func InvokePlugin(c *gin.Context) {
	pluginSrv := service.NewPluginService("./plugins/teatime-plugin-pinit/teatime-plugin-pinit")
	info, err := pluginSrv.GetInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info)
}
