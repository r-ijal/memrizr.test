package handler

import (
	"log"
	"memrizr/account/model/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokensReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"hello": "it's tokens",
	// })

	var req tokensReq

	if ok := bindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	// verify refresh JWT
	// refreshToken consist of SS, UID, tokenId
	refreshToken, err := h.TokenService.ValidateRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// get up-to-date user
	u, err := h.UserService.Get(ctx, refreshToken.UID)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create fresh pair of tokens
	tokens, err := h.TokenService.NewPairFromUser(ctx, u, refreshToken.ID.String())

	if err != nil {
		log.Printf("Failed to create tokens for user: %+v. Error: %v\n", u, err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}