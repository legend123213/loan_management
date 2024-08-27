package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/legend123213/loan_managment/Domain"
	usecase "github.com/legend123213/loan_managment/Usecase"
)

type UserController struct {
	UserUsecase usecase.UserServiceUsecase
}

	

func NewUserController(usecase usecase.UserServiceUsecase) *UserController {
	return &UserController{
		UserUsecase: usecase,
	}
}

func (controller *UserController) CreateUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var user *domain.User // Change the type to a pointer to domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}		
		user, err := controller.UserUsecase.AddUser(user) // Remove the address-of operator
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func (controller *UserController) GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		user, err := controller.UserUsecase.GetUser(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

// Remove the duplicate declaration of GetUsers method

func (controller *UserController) UpdateUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var user *domain.User // Change the type to a pointer to domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}		
		user, err := controller.UserUsecase.UpdateUser(user) // Remove the address-of operator
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func (controller *UserController) DeleteUser() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		err := controller.UserUsecase.DeleteUser(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User deleted successfully",
		})
	}
}

func (controller *UserController) FatchUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		users, err := controller.UserUsecase.GetUsers()
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}
