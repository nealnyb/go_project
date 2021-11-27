package dao

import (
	"legend/model"
	"legend/tool"
)

type ShenLongDao struct {
	*tool.Orm
}

func NewShenLongDao() *ShenLongDao {
	return &ShenLongDao{Orm:tool.DbEngine}
}

func (sld *ShenLongDao) ShenLongQueryDao() ([]model.ShenLong,error){
	var shenLongs []model.ShenLong
	if err := sld.Engine.Find(&shenLongs);err != nil {
		return nil,err
	}
	return shenLongs,nil
}