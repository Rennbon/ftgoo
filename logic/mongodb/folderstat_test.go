package mongodb

import (
	pb "ftgoo/logic/folderstat"
	"log"
	"testing"
)

type ucTest struct {
	in, out string
}

var fss FolderStatService

func TestDailyFlushing(t *testing.T) {
	fss.DailyFlushing()
}
func TestGetFolderStatNow(t *testing.T) {
	r, err := fss.GetFolderStatNow(&pb.GetFolderStatNowRequest{FolderId: ""})
	log.Println(r, err)
}

func TestGetFolderStatByDate(t *testing.T) {
	r, err := fss.GetFolderStatByDate(
		&pb.GetFolderStatByDateRequest{
			FolderId: "",
			//StartDate: time.Now(),
			//EndDate:   time.Now(),
		})
	log.Println(r, err)
}
