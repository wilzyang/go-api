package file

import (
	"context"

	"gorm.io/gorm"
)

type FileRepository struct {
	DB *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{
		DB: db,
	}
}

var (
	ctx = context.Background()
)

func (r *FileRepository) DoInsertData(title string, size int, link string, file_type string) (err error) {
	err = InsertData(ctx, title, size, link, file_type, r.DB)
	return
}
