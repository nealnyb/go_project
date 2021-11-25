package tool

import (
	"legend/model"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var DbEngine *Orm
type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm,error){
	db := cfg.Database
	conn := db.User + ":" + db.Password + "@tcp(" + db.Host + ":" +db.Port + ")/" + db.DbName + "?charset=" + db.Charset
	engine,err := xorm.NewEngine(db.Driver,conn)
	if err != nil {
		return nil,err
	}

	engine.ShowSQL(db.ShowSql)

	err = engine.Sync2(
	new(model.Zodiac),
	new(model.Home),
	new(model.HongHuang),
)
	if err != nil {
		return nil,err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm,nil
} 