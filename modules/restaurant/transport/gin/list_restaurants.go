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

func ListRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		var filter restaurantModel.Filter

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := paging.Process(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
