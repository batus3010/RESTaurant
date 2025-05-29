package restaurantModel

import "RESTaurant_v2/common"

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"` // tag
	Address string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }
