package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tm/internal/cfg"
)

func NewGormDb(configuration cfg.Cfg) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(configuration.DbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
