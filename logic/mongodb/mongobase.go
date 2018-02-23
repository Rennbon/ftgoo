package mongodb

import (
	"time"

	cnf "github.com/Rennbon/ftgoo/config"

	"gopkg.in/mgo.v2"
)

const (
	db_taskcenter  string = "taskcenter"
	col_task       string = "task"
	col_folder     string = "folder"
	col_folderStat string = "folderStatisticsNew"
)

var (
	taskCenterbase *mgo.Session
)

func init() {
	conf, err := cnf.LoadConfig()
	if err != nil {
		panic(err)
	}
	var keys []string
	keys = append(keys, "TaskCenter")
	err = cnf.CheckConfig(conf, keys)
	if err != nil {
		panic(err)
	}
	loadTaskSession(conf)
}
func loadTaskSession(c *cnf.Config) error {

	mongoDBInfo := &mgo.DialInfo{
		Addrs:     []string{c.TaskCenter.Addr},
		Timeout:   c.TaskCenter.Timeout * time.Second,
		PoolLimit: c.TaskCenter.PoolLimit,
		Database:  c.TaskCenter.Database,
	}
	session, err := mgo.DialWithInfo(mongoDBInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	taskCenterbase = session
	return nil
}
func CloneTaskCenter() *mgo.Session {
	return taskCenterbase.Clone()
}
func TaskColProvider() (*mgo.Session, *mgo.Collection) {
	session := taskCenterbase.Clone()
	col := session.DB(db_taskcenter).C(col_task)
	return session, col
}
func FolderColProvider() (*mgo.Session, *mgo.Collection) {
	session := taskCenterbase.Clone()
	col := session.DB(db_taskcenter).C(col_folder)
	return session, col
}
func FolderStatColProvider() (*mgo.Session, *mgo.Collection) {
	session := taskCenterbase.Clone()
	col := session.DB(db_taskcenter).C(col_folderStat)
	return session, col
}
