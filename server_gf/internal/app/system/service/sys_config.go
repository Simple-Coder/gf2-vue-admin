package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"server_gf/api/v1/system"
	systemConsts "server_gf/internal/app/system/consts"
	"server_gf/internal/app/system/model/entity"
	"server_gf/internal/app/system/service/internal/dao"
	"server_gf/internal/app/system/service/internal/do"
	"server_gf/utility"
)

type ISysConfigService interface {
	GetList(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error)
	Add(ctx context.Context, req *system.ConfigAddReq, userId uint64) (err error)
	Edit(ctx context.Context, req *system.ConfigEditReq, userId uint64) (err error)
	Get(ctx context.Context, id int) (res *system.ConfigGetRes, err error)
	Delete(ctx context.Context, ids []int) (err error)
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
				m = m.WhereLike(dao.SysConfig.Columns().ConfigName, "%"+req.ConfigName+"%")
			}
			if req.ConfigType != "" {
				m = m.Where(dao.SysConfig.Columns().ConfigType, gconv.Int(req.ConfigType))
			}
			if req.ConfigKey != "" {
				m = m.WhereLike(dao.SysConfig.Columns().ConfigKey, "%"+req.ConfigKey+"%")
			}
			if len(req.DateRange) > 0 {
				m = m.WhereGTE(dao.SysConfig.Columns().CreatedAt, req.DateRange[0])
				m = m.WhereLTE(dao.SysConfig.Columns().CreatedAt, req.DateRange[1])
				//m = m.Where("created_at >= ? AND created_at<=?", req.DateRange[0], req.DateRange[1])
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

func (s *configImpl) Add(ctx context.Context, req *system.ConfigAddReq, userId uint64) (err error) {
	err = g.Try(func() {
		err = s.CheckConfigKeyUnique(ctx, req.ConfigKey)
		utility.ErrIsNil(ctx, err)
		_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			CreateBy:    userId,
			Remark:      req.Remark,
		})
		utility.ErrIsNil(ctx, err, "添加系统参数失败")
		//清除缓存
		//Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// Get 获取系统参数
func (s *configImpl) Get(ctx context.Context, id int) (res *system.ConfigGetRes, err error) {
	res = new(system.ConfigGetRes)
	err = g.Try(func() {
		err = dao.SysConfig.Ctx(ctx).WherePri(id).Scan(&res.Data)
		utility.ErrIsNil(ctx, err, "获取系统参数失败")
	})
	return
}

// Edit 修改系统参数
func (s *configImpl) Edit(ctx context.Context, req *system.ConfigEditReq, userId uint64) (err error) {
	err = g.Try(func() {
		err = s.CheckConfigKeyUnique(ctx, req.ConfigKey, req.ConfigId)
		utility.ErrIsNil(ctx, err)
		_, err = dao.SysConfig.Ctx(ctx).WherePri(req.ConfigId).Update(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			UpdateBy:    userId,
			Remark:      req.Remark,
		})
		utility.ErrIsNil(ctx, err, "修改系统参数失败")
		//清除缓存
		//Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// CheckConfigKeyUnique 验证参数键名是否存在
func (s *configImpl) CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int64) (err error) {
	err = g.Try(func() {
		data := (*entity.SysConfig)(nil)
		m := dao.SysConfig.Ctx(ctx).Fields(dao.SysConfig.Columns().ConfigId).Where(dao.SysConfig.Columns().ConfigKey, configKey)
		if len(configId) > 0 {
			m = m.Where(dao.SysConfig.Columns().ConfigId+" != ?", configId[0])
		}
		err = m.Scan(&data)
		utility.ErrIsNil(ctx, err, "校验失败")
		if data != nil {
			utility.ErrIsNil(ctx, errors.New("参数键名重复"))
		}
	})
	return
}

// Delete 删除系统参数
func (s *configImpl) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(func() {
		_, err = dao.SysConfig.Ctx(ctx).Delete(dao.SysConfig.Columns().ConfigId+" in (?)", ids)
		utility.ErrIsNil(ctx, err, "删除失败")
		//清除缓存
		//Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}
