package dao

import (
	"legend/model"
	"legend/tool"
)

type HumanoidDao struct {
	*tool.Orm
}

func NewHumanoidDao() *HumanoidDao {
	return &HumanoidDao{Orm:tool.DbEngine}
}

func (hud *HumanoidDao) HumanoidQueryDao() ([]model.Humanoid,error) {
	var humanoids []model.Humanoid
	if err := hud.Engine.Find(&humanoids);err != nil {
		return nil,err
	}
	return humanoids,nil
}

func (hud *HumanoidDao) HumanoidRefreshUpdateDao(data *model.Humanoid) (int64,error){
	return hud.Engine.Id(data.Id).Update(data)
}