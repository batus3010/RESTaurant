package restaurantBiz

import (
	"RESTaurant_v2/common"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
)

//const (
//	ErrNameIsBlank    = "name cannot be blank"
//	ErrAddressIsBlank = "address cannot be blank"
//)

type createNewRestaurantBiz struct {
	store CreateRestaurantStore
}

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error
}

func NewCreateNewRestaurantBiz(store CreateRestaurantStore) *createNewRestaurantBiz {
	return &createNewRestaurantBiz{store: store}
}

func (biz createNewRestaurantBiz) CreateNewRestaurant(
	ctx context.Context,
	data *restaurantModel.RestaurantCreate,
) error {
	if err := data.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(restaurantModel.EntityName, err)
	}

	return nil
}
