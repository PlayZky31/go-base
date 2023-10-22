package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/PlayZky31/go-base/internal/common"
	"github.com/PlayZky31/go-base/internal/domain"
	"github.com/PlayZky31/go-base/pkg/customError"
)

type EmployeeRepo interface {
	GetEmployee(ctx context.Context, id string) (result domain.Employee, err error)
	GetListEmployee(ctx context.Context, commonQuery *common.QueryPagination) (result []domain.Employee, err error)
	CreateEmployee(ctx context.Context, ent *domain.Employee) (err error)
	UpdateEmployee(ctx context.Context, id string, ent *domain.Employee) (err error)
	DeleteEmployee(ctx context.Context, id string) (err error)
	CheckIfExist(ctx context.Context, id string) (err error)
}

func (r repoContainer) GetListEmployee(ctx context.Context, commonQuery *common.QueryPagination) (result []domain.Employee, err error) {
	exec := r.dbClient.Table(TableEmployee).Select("id,first_name,last_name,email,TO_CHAR(hire_date, 'YYYY-MM-DD') AS hire_date")

	if commonQuery.SearchKeyword != "" {
		exec = exec.Where("first_name LIKE '%@val%' OR last_name LIKE '%@val%' OR email LIKE '%@val%'", sql.Named("val", commonQuery.SearchKeyword))
	}

	getCount := exec.Count(&commonQuery.TotalItems)
	if getCount.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	exec = exec.Scopes(paginate(commonQuery)).Find(&result)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	return
}

func (r repoContainer) GetEmployee(ctx context.Context, id string) (result domain.Employee, err error) {
	exec := r.dbClient.Table(TableEmployee).Select("id,first_name,last_name,email,TO_CHAR(hire_date, 'YYYY-MM-DD') AS hire_date").Where("id=?", id).First(&result)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
			err = customError.ErrNotFound(TableEmployee)
		}
		return
	}
	return
}

func (r repoContainer) CreateEmployee(ctx context.Context, ent *domain.Employee) (err error) {
	exec := r.dbClient.Table(TableEmployee).Create(ent)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if exec.RowsAffected < 1 {
		err = customError.ErrQuery(fmt.Errorf("failed to create employee"))
		return
	}

	return
}

func (r repoContainer) UpdateEmployee(ctx context.Context, id string, ent *domain.Employee) (err error) {
	exec := r.dbClient.Table(TableEmployee).Where("id=?", id).Updates(ent)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if exec.RowsAffected < 1 {
		err = customError.ErrQuery(fmt.Errorf("failed to update employee"))
		return
	}

	return
}

func (r repoContainer) DeleteEmployee(ctx context.Context, id string) (err error) {
	exec := r.dbClient.Table(TableEmployee).Where("id=?", id).Delete(&domain.Employee{})
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if exec.RowsAffected < 1 {
		err = customError.ErrQuery(fmt.Errorf("failed to delete employee"))
		return
	}

	return
}
