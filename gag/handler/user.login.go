package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type loginReq struct {
	UUID     string `json:"uuid" form:"uuid" binding: "required,uuid"`
	Key      string `json:"key" form:"key" binding: "required,key"`
	ID       string `json:"id" form:"id" binding: "required,id"`
	Password string `json:"password" form:"password" binding: "required,password"`
}

type loginRes struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Major    string `json:"major"`
	ImageURL string `json:"image_url"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		UUID:        req.UUID,
		ID:          req.ID,
		AesPassword: req.Password,
	}

	err := h.UserService.Login(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(loginRes{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		ImageURL: user.ImageURL,
		Major:    "컴퓨터공학과",
	})

	c.IndentedJSON(http.StatusOK, res)
}
