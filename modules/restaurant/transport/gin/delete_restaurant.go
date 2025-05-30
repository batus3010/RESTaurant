package restaurantGin

import (
	"RESTaurant_v2/components/appctx"
	restaurantBiz "RESTaurant_v2/modules/restaurant/biz"
	restaurantStorage "RESTaurant_v2/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
