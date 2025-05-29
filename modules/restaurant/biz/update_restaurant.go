package restaurantBiz

import (
	"RESTaurant_v2/common"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
	"errors"
)

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

type UpdateRestaurantStore interface {
	Update(
		ctx context.Context,
		condition map[string]interface{},
		updateData *restaurantModel.RestaurantUpdate,
	) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantModel.Restaurant, error)
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantModel.RestaurantUpdate,
) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil && errors.Is(err, common.ErrDataNotFound) {
		return errors.New("data not found")
	}

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
