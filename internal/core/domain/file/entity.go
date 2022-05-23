package file

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	IsError bool   `json:"is_error"`
	Data    string `json:"data"`
}

//data inserted to database
type FileList struct {
	gorm.Model
	Filename     string    `gorm:"filename"`
	Size         int64     `gorm:"size"`
	MediaLink    string    `gorm:"media_link"`
	DownloadLink string    `gorm:"download_link"`
	FileType     string    `gorm:"file_type"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
}
