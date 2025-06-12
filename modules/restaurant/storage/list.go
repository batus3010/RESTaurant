package restaurantStorage

import (
	"RESTaurant_v2/common"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantModel.Restaurant, error) {
	db := s.db
	var result []restaurantModel.Restaurant

	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}

	db = db.Where("status NOT IN (0)") // for deleting

	if err := db.Table(restaurantModel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	if err := db.Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}
	return result, nil
}
