package restaurantStorage

import (
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
