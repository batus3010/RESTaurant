package main

import (
	"RESTaurant_v2/components/appctx"
	restaurantGin "RESTaurant_v2/modules/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dns: connection stream to MySQL.
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)

	db = db.Debug()
	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		restaurant := v1.Group("/restaurants")
		{
			restaurant.POST("", restaurantGin.CreateRestaurant(appCtx))
			restaurant.GET("/:id", restaurantGin.GetRestaurant(appCtx))
			restaurant.PUT("/:id", restaurantGin.UpdateRestaurant(appCtx))
			restaurant.GET("", restaurantGin.ListRestaurant(appCtx))
			restaurant.DELETE("/:id", restaurantGin.DeleteRestaurant(appCtx))
		}
	}

	r.Run("localhost:8080")
}
