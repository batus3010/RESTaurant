package restaurantGin

import (
	restaurantBiz "RESTaurant_v2/modules/restaurant/biz"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	restaurantStorage "RESTaurant_v2/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantModel.RestaurantCreate
		if err := c.ShouldBind(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewCreateNewRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(c, &newData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dataId": newData.Id})
	}
}
