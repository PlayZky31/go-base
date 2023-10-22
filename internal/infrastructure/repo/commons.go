package repo

import (
	"context"

	"gorm.io/gorm"

	"github.com/PlayZky31/go-base/internal/common"
	"github.com/PlayZky31/go-base/pkg/customError"
)

func (r repoContainer) CheckIfExist(ctx context.Context, id string) (err error) {
	var total int64
	exec := r.dbClient.Table(TableEmployee).Where("id=?", id).Count(&total)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if total < 1 {
		err = customError.ErrNotFound(TableEmployee)
		return
	}

	return
}

func paginate(commonQuery *common.QueryPagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (commonQuery.Page - 1) * commonQuery.ItemsPerPage
		return db.Offset(offset).Limit(commonQuery.ItemsPerPage)
	}
}
