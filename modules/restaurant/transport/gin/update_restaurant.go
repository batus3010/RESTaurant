package restaurantGin

import (
	"RESTaurant_v2/common"
	"RESTaurant_v2/components/appctx"
	restaurantBiz "RESTaurant_v2/modules/restaurant/biz"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	restaurantStorage "RESTaurant_v2/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var data restaurantModel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
