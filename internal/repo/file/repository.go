package file

import (
	"context"

	"github.com/wilzyang/go-api/internal/core/domain/file"
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

func (r *FileRepository) DoInsertData(title string, data file.FileList) (err error) {
	err = InsertData(ctx, data, r.DB)
	return
}
