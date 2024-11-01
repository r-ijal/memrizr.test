package handler

import (
	"log"
	"net/http"
	// "time"
	"memrizr/account/model"
	"memrizr/account/model/apperrors"

	"github.com/gin-gonic/gin"
)

type signinReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signin handler
func (h *Handler) Signin(c *gin.Context) {
	// time.Sleep(1 * time.Second) // demonstrate a timeout
	// c.JSON(http.StatusOK, gin.H{
	// 	"hello": "it's signin",
	// })
	var req signinReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	ctx := c.Request.Context()
	err := h.UserService.Signin(ctx, u)

	if err != nil {
		log.Printf("Failed to sign in user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"tokens": tokens,
	// })

	// Custom response structure to match expected JSON format
	c.JSON(http.StatusOK, gin.H{
		"tokens": gin.H{
			"idToken":      tokens.IDToken.SS,
			"refreshToken": tokens.RefreshToken.SS,
		},
	})
}
