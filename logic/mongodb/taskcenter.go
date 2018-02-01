package mongodb

import (
	"time"

	"ftgoo/logic/errors"

	"gopkg.in/mgo.v2/bson"
)

type taskCenter struct {
}

/* 获取开始时间在指定时间段中的非预置费归档项目组
startTime:开始时间 endTime:截止时间 */
func (taskCenter) GetFoldersCreateTimeBetween(startTime time.Time, endTime time.Time) ([]*Folder, error) {
	session, col := FolderColProvider()
	defer session.Close()
	query := bson.M{
		"Archived":   false,
		"Predefined": false,
		"CreateTime": bson.M{
			"$lt":  endTime,
			"$gte": startTime,
		},
	}
	pro := bson.M{
		"FolderID": 1,
	}
	var folders []*Folder
	//所有符合条件的项目
	err := col.Find(query).Select(pro).All(&folders)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

/* 获取指定项目id且创建时间小于指定时间的未删除任务组
folderId:项目id,tm:指定早于的时间(time.Time{}时不进入筛选) */
func (taskCenter) GetTasksByFolderIdAndTime(folderId string, date time.Time) ([]*Task, error) {
	session, col := TaskColProvider()
	defer session.Close()
	query := bson.M{
		"FolderID": folderId,
		"IsDelete": false,
	}
	if date.Equal(time.Time{}) {
		query["CreateTime"] = bson.M{
			"$lt": date,
		}
	}
	pro := bson.M{
		"Status":           1,
		"Deadline":         1,
		"stm":              1,
		"ChargeAccountID":  1,
		"Members":          1,
		"CreateTime":       1,
		"StatusModifyTime": 1,
	}
	var tasks []*Task
	err := col.Find(query).Select(pro).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

/* 获取时间段统计
folderId:项目id,startTime:开始时间，endTime:截止时间 */
func (taskCenter) GetFolderDailyStatisticsByDate(folderId string, startTime time.Time, endTime time.Time) ([]*FolderStatistics, error) {
	session, col := FolderColProvider()
	defer session.Close()
	query := bson.M{
		"fid": folderId,
		"dt": bson.M{
			"$gte": startTime,
			"$lt":  endTime,
		},
	}
	var fsts []*FolderStatistics
	err := col.Find(query).All(&fsts)
	if err != nil {
		return nil, err
	}
	return fsts, nil
}

/* 插入项目统计
fsts:项目统计实体 */
func (taskCenter) InsertFolderStatistics(fsts *FolderStatistics) (bool, error) {
	if fsts == nil {
		return false, errors.ERR_PARAMETER
	}
	session, col := FolderStatColProvider()
	defer session.Close()
	err := col.Insert(fsts)
	if err != nil {
		return false, err
	}
	return true, nil
}
