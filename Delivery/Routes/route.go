package routes

// gin-gonic
import (
	"github.com/gin-gonic/gin"
	controller "github.com/legend123213/loan_managment/Delivery/Controller"
)
type Route struct {
	UserController *controller.UserController
}

func NewRoute(userController *controller.UserController) *Route {
	return &Route{
		UserController: userController,
	}
}

func (route *Route) SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.POST("/users", route.UserController.CreateUser())
	r.GET("/users/:id", route.UserController.GetUser())
	r.GET("/users", route.UserController.FatchUsers())
	r.PUT("/users", route.UserController.UpdateUser())
	r.DELETE("/users/:id", route.UserController.DeleteUser())
	return r
}



