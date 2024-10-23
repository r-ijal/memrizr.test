package handler

import (
	// "net/http"

	"log"
	"memrizr/account/model"
	"memrizr/account/model/apperrors"

	"github.com/gin-gonic/gin"
)

// signupReq is not exported, hence the lowercase name
// it is used for validation and json marshalling
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"hello": "it's signup",
	// })

	var req signupReq

	// Bind incoming json to struct and check for validation errors
	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email: req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
}
