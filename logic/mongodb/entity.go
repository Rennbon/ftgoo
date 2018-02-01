package mongodb

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

var maxDate time.Time = time.Date(8000, 1, 1, 0, 0, 0, 0, time.Local)

type FolderStatistics struct {
	Id                bson.ObjectId `bson:"_id,omitempty"`
	FolderId          string        `bson:"fid"` //项目id
	CreateTime        time.Time     `bson:"ct"`  //创建时间
	Date              time.Time     `bson:"dt"`  //统计的当天日期
	Amount            int           `bson:"at"`  //任务总数
	Underway          int           `bson:"u"`   //进行中任务总数
	Underway_N        int           `bson:"un"`  //正常状态进行中任务数
	Underway_A        int           `bson:"ua"`  //逾期进行中任务数
	Underway_U        int           `bson:"uu"`  //未设定截止日期的进行中任务数
	Timespan          int64         `bson:"t"`   //总逾期时间 单位：小时
	Timespan_Und      int64         `bson:"tu"`  //进行中任务总逾期时间 单位：小时
	Timespan_Com      int64         `bson:"tc"`  //已完成任务总逾期时间 单位：小时
	Completed         int           `bson:"c"`   //已完成任务总数
	Completed_N       int           `bson:"cn"`  //正常已完成任务数
	Completed_A       int           `bson:"ca"`  //逾期已完成任务总数
	Completed_U       int           `bson:"cu"`  //未设定截止日期已完成任务数
	MemberAmount      int           `bson:"ma"`  //参与成员数
	ChargeAmount      int           `bson:"cha"` //任务负责人数
	NewTasks          int           `bson:"nts"` //当天新增任务数
	CompletedTasks    int           `bson:"cts"` //当天完成任务数
	Timespan_BurnDown int64         `bson:"tbd"` //进行中燃尽时间
}

type Folder struct {
	FolderId   string `bson:"FolderID"`   //项目id
	Archived   bool   `bson:"Archived"`   //归档
	Predefined bool   `bson:"Predefined"` //预置
}
type TaskStatus int

const (
	Cs_Incomplete TaskStatus = 0 //未完成
	Cs_Complete   TaskStatus = 1 //完成
)

type Task struct {
	TaskId           string       `bson:"TaskID"`           //任务ID
	FolderId         string       `bson:"FolderID"`         //所属项目ID
	Deadline         time.Time    `bson:"Deadline"`         //截止时间
	StartTime        time.Time    `bson:"stm"`              //开始时间
	ChargeAccountId  string       `bson:"ChargeAccountID"`  //负责人
	Status           TaskStatus   `bson:"Status"`           //状态
	StatusModifyTime time.Time    `bson:"StatusModifyTime"` //修改时间，配合status为1即完成时间
	UpdateTime       time.Time    `bson:"UpdateTime"`       //更新时间
	CreateTime       time.Time    `bson:"CreateTime"`       //创建时间
	IsDelete         bool         `bson:"IsDelete"`         //是否删除
	Members          []TaskMember `bson:"Members"`          //成员
}

type MemberType int

const (
	Cs_Member   MemberType = 0 //成员
	Cs_Releaser MemberType = 1 //托付人
)

type MemberStatus int

const (
	Cs_Normal MemberStatus = 0 //正常
)

type TaskMember struct {
	aId    string       `bson:"AccountID"`   //成员id
	Type   MemberType   `bson:"Type"`        //成员类型  0:成员 1:托付者 （其他无用，这里只读）
	Status MemberStatus `bson:"ApplyStatus"` //状态 0:正常 （其他无用，这里只读）
}
