package routes

import (
	"gin-api-project/controllers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"}, // 許可するオリジン
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                 // 許可するHTTPメソッド
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},      // 許可するヘッダー
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Cookie などの認証情報を許可
		MaxAge:           12 * time.Hour, // プリフライトリクエストのキャッシュ時間
	}))

	api := r.Group("/api")

	api.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"text": "hello",
		})
	})

	{
		api.GET("/user", controllers.GetItems)
		api.POST("/user", controllers.CreateItem)
		// PUTとDELETEのルートも追加

		api.GET("/policy", controllers.ListPolicies)
	}

	return r
}
