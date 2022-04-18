package user

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/h1ght1me/auth-micro/config"
	dbuser "github.com/h1ght1me/auth-micro/pkg/database/user"
	"github.com/h1ght1me/auth-micro/pkg/utils"
	"github.com/h1ght1me/auth-micro/web"
)

type Service struct {
	Service *dbuser.Service
	Config  *config.Config
}

func NewService(userService *dbuser.Service) *Service {
	return &Service{Service: userService, Config: userService.Config}
}

func (s *Service) Authenticator(c *gin.Context) (interface{}, error) {
	var authData web.Auth
	if err := c.ShouldBind(&authData); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	user, err := s.Service.GetByName(authData.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, jwt.ErrFailedAuthentication
		}
		return nil, err
	}
	if utils.CheckPasswordHash(authData.Password, user.Password) {
		return &web.JWTData{ID: user.ID, Name: user.Name, Access: user.Access}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

func (s *Service) GetUsers(c *gin.Context) {
	users, err := s.Service.Users()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, web.APIResponse{Errors: "internal server error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *Service) GetUser(c *gin.Context) {
	id := c.Param("userid")
	userID, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIResponse{Errors: "userid must be an integer"})
		return
	}
	user, err := s.Service.User(userID)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			web.APIResponse{Errors: fmt.Sprintf("user with id %d not found", userID)},
		)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Service) Register(c *gin.Context) {
	authData := new(web.Auth)
	err := c.BindJSON(authData)
	if err != nil {
		fmt.Println(err)
		return
	}
	u := new(dbuser.User)
	u.Name = authData.Name
	u.Password, err = utils.HashPassword(authData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIResponse{Errors: err})
	}
	err = s.Service.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIResponse{Errors: err})
	}
	c.JSON(http.StatusCreated, web.APIResponse{OK: true, Result: u})
}

func (s *Service) DeleteUser(c *gin.Context) {
	id := c.Param("userid")

	userID, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIResponse{Errors: "userid must be an integer"})
		return
	}

	_, err = s.Service.User(userID)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			web.APIResponse{Errors: fmt.Sprintf("user with id %d not found", userID)},
		)
		return
	}
	if err = s.Service.DeleteUser(userID); err != nil {
		log.Println(err)
		c.JSON(
			http.StatusBadRequest,
			web.APIResponse{Errors: fmt.Sprintf("user with id %d not found", userID)},
		)
	} else {
		c.JSON(http.StatusOK, web.APIResponse{OK: true})
	}
}
