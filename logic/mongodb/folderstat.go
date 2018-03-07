package mongodb

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/Rennbon/ftgoo/tool"

	pb "github.com/Rennbon/ftgoo/logic/folderstat"

	"github.com/Rennbon/ftgoo/logic/errors"

	. "github.com/ahmetb/go-linq"
)

type FolderStatService struct {
}

var tcdb taskCenter

//项目统计日刷
func (FolderStatService) DailyFlushing(date time.Time) error {
	defer tool.CallRecover()
	dateRight := tool.GetDate(date)
	dateLeft := dateRight.AddDate(0, -1, 0)
	iceAge := time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local)
	runtime.GOMAXPROCS(runtime.NumCPU())
	chleft := make(chan time.Time)
	chright := make(chan time.Time)
	wg := sync.WaitGroup{}
	go func() {
		i := 0
		for {
			if dateLeft.Before(iceAge) || i > 10000 {
				break
			}
			i++
			wg.Add(1)
			chleft <- dateLeft
			chright <- dateRight
			dateRight = dateLeft
			dateLeft = dateRight.AddDate(0, -1, 0)
		}
		close(chleft)
		close(chright)
	}()
	for left := range chleft {
		go func(left, right time.Time) {
			err := dailyFlushingBetween(left, right)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(left, <-chright)
	}
	wg.Wait()
	return nil
}

func dailyFlushingBetween(start time.Time, end time.Time) error {
	defer tool.CallRecover()
	folders, err := tcdb.GetFoldersCreateTimeBetween(start, end)
	if err != nil {
		log.Print(err)
		return err
	}
	dateNow := tool.GetDate(time.Now().Local())             //当前时间
	fstsDate := dateNow.AddDate(0, 0, -1)                   //当前时间前一天，即需要统计的时间
	ddCompare := dateNow.Add(time.Hour * time.Duration(-1)) //对比时间
	effd := 0
	//遍历项目
	for _, folder := range folders {
		//go func(folder *Folder) {
		tasks, err := tcdb.GetTasksByFolderIdAndTime(folder.FolderId, dateNow)
		if err != nil {
			log.Print(err)
			continue
			//return
		}
		fsts, err := aggregateFolderStats(tasks, folder.FolderId, ddCompare, fstsDate)
		if err != nil {
			log.Print(err)
			continue
			//return
		}
		if fsts == nil {
			continue
			//return
		}
		fsts.CreateTime = time.Now().Local()
		err = tcdb.InsertFolderStatistics(fsts)
		if err != nil {
			log.Print(err)
			continue
			//return
		}
		effd++
		//}(folder)
		//log.Println("folder:", folder.FolderId, "completed!")
	}
	log.Print(start.Format("2006-01-02"), "-", end.Format("2006-01-02"),
		" completed,total:", len(folders), ",real:", effd)
	return nil
}

//获取时间段内的项目统计
func (FolderStatService) GetFolderStatByDate(request *pb.GetFolderStatByDateRequest) (*pb.GetFolderStatByDateResponse, error) {
	defer tool.CallRecover()
	response := &pb.GetFolderStatByDateResponse{}
	if request.StartDate > request.EndDate || !tool.CheckUnix(request.StartDate) || !tool.CheckUnix(request.EndDate) {
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
	fstses, err := tcdb.GetFolderDailyStatisticsByDate(request.FolderId, startTime, realEndDate)
	if err != nil {
		response.Result = &pb.ExecuteResponse{Success: false, ErrMsg: err.Error()}
		return response, err
	}
	//填充当天实时数据
	if endDate.Equal(dateNow) {
		tasks, err := tcdb.GetTasksByFolderIdAndTime(request.FolderId, time.Time{})
		if err != nil {
			response.Result = &pb.ExecuteResponse{Success: false, ErrMsg: err.Error()}
			return response, err
		}
		todayFsts, err := aggregateFolderStats(tasks, request.FolderId, timeNow, dateNow)
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
			if len(fstses) == dayCount+1 {
				break
			}
			tempDate = startDate.AddDate(0, 0, i)

			existFlag := From(fstses).
				AnyWith(
					func(p interface{}) bool {
						return p.(*FolderStatistics).Date.Equal(tempDate)
					},
				)
			if !existFlag {
				fstses = append(fstses, &FolderStatistics{Date: tempDate, CreateTime: tempDate, FolderId: request.FolderId})
			}
		}
	}
	///////////////////////////////////排序填充无数据日期END//////////////////////////////
	var fstsesSort []*FolderStatistics
	From(fstses).
		SortT(
			func(f1 *FolderStatistics, f2 *FolderStatistics) bool {
				return f1.Date.Before(f2.Date)
			},
		).ToSlice(&fstsesSort)
	response.Result = &pb.ExecuteResponse{Success: true}
	response.Folderstats = cvt_mg_pb_folderstatses(fstsesSort)
	return response, nil
}

func (FolderStatService) GetFolderStatNow(request *pb.GetFolderStatNowRequest) (*pb.GetFolderStatNowResponse, error) {
	defer tool.CallRecover()
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
	if fsts != nil {
		fsts = &FolderStatistics{}
		fsts.Date = dateNow
		fsts.FolderId = request.FolderId
		fsts.CreateTime = dateNow
	}
	response.Result = &pb.ExecuteResponse{Success: true, ErrMsg: ""}
	response.Folderstat = cvt_mg_pb_folderstatsone(fsts)
	return response, nil
}

/* 聚合任务到项目统计实体
tasks:任务数组，folderId:�������������id，compareTime:��比�����时间,date:统计所属时间段，
createTime未赋值,tasks len=0��返回nil,nil */
func aggregateFolderStats(tasks []*Task, folderId string, compareTime time.Time, date time.Time) (*FolderStatistics, error) {
	defer tool.CallRecover()
	tasksCount := int32(len(tasks))
	if tasksCount == 0 {
		return nil, nil
	}
	date = tool.GetDate(date)
	fsts := &FolderStatistics{}
	fsts.FolderId = folderId
	fsts.Date = date
	fsts.CreateTime = date
	fsts.Amount = tasksCount
	var (
		chargeIds []string //存负责人总数
		memberIds []string //存成员总数
	)

	for _, task := range tasks {
		go func(task *Task) {
			chargeIds = append(chargeIds, task.ChargeAccountId)
			for _, member := range task.Members {
				if member.Status == Cs_Normal && (member.Type == Cs_Member || member.Type == Cs_Releaser) {
					memberIds = append(memberIds, member.AId)
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
		}(task)
	}
	if fsts.Underway_A > 0 && fsts.Timespan_Und == 0 {
		fsts.Timespan_Und = 1
	}
	if fsts.Completed_A > 0 && fsts.Timespan_Com == 0 {
		fsts.Timespan_Com = 1
	}
	fsts.Timespan = fsts.Timespan_Und + fsts.Timespan_Com
	From(chargeIds).Distinct().ToSlice(&chargeIds)
	fsts.ChargeAmount = int32(len(chargeIds))
	From(memberIds).Distinct().ToSlice(&memberIds)
	fsts.MemberAmount = int32(len(memberIds))
	return fsts, nil
}
