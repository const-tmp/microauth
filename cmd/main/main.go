package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"github.com/h1ght1me/auth-micro/config"
	"github.com/h1ght1me/auth-micro/pkg/database"
	dbuser "github.com/h1ght1me/auth-micro/pkg/database/user"
	"github.com/h1ght1me/auth-micro/web"
	"github.com/h1ght1me/auth-micro/web/auth"
	"github.com/h1ght1me/auth-micro/web/user"
)

// main entry point
func main() {
	env, exists := os.LookupEnv("ENV")
	if !exists {
		env = "dev"
	}
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("can't load config: %s", err)
	}
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	err = database.RunMigration(db)
	if err != nil {
		log.Fatal("migration failed")
	}

	userDB := dbuser.NewService(db, cfg)
	userWeb := user.NewService(userDB)
	authMiddleware, err := auth.Middleware(userWeb)
	if err != nil {
		log.Fatal("creating auth middleware failed")
	}
	router := gin.Default()

	public := router.Group("/")
	{
		public.POST("/login", authMiddleware.LoginHandler)
		public.POST("/register", userWeb.Register)
	}

	private := router.Group("/")
	//{
	//	private.POST("/login", loginEndpoint)
	//	private.POST("/submit", submitEndpoint)
	//	private.POST("/read", readEndpoint)
	//}
	private.Use(authMiddleware.MiddlewareFunc())

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(
			http.StatusNotFound,
			web.APIResponse{
				OK:     false,
				Errors: "resource not found",
			},
		)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router); err != nil {
		log.Fatal(err)
	}
}
