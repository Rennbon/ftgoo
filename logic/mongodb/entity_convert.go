package mongodb

import (
	"time"

	pb "github.com/Rennbon/ftgoo/logic/folderstat"
)

func cvt_mg_pb_folderstatses(models []*FolderStatistics) []*pb.FolderStatistics {
	var dtos []*pb.FolderStatistics
	for _, model := range models {
		dto := cvt_mg_pb_folderstatsone(model)
		dtos = append(dtos, dto)
	}
	return dtos
}

func cvt_mg_pb_folderstatsone(model *FolderStatistics) *pb.FolderStatistics {
	if model == nil {
		return nil
	}
	return &pb.FolderStatistics{
		FolderId:         model.FolderId,                   //1
		CreateTime:       time.Time.Unix(model.CreateTime), //2
		Date:             time.Time.Unix(model.Date),       //3
		Amount:           model.Amount,                     //4
		Underway:         model.Underway,                   //5
		UnderwayA:        model.Underway_A,                 //6
		UnderwayN:        model.Underway_N,                 //7
		UnderwayU:        model.Underway_U,                 //8
		Timespan:         model.Timespan,                   //9
		TimespanUnd:      model.Timespan_Und,               //10
		TimespanCom:      model.Timespan_Com,               //11
		Completed:        model.Completed,                  //12
		CompletedN:       model.Completed_N,                //13
		CompletedA:       model.Completed_A,                //14
		CompletedU:       model.Completed_U,                //15
		MemberAmount:     model.MemberAmount,               //16
		ChargeAmount:     model.ChargeAmount,               //17
		NewTasks:         model.NewTasks,                   //18
		CompletedTasks:   model.CompletedTasks,             //19
		TimespanBurnDown: model.Timespan_BurnDown,          //t20
	}
}
