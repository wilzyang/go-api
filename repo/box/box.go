package box

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ResearchReport struct {
	FileKey       string    `gorm:"file_key"`
	Title         string    `gorm:"title"`
	Size          int       `gorm:"size"`
	CreatedAt     time.Time `gorm:"created_at"`
	LastUpdatedAt time.Time `gorm:"last_updated_at"`
	Link          string    `gorm:"box_shared_link"`
	FileId        string    `gorm:"box_file_id"`
	TotalDownload int       `gorm:"total_download"`
}

func InsertData(ctx context.Context, title string, size int, link string, file_id string, db *gorm.DB) error {

	fk := strings.Split(title, ".")
	filekey := fk[0] //get filekey from file name

	filedata := &ResearchReport{
		FileKey:       filekey,
		Title:         title,
		Size:          size,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		Link:          link,
		FileId:        file_id,
		TotalDownload: 0,
	}

	err := db.Select("file_key", "title", "size", "created_at", "last_updated_at", "box_shared_link", "box_file_id", "total_download").Create(filedata).Error

	if err != nil {
		e := fmt.Sprintf("Database table [research_report] Insert : %v", err)
		return errors.New(e)
	}
	return nil
}

func CheckIdByFilekey(ctx context.Context, filekey string, db *gorm.DB) (Id string, err error) {

	err = db.Table("research_report").Select("box_file_id").Where("file_key = ?", filekey).First(&Id).Error

	if err != nil {
		e := fmt.Sprintf("Database table [research_report] query : %v", err)
		return e, errors.New(e)
	}

	return
}

func DeleteByFilekey(ctx context.Context, filekey string, db *gorm.DB) error {

	err := db.Table("research_report").Where("file_key = ?", filekey).Delete(&ResearchReport{})

	if err != nil {
		e := fmt.Sprintf("Database table [research_report] Delete : %v", err)
		return errors.New(e)
	}

	return nil
}
