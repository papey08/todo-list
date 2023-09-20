package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/internal/app"
)

// New creates http server which works with given app
func New(addr string, a app.App) *http.Server {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	api := router.Group("todo-list/api")
	appRouter(api, a)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
