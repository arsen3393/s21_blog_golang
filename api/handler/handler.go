package handler

import (
	"Go_Day06/models"
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

type Handler struct {
	UserModel models.UserModel
	PostModel models.PostModel
}

func NewHandler(userModel models.UserModel, postModel models.PostModel) *Handler {
	return &Handler{
		UserModel: userModel,
		PostModel: postModel,
	}
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func (handler *Handler) SetupRouters() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"}, // Разрешить запросы с любых источников (можно заменить на конкретные URL)
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second, // 100 секунда
		Limit: 100,         // Лимит 5 запросов
	})

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
	})

	router.Use(mw)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.json")))
	RegisterUserRoutes(router, &handler.UserModel)
	RegisterPostRoutes(router, &handler.PostModel)
	return router
}
