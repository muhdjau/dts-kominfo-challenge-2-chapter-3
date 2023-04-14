package routers

import (
	"challenge-chapter-3-sesi-2/controllers"
	"challenge-chapter-3-sesi-2/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	roleRouter := router.Group("/role")
	{
		roleRouter.POST("/", controllers.CreateRole)
		roleRouter.GET("/", controllers.GetRole)
	}

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := router.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/create", controllers.CreateProduct)
		productRouter.GET("/all", controllers.GetAllProducts)
		productRouter.GET("/:productID", middlewares.ProductAuthorization(), controllers.GetProductById)
		productRouter.PUT("/:productID", middlewares.RoleMiddleWare(), middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productID", middlewares.RoleMiddleWare(), middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return router
}