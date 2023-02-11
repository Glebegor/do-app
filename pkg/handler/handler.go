package handler

import(
	"github.com/gin-gonic/gin"
)

type handler struct {
}

func (h *handler) initRoutes() *gin.Engine {
	route := gin.New()

	auth := route.Group("/auth",)
	{
		auth.POST("/sing-up")
		auth.POST("/sing-in")
	}
	api := route.Group("/todo")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}
	return route
}
