package internal

import (
	"Server/src/server/base"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}



func (m *Module) OnInit() {
	m.Skeleton = skeleton
	InitLoginTables()

}

func InitLoginTables() {
	//db := mysql.MysqlDB()
	//db.AutoMigrate(&Account{})
}


func (m *Module) OnDestroy() {

}
