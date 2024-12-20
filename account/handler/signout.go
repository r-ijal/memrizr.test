package handler

import (
	"memrizr/account/model"
	"memrizr/account/model/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"hello": "it's signout",
	// })

	user := c.MustGet("user")

	ctx := c.Request.Context()
	if err := h.TokenService.Signout(ctx, user.(*model.User).UID); err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user signed out successfully!",
	})
}