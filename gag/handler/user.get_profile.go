package handler

import (
	"fmt"
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type profileRes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Major            string `json:"major"`
	ImageURL         string `json:"image_url"`
	IsAlarmOn        bool   `json:"isAlarmOn"`
	IsProfileVisible bool   `json:"isProfileVisible"`
}

func (h *Handler) GetProfile(c *gin.Context) {

	var req defaultReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		ID: req.ID,
	}

	err := h.UserService.GetProfile(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}
	fmt.Print("messages")
	data := profileRes{
		Id:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		ImageURL:         user.ImageURL,
		Major:            "컴퓨터공학과",
		IsAlarmOn:        user.IsAlarm,
		IsProfileVisible: user.IsProfileVisible,
	}

	res := app.NewSuccess(data)

	c.IndentedJSON(http.StatusOK, res)
}
