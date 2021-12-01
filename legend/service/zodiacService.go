package service

import (
	"legend/dao"
	"legend/model"
	"legend/tool"
)

type ZodiacService struct {
}



func (zs *ZodiacService) ZodiacQuery() ([]model.Zodiac,error) {
	//操作数据库
	zodiacDao := dao.NewZoaZodiacDao()
	return zodiacDao.QueryZodiac()
}

func (zs *ZodiacService) ZodiacAdd() bool {
	//将生肖信息保存到数据库
	zodiacInfo := model.Zodiac{
		Refresh: "",	
	}
	zodiacDao := dao.ZodiacDao{Orm:tool.DbEngine}
	res := zodiacDao.InsertZodiac(zodiacInfo)
	return res > 0
}

func (zs *ZodiacService) ZodiacRefreshUpdateService(data *model.Zodiac) (int64,error){
	//操作数据库
	zodiacDao := dao.NewZoaZodiacDao()
	return zodiacDao.ZodiacRefreshUpdateDao(data)
}