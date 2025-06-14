package restaurantGin

import (
	"RESTaurant_v2/common"
	"RESTaurant_v2/components/appctx"
	restaurantBiz "RESTaurant_v2/modules/restaurant/biz"
	restaurantModel "RESTaurant_v2/modules/restaurant/model"
	restaurantStorage "RESTaurant_v2/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantModel.RestaurantCreate
		if err := c.ShouldBind(&newData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewCreateNewRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(c, &newData); err != nil {
			panic(err)
		}

		newData.Mask(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.FakeID))
	}
}
