package handler

import (
	"memrizr/account/handler/middleware"
	"memrizr/account/model"
	"memrizr/account/model/apperrors"
	"net/http"
	"time"

	// "os"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// Config will hold services that will eventually be injected
// into this handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	BaseURL         string
	TimeoutDuration time.Duration
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly wit a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	// Create an account group
	// g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))
	g := c.R.Group(c.BaseURL)

	if gin.Mode() != gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
	} else {
		g.GET("me", h.Me)
	}

	// g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
