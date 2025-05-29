package restaurantGin

import (
	restaurantBiz "RESTaurant_v2/modules/restaurant/biz"
	restaurantStorage "RESTaurant_v2/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRestaurant(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSqlStore(db)
		biz := restaurantBiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dataId": data})
	}
}
