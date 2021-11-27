package service

import (
	"legend/dao"
	"legend/model"
)

type HomeService struct {
}

func (hs *HomeService) HomeQueryService() ([]model.Home, error) {
	//操作数据库
	homeDao := dao.NewHomeDao()
	return homeDao.HomeQueryDao()
}

func (hs *HomeService) HomeRefeshUpdateService(data *model.Home) (int64,error){
	//操作数据库
	homeDao := dao.NewHomeDao()
	return homeDao.HomeRefeshUpdateDao(data)
}