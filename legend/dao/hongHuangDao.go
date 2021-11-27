package dao

import (
	"legend/model"
	"legend/tool"
)

type HongHuangDao struct {
	*tool.Orm
}

func NewHongHuangDao() *HongHuangDao { 
	return &HongHuangDao{Orm:tool.DbEngine}
 }

 func (hhd *HongHuangDao) HongHuangQueryDao() ([]model.HongHuang,error) {
	var honghuangs []model.HongHuang
	if err := hhd.Engine.Find(&honghuangs);err != nil {
		return nil,err
	}
	return honghuangs,nil
 }

 func (hhd *HongHuangDao) HongHuangRefreshUpdateDao(data *model.HongHuang) (int64,error) {
	 return hhd.Engine.Id(data.Id).Update(data)
 }