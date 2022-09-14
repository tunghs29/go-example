package main

import (
	"Test/component"
	"Test/database"
	"Test/modules/product/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
)

func main() {

	secretKey := os.Getenv("SecretKey")

	port := os.Getenv("Port")

	client, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Connection Failed to Database")
		log.Fatal(err)
	}

	if err := runService(client, secretKey, port); err != nil {
		log.Fatal("Cannot run service. Try again.")
		log.Fatalln(err)
	}
}

func runService(client *mongo.Client, secretKey string, port string) error {
	c := cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowAllOrigins: true,
		AllowFiles:      true,
	})
	appCtx := component.NewAppContext(client, secretKey)

	routers := gin.Default()
	routers.SetTrustedProxies([]string{"172.16.0.4"})
	routers.Use(c)

	routers.GET("/heath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := routers.Group("/v1")

	{
		product := v1.Group("/products")
		{
			product.GET("/", controller.ListProduct(appCtx))
			product.POST("/", controller.ImportMasterData(appCtx))
		}
	}

	return routers.Run(":" + port)
}
