package dao

import (
	"legend/model"
	"legend/tool"
)

type ShengYuDao struct {
	*tool.Orm
}

func NewShengYuDao() *ShengYuDao {
	return &ShengYuDao{Orm:tool.DbEngine}
}

func (syd *ShengYuDao) ShengYuQueryDao() ([]model.ShengYu,error){
	var shengYus []model.ShengYu
	if err := syd.Engine.Find(&shengYus); err != nil {
		return nil,err
	}
	return shengYus,nil
}