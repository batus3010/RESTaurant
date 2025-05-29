package restaurantStorage

import (
	"gorm.io/gorm"
)

// encapsulation
type sqlStore struct {
	db *gorm.DB
}
