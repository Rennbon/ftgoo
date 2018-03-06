package mongodb

import (
	"log"
	"testing"
	"time"

	pb "github.com/Rennbon/ftgoo/logic/folderstat"
)

type ucTest struct {
	in, out string
}

var fss FolderStatService

func TestDailyFlushing(t *testing.T) {
	fss.DailyFlushing(time.Now().Local())
}
func TestGetFolderStatNow(t *testing.T) {
	r, err := fss.GetFolderStatNow(&pb.GetFolderStatNowRequest{FolderId: "eed3c2db-c690-4ed6-ada7-5f0aae18b4c2"})
	log.Println(r, err)
}

func TestGetFolderStatByDate(t *testing.T) {
	r, err := fss.GetFolderStatByDate(
		&pb.GetFolderStatByDateRequest{
			FolderId:  "dfbafb8e-fecc-4f85-8ed7-94318a585e9b",
			StartDate: time.Date(2018, 2, 27, 0, 0, 0, 0, time.Local).Unix(),
			EndDate:   time.Date(2018, 3, 5, 0, 0, 0, 0, time.Local).Unix(),
		})
	for _, v := range r.Folderstats {
		log.Println(v, err)
		log.Println(time.Unix(v.Date, 0).Format("2006-01-02 15:04:05"))
	}

}
