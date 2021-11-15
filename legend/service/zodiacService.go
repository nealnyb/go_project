package service

import (
	"legend/dao"
	"legend/model"
	"legend/tool"
)

type ZodiacService struct {
}

func (zs *ZodiacService) ZodiacAdd() bool {
	//将生肖信息保存到数据库
	zodiacInfo := model.Zodiac{
		Name: "牛",
		Addr: "生肖殿堂",
		RefreshTime: 7200,
		Path: "下5左2右1",	
	}
	zodiacDao := dao.ZodiacDao{Orm:tool.DbEngine}
	res := zodiacDao.InsertZodiac(zodiacInfo)
	return res > 0
}