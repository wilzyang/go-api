package box

import (
	"context"

	"github.com/wilzyang/go-api/domain/box"
	"gorm.io/gorm"
)

type BoxRepository struct {
	DB *gorm.DB
}

func NewBoxRepository(db *gorm.DB) *BoxRepository {
	return &BoxRepository{
		DB: db,
	}
}

var (
	ctx = context.Background()
)

func (r *BoxRepository) DoInsertData(title string, size int, link string, file_id string) (err error) {
	err = InsertData(ctx, title, size, link, file_id, r.DB)
	return
}

func (r *BoxRepository) DoCheckIdByFilekey(filekey string) (data box.ResearchReport, err error) {
	Id, err := CheckIdByFilekey(ctx, filekey, r.DB)

	//error checks
	if err != nil {
		return data, err
	}
	fileid := ResearchReport{
		FileId: Id,
	}

	return toDomainFileId(fileid), err
}

func (r *BoxRepository) DoDeleteByFilekey(filekey string) (err error) {
	err = DeleteByFilekey(ctx, filekey, r.DB)
	return
}
