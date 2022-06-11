package controller

import (
	"context"
	"server_gf/api/v1/system"
	systemService "server_gf/internal/app/system/service"
)

type configController struct {
}

var SysConfig = configController{}

// List 系统参数列表--..
func (c *configController) List(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	res, err = systemService.SysConfig().ListSysConfigs(ctx, req)
	return
}
