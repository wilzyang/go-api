package file

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wilzyang/go-api/internal/core/domain/file"
	"gorm.io/gorm"
)

type FileList struct {
	Filename     string    `gorm:"filename"`
	Size         int64     `gorm:"size"`
	MediaLink    string    `gorm:"media_link"`
	DownloadLink string    `gorm:"download_link"`
	FileType     string    `gorm:"file_type"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
}

func InsertData(ctx context.Context, data file.FileList, db *gorm.DB) error {

	err := db.Create(&data).Error

	if err != nil {
		e := fmt.Sprintf("Database table Insert : %v", err)
		return errors.New(e)
	}
	return nil
}

func GetData(ctx context.Context, table string, condition map[string]interface{}, db *gorm.DB) (data FileList, err error) {

	err = db.Table(table).Where(condition).Find(data).Error

	return
}
