package api

import (
	"api-booking/controller/server/api/admin"
	"api-booking/controller/server/api/admin/account"
	"api-booking/controller/server/api/users"
	"api-booking/middlewares"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func NewHandler(db *sql.DB, client *redis.Client) http.Handler {
	// Force log's color
	gin.ForceConsoleColor()
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("/api")
	{
		v1.GET("/logout", middlewares.AuthJwtVerify(), func(c *gin.Context) { account.LogoutAccount(c, client) })
		admin.Routes(v1, db, client)
		users.Routes(v1, db)
	}

	return router
}
