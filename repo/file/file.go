package file

import (
	"context"
	"errors"
	"fmt"

	"github.com/wilzyang/go-api/domain/file"
	"gorm.io/gorm"
)

func InsertData(ctx context.Context, data file.FileList, db *gorm.DB) error {

	err := db.Create(data).Error

	if err != nil {
		e := fmt.Sprintf("Database table Insert : %v", err)
		return errors.New(e)
	}
	return nil
}
