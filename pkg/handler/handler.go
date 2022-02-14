package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		api.POST("*longUrl", h.h1)
		api.GET("*shortUrl", h.h2)
	}

	return router
}

func (h *Handler) h1(c *gin.Context) {
	fmt.Println("print OK POST")
	c.JSON(200, "OK POST")
}
func (h *Handler) h2(c *gin.Context) {
	fmt.Println("print OK GET")
	c.JSON(200, "OK GET")
}

func (h *Handler) saveLongUrl(c *gin.Context) {
	longUrl := c.Param("longUrl")

	if longUrl == "" {
		c.AbortWithStatusJSON(500, struct {
			Message string `json:"message"`
		}{Message: "Invalid url (empty)"})
		return
	}

	shortUrl, err := h.services.SaveLongUrl(longUrl)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"shortUrl": shortUrl})
}

func (h *Handler) getLongUrl(c *gin.Context) {
	longUrl, err := h.services.GetLongUrl(c.Param("shortUrl"))
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"longUrl": longUrl})
}
