package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type registerPostReq struct {
	id      string `json:"id" form:"id" binding: "required,id"`
	sid     string `json:"sid" form:"sid" binding: "required,sid"`
	title   string `json:"title" form:"title" binding: "required,title"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type registerPostRes struct {
	id uint `json:"id" form:"id" binding: "required,id"`
}

func (h *Handler) RegisterPost(c *gin.Context) {
	var req registerPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	post := &model.Post{
		Writer:  req.id,
		SID:     req.sid,
		Title:   req.title,
		Content: req.content,
	}

	err := h.SubjectService.RegisterPost(c, post)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(registerPostRes{
		id: post.ID,
	})

	c.IndentedJSON(http.StatusOK, res)
}
