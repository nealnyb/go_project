package service

import (
	"legend/dao"
	"legend/model"
)

type HongHuangService struct {
}

func (hhs *HongHuangService) HongHuangQueryService() ([]model.HongHuang,error) {
	//操作数据库
	hongHuangDao := dao.NewHongHuangDao()
	return hongHuangDao.HongHuangQueryDao()
}