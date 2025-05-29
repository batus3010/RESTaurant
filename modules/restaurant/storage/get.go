package restaurantStorage

import (
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantModel.Restaurant, error) {
	db := s.db
	var data restaurantModel.Restaurant

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
