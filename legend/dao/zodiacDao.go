package dao

import (
	"legend/model"
	"legend/tool"
	"log"
)

type ZodiacDao struct {
	 *tool.Orm
}

//实例化dao对象
func NewZoaZodiacDao() *ZodiacDao {
	return &ZodiacDao{Orm:tool.DbEngine}
}

//从数据库查询数据，并返回
func (zd *ZodiacDao) QueryZodiac() ([]model.Zodiac,error) {
	var zodiacs []model.Zodiac
	if err := zd.Engine.Find(&zodiacs); err != nil {
		return nil,err
	}
	return zodiacs,nil
}



func (zd *ZodiacDao) InsertZodiac(zodiac model.Zodiac) int64 {
	res,err := zd.InsertOne(&zodiac)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res

}