package service

import (
	"legend/dao"
	"legend/model"
)

type HumanoidService struct {
}

func (hus *HumanoidService) HumanoidQueryService() ([]model.Humanoid,error) {
	//操作数据库
	humanoidDao := dao.NewHumanoidDao()
	return humanoidDao.HumanoidQueryDao()
}