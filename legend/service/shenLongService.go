package service

import (
	"legend/dao"
	"legend/model"
)

type ShenLongService struct {
}

func (sls *ShenLongService) ShenLongQueryService() ([]model.ShenLong,error) {
	//操作数据库
	shenLongDao := dao.NewShenLongDao()
	return shenLongDao.ShenLongQueryDao()
}