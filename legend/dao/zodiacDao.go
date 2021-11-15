package dao

import (
	"legend/model"
	"legend/tool"
	"log"
)

type ZodiacDao struct {
	 *tool.Orm
}

func (zd *ZodiacDao) InsertZodiac(zodiac model.Zodiac) int64 {
	res,err := zd.InsertOne(&zodiac)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res

}