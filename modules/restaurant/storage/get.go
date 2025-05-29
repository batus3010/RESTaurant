package restaurantStorage

import (
	"RESTaurant_v2/common"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantModel.Restaurant, error) {
	db := s.db
	var data restaurantModel.Restaurant

	if err := db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDataNotFound
		}
		return nil, err
	}

	return &data, nil
}
