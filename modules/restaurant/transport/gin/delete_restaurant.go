package restaurantGin

import (
	"RESTaurant_v2/common"
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
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
