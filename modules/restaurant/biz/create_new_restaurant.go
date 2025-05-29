package restaurantBiz

import (
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	"context"
	"errors"
	"strings"
)

const (
	ErrNameIsBlank    = "name cannot be blank"
	ErrAddressIsBlank = "address cannot be blank"
)

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
	err := ensureNotBlank(data.Name, ErrNameIsBlank)
	if err != nil {
		return err
	}
	err = ensureNotBlank(data.Address, ErrAddressIsBlank)
	if err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func ensureNotBlank(name string, err string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New(err)
	}
	return nil
}
