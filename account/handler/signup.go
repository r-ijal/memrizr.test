package handler

import (
	"net/http"

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

	ctx := c.Request.Context()
	err := h.UserService.Signup(ctx, u)
	// err := h.UserService.Signup(c, u)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

    // create token pair as strings
    tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")

    if err != nil {
        log.Printf("Failed to create tokens for user: %v\n", err.Error())

        // may eventually implement rollback logic here
        // meaning, if we fail to create tokens after creating a user,
        // we make sure to clear/delete the created user in the database

        c.JSON(apperrors.Status(err), gin.H{
            "error": err,
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "tokens": tokens,
    })
}
