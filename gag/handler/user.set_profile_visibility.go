package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetProfileVisibility(c *gin.Context) {
	var req alarmReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		ID:               req.ID,
		RsaPrivateKey:    req.Key,
		IsProfileVisible: req.Current,
	}

	err := h.UserService.SetProfileVisibility(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(loginRes{
		Name:     user.Name,
		Email:    user.Email,
		ImageURL: user.ImageURL,
	})

	c.IndentedJSON(http.StatusOK, res)
}
