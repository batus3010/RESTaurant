package main

import (
	"RESTaurant_v2/components/appctx"
	"RESTaurant_v2/components/uploadprovider"
	"RESTaurant_v2/middleware"
	restaurantGin "RESTaurant_v2/modules/restaurant/transport/gin"
	uploadGin "RESTaurant_v2/modules/upload/transport/gin"
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
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)

	db = db.Debug()

	s3Provider := uploadprovider.NewAWSS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(db, s3Provider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx)) // global middleware

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		v1.POST("/upload", uploadGin.UploadImage(appCtx))
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
