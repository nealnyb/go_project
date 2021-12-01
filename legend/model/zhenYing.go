package model

type ZhenYing struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(30)" json:"name"`
	Addr string `xorm:"varchar(30)" json:"addr"`
	RefreshTime int64 `xorm:"bigint" json:"refresh_time"`
	Path string `xorm:"varchar(50)" json:"path"`
	Refresh string `xorm:"varchar(50)" json:"refresh"`
}