package auth

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"log"
	"time"

	dbuser "github.com/h1ght1me/auth-micro/pkg/database/user"
	"github.com/h1ght1me/auth-micro/pkg/permissions"
	"github.com/h1ght1me/auth-micro/pkg/utils"
	"github.com/h1ght1me/auth-micro/web"
)

func Middleware(userDB *dbuser.Service) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "jwt auth",
		Key:         []byte(userDB.Config.Server.SecretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*JWTData); ok {
				return jwt.MapClaims{identityKey: v.ID, "name": v.Name, "access": v.Access}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id, ok := claims[identityKey].(string)
			if !ok {
				log.Panicln("can't extract userID from jwt claims")
			}
			userID, err := uuid.FromString(id)
			if err != nil {
				log.Panicf("can't extract userID from jwt claims: %s", err)
			}
			access, ok := claims["access"].(float64)
			if !ok {
				log.Panicln("can't extract access from jwt claims")
			}
			return &JWTData{
				ID:     userID,
				Name:   claims["name"].(string),
				Access: permissions.Permission(access),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var authData Auth
			if err := c.ShouldBind(&authData); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			user, err := userDB.GetByName(authData.Name)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, jwt.ErrFailedAuthentication
				}
				return nil, err
			}
			if utils.CheckPasswordHash(authData.Password, user.Password) {
				return &JWTData{ID: user.ID, Name: user.Name, Access: user.Access}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.JSON(code, web.APIResponse{OK: true, Result: LoginResponse{Token: message}})
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			log.Printf("path = %s\n", c.Request.URL.Path)
			if v, ok := data.(*JWTData); ok && v.Access&permissions.Access > 0 {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, web.APIResponse{Errors: message})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware, err
}
