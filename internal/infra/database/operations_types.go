package database

import "gorm.io/gorm"

type OperationsTypesRepository struct {
	Db *gorm.DB
}
