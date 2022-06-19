package controller

import (
	"context"
	"server_gf/api/v1/system"
	systemService "server_gf/internal/app/system/service"
)

type configController struct {
}

var SysConfig = configController{}

// List 系统参数列表--....
func (c *configController) List(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	res, err = systemService.SysConfig().GetList(ctx, req)
	return
}

// Get 获取系统参数
func (c *configController) Get(ctx context.Context, req *system.ConfigGetReq) (res *system.ConfigGetRes, err error) {
	res, err = systemService.SysConfig().Get(ctx, req.Id)
	return
}

// Add 添加系统参数
func (c *configController) Add(ctx context.Context, req *system.ConfigAddReq) (res *system.ConfigAddRes, err error) {
	//err = systemService.SysConfig().Add(ctx, req, service.Context().GetUserId(ctx))
	err = systemService.SysConfig().Add(ctx, req, 123)
	return
}

// Edit 修改系统参数
func (c *configController) Edit(ctx context.Context, req *system.ConfigEditReq) (res *system.ConfigEditRes, err error) {
	//err = systemService.SysConfig().Edit(ctx, req, service.Context().GetUserId(ctx))
	err = systemService.SysConfig().Edit(ctx, req, 123)
	return
}

// Delete 删除系统参数
func (c *configController) Delete(ctx context.Context, req *system.ConfigDeleteReq) (res *system.ConfigDeleteRes, err error) {
	err = systemService.SysConfig().Delete(ctx, req.Ids)
	return
}
