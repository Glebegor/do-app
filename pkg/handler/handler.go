package handler

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
			
			item := lists.Group(":id/items")
			{
				item.POST('/')
				item.GET('/')
				item.GET('/:item_id')
				item.PUT('/:item_id')
				item.DELETE('/:item_id')
			}
		}
	}
}
