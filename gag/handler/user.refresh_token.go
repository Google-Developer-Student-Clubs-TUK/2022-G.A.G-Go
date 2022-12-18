package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type tokenReq struct {
	Token string `json:"token" form:"token" binding: "required,token"`
	ID    string `json:"id" form:"id" binding: "required,id"`
	Key   string `json:"key" form:"key" binding: "required,key"`
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req tokenReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		ID:    req.ID,
		Token: req.Token,
	}

	err := h.UserService.SetToken(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(
		"token registerd",
	)

	c.IndentedJSON(http.StatusOK, res)
}
