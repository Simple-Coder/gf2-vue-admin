package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"server_gf/api/v1/system"
	systemConsts "server_gf/internal/app/system/consts"
	"server_gf/internal/app/system/service/internal/dao"
	"server_gf/utility"
)

type ISysConfigService interface {
	GetList(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error)
}
type configImpl struct {
}

var sysConfigService = configImpl{}

func SysConfig() ISysConfigService {
	return &sysConfigService
}

func (s *configImpl) GetList(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	res = new(system.ConfigSearchRes)
	err = g.Try(func() {
		m := dao.SysConfig.Ctx(ctx)
		if req != nil {
			if req.ConfigName != "" {
				m = m.WhereLike(dao.SysConfig.Columns().ConfigName, req.ConfigName)
			}
			if req.ConfigType != "" {
				m = m.Where(dao.SysConfig.Columns().ConfigType, gconv.Int(req.ConfigType))
			}
			if req.ConfigKey != "" {
				m = m.Where(dao.SysConfig.Columns().ConfigKey, gconv.Int(req.ConfigKey))
			}
			if len(req.DateRange) > 0 {
				m = m.WhereGTE(dao.SysConfig.Columns().CreatedAt, req.DateRange[0])
				m = m.WhereLTE(dao.SysConfig.Columns().CreatedAt, req.DateRange[1])
			}
		}
		res.Total, err = m.Count()
		utility.ErrIsNil(ctx, err, "获取数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Page(req.PageNum, req.PageSize).OrderAsc(dao.SysConfig.Columns().ConfigId).Scan(&res.List)
		utility.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}
