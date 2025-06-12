package restaurantStorage

import (
	"RESTaurant_v2/common"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
)

func (s *sqlStore) Update(
	ctx context.Context,
	condition map[string]interface{},
	updateData *restaurantModel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where(condition).Updates(updateData).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
