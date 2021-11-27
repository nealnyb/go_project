package service

import (
	"legend/dao"
	"legend/model"
)

type ShengYuService struct {
}

func (sys *ShengYuService) ShengYuQueryService() ([]model.ShengYu,error){
	//操作数据库
	shengYuDao := dao.NewShengYuDao()
	return shengYuDao.ShengYuQueryDao()
}