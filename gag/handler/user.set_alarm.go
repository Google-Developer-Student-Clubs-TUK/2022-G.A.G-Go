package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type alarmReq struct {
	Key     string `json:"key" form:"key" binding: "required,key"`
	ID      string `json:"id" form:"id" binding: "required,id"`
	Current bool   `json:"isProfileVisible" form:"current" binding: "required"`
}

func (h *Handler) SetAlarm(c *gin.Context) {
	var req alarmReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		ID:            req.ID,
		RsaPrivateKey: req.Key,
		IsAlarm:       req.Current,
	}

	err := h.UserService.SetAlarm(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	result := true
	if req.Current {
		result = false
	}

	res := app.NewSuccess(
		result,
	)

	c.IndentedJSON(http.StatusOK, res)
}
