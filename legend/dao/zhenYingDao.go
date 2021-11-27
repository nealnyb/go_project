package dao

import (
	"legend/model"
	"legend/tool"
)

type ZhenYingDao struct {
	*tool.Orm
}

func NewZhenYingDao() *ZhenYingDao {
	return &ZhenYingDao{Orm:tool.DbEngine}
}

func (zyd *ZhenYingDao) ZhenYingQueryDao() ([]model.ZhenYing,error) {
	var zhenYings []model.ZhenYing
	if err := zyd.Engine.Find(&zhenYings);err != nil {
		return nil,err
	}
	return zhenYings,nil
}

func (zyd *ZhenYingDao) ZhenYingRefreshUpdateDao(data *model.ZhenYing) (int64,error) {
	return zyd.Engine.Id(data.Id).Update(data)
}