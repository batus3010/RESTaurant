package restaurantBiz

import (
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
)

type getRestaurantBiz struct {
	store GetRestaurantStore
}

type GetRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantModel.Restaurant, error)
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id int,
) (*restaurantModel.Restaurant, error) {

	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}
