package restaurantStorage

import (
	"gorm.io/gorm"
)

// encapsulation
type sqlStore struct {
	db *gorm.DB
}

//func NewSqlStore(db *gorm.DB) *sqlStore {
//	return &sqlStore{
//		db: db,
//	}
//}
