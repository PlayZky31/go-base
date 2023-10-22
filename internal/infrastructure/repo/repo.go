package repo

import (
	"gorm.io/gorm"
)

type Task interface {
}

type repoContainer struct {
	dbClient *gorm.DB
}

func NewRepo(dbClient *gorm.DB) Task {
	return &repoContainer{
		dbClient: dbClient,
	}
}
