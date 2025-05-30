package restaurantModel

import (
	"RESTaurant_v2/common"
	"strings"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string        `json:"name" gorm:"column:name;"`
	Address string        `json:"addr" gorm:"column:addr;"`
	Logo    *common.Image `json:"logo" gorm:"column:logo;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return common.ErrNameIsBlank
	}
	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return common.ErrAddressIsBlank
	}
	return nil
}
