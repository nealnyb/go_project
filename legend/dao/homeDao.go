package dao

import (
	"legend/model"
	"legend/tool"
)

type HomeDao struct {
	*tool.Orm
}

func NewHomeDao() *HomeDao {
	return &HomeDao{Orm:tool.DbEngine}
}

func (hd *HomeDao) HomeQueryDao() ([]model.Home,error) {
	var homes []model.Home
	if err := hd.Engine.Find(&homes); err != nil {
		return nil,err
	}
	return homes,nil
}