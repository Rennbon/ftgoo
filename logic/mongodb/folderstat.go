package mongodb

import (
	"ftgoo/tool"
	"log"
	"time"

	pb "ftgoo/logic/folderstat"

	"ftgoo/logic/errors"

	. "github.com/ahmetb/go-linq"
)

type FolderStatService struct {
}

var tcdb taskCenter

//项目统计日刷
func (FolderStatService) DailyFlushing() error {
	dateRight := tool.GetDate(time.Now().Local())
	dateLeft := dateRight.AddDate(0, -1, 0)
	iceAge := time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local)
	i := 0
	for {
		log.Println(dateRight, dateLeft)
		if dateLeft.Before(iceAge) || i > 10000 {
			break
		}
		i++
		err := dailyFlushingBetween(dateLeft, dateRight)
		if err != nil {
			log.Println(err)
		}
		dateRight = dateLeft
		dateLeft = dateRight.AddDate(0, -1, 0)
	}
	return nil
}
func dailyFlushingBetween(start time.Time, end time.Time) error {
	/* 	session := CloneTaskCenter()
	   	defer session.Close()
	   	fcol := session.DB(db_taskcenter).C(col_folder)
	   	fqy := bson.M{
	   		"Archived":   false,
	   		"Predefined": false,
	   		"CreateTime": bson.M{
	   			"$lt":  end,
	   			"$gte": start,
	   		},
	   	}
	   	projection := bson.M{
	   		"FolderID": 1,
	   	} */
	//var folders []Folder
	//所有符合条件的项目
	//err := fcol.Find(fqy).Select(projection).All(&folders)
	folders, err := tcdb.GetFoldersCreateTimeBetween(start, end)
	if err != nil {
		log.Print(err)
		return err
	}
	/* tcol := session.DB(db_taskcenter).C(col_task)
	tpro := bson.M{
		"Status":           1,
		"Deadline":         1,
		"stm":              1,
		"ChargeAccountID":  1,
		"Members":          1,
		"CreateTime":       1,
		"StatusModifyTime": 1,
	} */

	dateNow := tool.GetDate(time.Now().Local())             //当前时间
	fstsDate := dateNow.AddDate(0, 0, -1)                   //当前时间前一天，即需要统计的时间
	ddCompare := dateNow.Add(time.Hour * time.Duration(-1)) //对比时间

	//遍历项目
	for _, folder := range folders {
		/* var tasks []Task
		tqy := bson.M{
			"FolderID": folder.FolderId,
			"IsDelete": false,
			"CreateTime": bson.M{
				"$lt": dateNow,
			},
		}
		err = tcol.Find(tqy).Select(tpro).All(&tasks) */
		tasks, err := tcdb.GetTasksByFolderIdAndTime(folder.FolderId, dateNow)
		if err != nil {
			log.Print(err)
			return err
		}
		fsts, err := aggregateFolderStats(tasks, folder.FolderId, ddCompare, fstsDate)
		if err != nil {
			return err
		}
		if fsts == nil {
			break
		}
		fsts.CreateTime = time.Now().Local()
		/*fsts := &FolderStatistics{}
		fsts.FolderId = folder.FolderId
		fsts.Date = fstsDate
		fsts.CreateTime = time.Now().Local()
		fsts.Amount = len(tasks)
		var (
			chargeIds []string //存负责人总数
			memberIds []string //存成员总数
		)

		for _, task := range tasks {
			chargeIds = append(chargeIds, task.ChargeAccountId)
			for _, member := range task.Members {
				if member.Status == Cs_Normal && (member.Type == Cs_Member || member.Type == Cs_Releaser) {
					memberIds = append(memberIds, member.aId)
				}
			}
			if tool.GetDate(task.CreateTime) == fstsDate {
				fsts.NewTasks++
			}

			if task.Status == Cs_Incomplete { //进行中
				fsts.Underway++ //进行中任务书
				if task.Deadline.After(ddCompare) {
					fsts.Underway_N++ //正常状态进行中任务数
					if task.Deadline.After(maxDate) {
						fsts.Underway_U++ //未设定截止日期的进行中任务数
					} else {

					}
				} else {
					fsts.Underway_A++                                                //逾期进行中任务数
					fsts.Timespan_Und += int64(ddCompare.Sub(task.Deadline).Hours()) //进行中任务总逾期时间 单位：小时

				}

				if task.Deadline.Before(maxDate) {
					fsts.Timespan_BurnDown += int64(task.Deadline.Sub(ddCompare).Hours()) //时间燃尽图
				}

			} else { //已完成
				fsts.Completed++ //已完成任务总数
				if task.Deadline.Before(task.StatusModifyTime) {
					fsts.Timespan_Com += int64(task.StatusModifyTime.Sub(task.Deadline).Hours()) //已完成任务总逾期时间 单位：小时
					fsts.Completed_A++                                                           //逾期已完成任务总数
				} else {
					fsts.Completed_N++ //正常已完成任务数
					if task.Deadline == maxDate {
						fsts.Completed_U++ //未设定截止日期已完成任务数
					}
				}

				if tool.GetDate(task.StatusModifyTime) == fstsDate {
					fsts.CompletedTasks++
				}
			}
		}
		if fsts.Underway_A > 0 && fsts.Timespan_Und == 0 {
			fsts.Timespan_Und = 1
		}
		if fsts.Completed_A > 0 && fsts.Timespan_Com == 0 {
			fsts.Timespan_Com = 1
		}
		fsts.Timespan = fsts.Timespan_Und + fsts.Timespan_Com
		From(chargeIds).Distinct().ToSlice(&chargeIds)
		fsts.ChargeAmount = len(chargeIds)
		From(memberIds).Distinct().ToSlice(&memberIds)
		fsts.MemberAmount = len(memberIds) */

		/* fscol := session.DB(db_taskcenter).C(col_folderStat)
		err = fscol.Insert(fsts) */
		ok, err := tcdb.InsertFolderStatistics(fsts)
		if err != nil {
			log.Print(err)
		}
	}
	log.Print(len(folders))
	return nil
}

