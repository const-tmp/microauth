package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/h1ght1me/auth-micro/config"
	"github.com/h1ght1me/auth-micro/pkg/utils"
	"github.com/h1ght1me/auth-micro/web"
	"github.com/h1ght1me/auth-micro/web/auth"
	"net/http"

	dbuser "github.com/h1ght1me/auth-micro/pkg/database/user"
)

type Service struct {
	Service *dbuser.Service
	Config  *config.Config
}

func NewService(userService *dbuser.Service) *Service {
	return &Service{
		Service: userService,
		Config:  userService.Config,
	}
}

func (t *Service) GetUsers(c *gin.Context) {
	users, err := t.Service.Users()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (t *Service) GetUser(c *gin.Context) {
	id := c.Param("userid")
	userID, err := uuid.FromString(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "userid must be an integer"})
		return
	}
	user, err := t.Service.User(userID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user with id %d not found", userID)})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (t *Service) Register(c *gin.Context) {
	authData := new(auth.Auth)
	err := c.BindJSON(authData)
	if err != nil {
		fmt.Println(err)
		return
	}
	u := new(dbuser.User)
	u.Name = authData.Name
	u.Password, err = utils.HashPassword(authData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIResponse{
			OK:     false,
			Result: nil,
			Errors: err,
		})
	}
	err = t.Service.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIResponse{
			OK:     false,
			Result: nil,
			Errors: err,
		})
	}
	c.IndentedJSON(http.StatusCreated, web.APIResponse{
		OK:     true,
		Result: u,
		Errors: nil,
	})
}

//func (t *Service) CreateUser(c *gin.Context) {
//	var request CreateUserRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	if err := t.Service.CreateUser(request.User.httpToModel()); err != nil {
//		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
//	} else {
//		c.IndentedJSON(http.StatusCreated, request.User)
//	}
//}

func (t *Service) DeleteUser(c *gin.Context) {
	id := c.Param("userid")

	userID, err := uuid.FromString(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "userid must be an integer"})
		return
	}

	_, err = t.Service.User(userID)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("user with id %d not found", userID),
			},
		)
		return
	}
	if err = t.Service.DeleteUser(userID); err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"message": fmt.Sprintf("user with id %d not found", userID),
			},
		)
	} else {
		c.IndentedJSON(http.StatusOK, struct{}{})
	}
}
