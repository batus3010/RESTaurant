package common

import (
	"errors"
)

var (
	ErrDataNotFound    = errors.New("data not found")
	ErrDataBeenDeleted = errors.New("data has been deleted")
	ErrNameIsBlank     = errors.New("name cannot be blank")
	ErrAddressIsBlank  = errors.New("address cannot be blank")
)

const (
	DbTypeRestaurant = 1
)
