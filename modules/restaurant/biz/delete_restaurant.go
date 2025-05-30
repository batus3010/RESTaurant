package restaurantBiz

import (
	"RESTaurant_v2/common"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
	"errors"
)

// soft delete, only update status to 0 (means deleted)
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

type DeleteRestaurantStore interface {
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

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz deleteRestaurantBiz) DeleteRestaurant(
	ctx context.Context,
	id int,
) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil && errors.Is(err, common.ErrDataNotFound) {
		return common.ErrDataNotFound
	}

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrDataBeenDeleted
	}

	zero := 0

	if err := biz.store.Update(ctx,
		map[string]interface{}{"id": id},
		&restaurantModel.RestaurantUpdate{Status: &zero}); err != nil {
		return err
	}

	return nil
}
