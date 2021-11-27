package service

import (
	"legend/dao"
	"legend/model"
)

type ZhenYingService struct {
}

func (zys *ZhenYingService) ZhenYingQueryService() ([]model.ZhenYing,error){
	//操作数据库
	zhenYingDao := dao.NewZhenYingDao()
	return zhenYingDao.ZhenYingQueryDao()
}

func (zys *ZhenYingService) ZhenYingRefreshUpdateService(data *model.ZhenYing) (int64,error) {
	//操作数据库
	zhenYingDao := dao.NewZhenYingDao()
	return zhenYingDao.ZhenYingRefreshUpdateDao(data)
}
