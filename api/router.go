package api

import (
	"io"
	"log"
	"os"

	"github.com/example/gin_framework/controller"
	"github.com/example/gin_framework/middleware"
	"github.com/example/gin_framework/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/docs"
)

var (
	userService  service.UserService  = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJwtService()

	userController  controller.UserController  = controller.New(userService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func (server *Server) initRouter() {
	// setupLogOutput()

	docs.SwaggerInfo_swagger.Title = "User authentication basic"
	docs.SwaggerInfo_swagger.Description = "User authentication basic - Description"
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = "localhost:8080"
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Schemes = []string{"http"}
	//
	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(gin.Recovery(), middleware.Logger())
	router.POST("/login", loginController.Login)
	// router := gin.Default()
	userRoutes := router.Group("/user", middleware.AuthorizeJWT())
	//  /user/heathCheck
	//  /user/
	userRoutes.GET("/heathCheck", userController.HeathCheck)

	userRoutes.GET("/", userController.GetAllUsers)
	userRoutes.POST("/", userController.Create)

	userRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.router = router

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)

	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
