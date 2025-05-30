package restaurantModel

import "RESTaurant_v2/common"

type Restaurant struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"` // tag
	Address string         `json:"addr" gorm:"column:addr;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }
