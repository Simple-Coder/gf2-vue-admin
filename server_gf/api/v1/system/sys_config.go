package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "server_gf/api/v1/common"
	systemEntity "server_gf/internal/app/system/model/entity"
)

type ConfigSearchReq struct {
	g.Meta     `path:"/config/list" tags:"系统参数管理" method:"get" summary:"系统参数列表"`
	ConfigName string `p:"configName"` //参数名称
	ConfigKey  string `p:"configKey"`  //参数键名
	ConfigType string `p:"configType"` //状态
	commonApi.PageReq
}

type ConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*systemEntity.SysConfig `json:"list"`
	commonApi.ListRes
}