//获取时间段内的项目统计
func (FolderStatService) GetFolderStatByDate(request *pb.GetFolderStatByDateRequest) (*pb.GetFolderStatByDateResponse, error) {
	response := &pb.GetFolderStatByDateResponse{}
	if request.StartDate > request.EndDate {
		response.Result = &pb.ExecuteResponse{Success: false, ErrMsg: errors.ERR_PARAMETER.Error()}
		return response, errors.ERR_PARAMETER
	}
	startTime := time.Unix(request.StartDate, 0)
	endTime := time.Unix(request.EndDate, 0)
	endDate := tool.GetDate(endTime)
	ms, _ := time.ParseDuration("-1ms")
	realEndDate := endDate.AddDate(0, 0, 1).Add(ms)
	timeNow := time.Now().Local()
	dateNow := tool.GetDate(timeNow)
	if endDate.Equal(dateNow) {
		realEndDate = endDate.Add(ms)
	}
	//历史时间段数据
	fstses, err := tcdb.GetFolderDailyStatisticsByDate(request.Folder_Id, startTime, realEndDate)
	if err != nil {
		response.Result = &pb.ExecuteResponse{Success: false, ErrMsg: err.Error()}
		return response, err
	}
	//填充当天实时数据
	if endDate.Equal(dateNow) {
		tasks, err := tcdb.GetTasksByFolderIdAndTime(request.Folder_Id, time.Time{})
		if err != nil {
			response.Result = &pb.ExecuteResponse{Success: false, ErrMsg: err.Error()}
			return response, err
		}
		todayFsts, err := aggregateFolderStats(tasks, request.Folder_Id, timeNow, dateNow)
		if err != nil {
			return nil, err
		}
		if todayFsts != nil {
			fstses = append(fstses, todayFsts)
		}
	}
	/////////////////////////////////排序填充无数据日期START//////////////////////////////////
	startDate := tool.GetDate(startTime)
	dayCount := int(endDate.Sub(startDate).Hours() / 24)

	if len(fstses) != dayCount+1 {
		var tempDate time.Time
		for i := 0; i <= dayCount; i++ {
			tempDate = startDate.AddDate(0, 0, i)

			existFlag := From(fstses).
				AnyWith(
					func(p interface{}) bool {
						return p.(FolderStatistics).Date.Equal(tempDate)
					},
				)
			if !existFlag {
				fstses = append(fstses, &FolderStatistics{Date: tempDate, FolderId: request.Folder_Id})
			}
		}
	}
	///////////////////////////////////排序填充无数据日期END//////////////////////////////
	From(fstses).
		OrderBy(
			func(p interface{}) interface{} {
				return p.(FolderStatistics).Date
			},
		).ToSlice(&fstses)
	return response, nil
}
func (FolderStatService) GetFolderStatNow(request *pb.GetFolderStatNowRequest) (*pb.GetFolderStatNowResponse, error) {
	response := &pb.GetFolderStatNowResponse{}
	tasks, err := tcdb.GetTasksByFolderIdAndTime(request.FolderId, time.Time{})
	if err != nil {
		response.Result = &pb.ExecuteResponse{Success: false, ErrMsg: err.Error()}
		return response, err
	}
	timeNow := time.Now().Local()
	dateNow := tool.GetDate(timeNow)
	fsts, err := aggregateFolderStats(tasks, request.FolderId, timeNow, dateNow)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/* 聚合任务到项目统计实体
tasks:任务数组，folderId:项目id，compareTime:对比的时间,date:统计所属时间段，
createTime未赋值,tasks len=0是返回nil,nil */
func aggregateFolderStats(tasks []*Task, folderId string, compareTime time.Time, date time.Time) (*FolderStatistics, error) {
	tasksCount := len(tasks)
	if tasksCount == 0 {
		return nil, nil
	}
	date = tool.GetDate(date)
	fsts := &FolderStatistics{}
	fsts.FolderId = folderId
	fsts.Date = date
	fsts.Amount = tasksCount
	var (
		chargeIds []string //存负责人总数
		memberIds []string //存成员总数
	)

	for _, task := range tasks {
		chargeIds = append(chargeIds, task.ChargeAccountId)
		for _, member := range task.Members {
			if member.Status == Cs_Normal && (member.Type == Cs_Member || member.Type == Cs_Releaser) {
				memberIds = append(memberIds, member.aId)
			}
		}
		if tool.GetDate(task.CreateTime) == date {
			fsts.NewTasks++
		}

		if task.Status == Cs_Incomplete { //进行中
			fsts.Underway++ //进行中任务书
			if task.Deadline.After(compareTime) {
				fsts.Underway_N++ //正常状态进行中任务数
				if task.Deadline.After(maxDate) {
					fsts.Underway_U++ //未设定截止日期的进行中任务数
				} else {

				}
			} else {
				fsts.Underway_A++                                                  //逾期进行中任务数
				fsts.Timespan_Und += int64(compareTime.Sub(task.Deadline).Hours()) //进行中任务总逾期时间 单位：小时

			}

			if task.Deadline.Before(maxDate) {
				fsts.Timespan_BurnDown += int64(task.Deadline.Sub(compareTime).Hours()) //时间燃尽图
			}

		} else { //已完成
			fsts.Completed++ //已完成任务总数
			if task.Deadline.Before(task.StatusModifyTime) {
				fsts.Timespan_Com += int64(task.StatusModifyTime.Sub(task.Deadline).Hours()) //已完成任务总逾期时间 单位：小时
				fsts.Completed_A++                                                           //逾期已完成任务总数
			} else {
				fsts.Completed_N++ //正常已完成任务数
				if task.Deadline == maxDate {
					fsts.Completed_U++ //未设定截止日期已完成任务数
				}
			}

			if tool.GetDate(task.StatusModifyTime) == date {
				fsts.CompletedTasks++
			}
		}
	}
	if fsts.Underway_A > 0 && fsts.Timespan_Und == 0 {
		fsts.Timespan_Und = 1
	}
	if fsts.Completed_A > 0 && fsts.Timespan_Com == 0 {
		fsts.Timespan_Com = 1
	}
	fsts.Timespan = fsts.Timespan_Und + fsts.Timespan_Com
	From(chargeIds).Distinct().ToSlice(&chargeIds)
	fsts.ChargeAmount = len(chargeIds)
	From(memberIds).Distinct().ToSlice(&memberIds)
	fsts.MemberAmount = len(memberIds)
	return fsts, nil
}
