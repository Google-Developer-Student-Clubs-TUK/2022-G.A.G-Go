package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type registerPostReq struct {
	writer  string `json:"writer" form:"writer" binding: "required,writer"`
	sid     string `json:"sid" form:"sid" binding: "required,sid"`
	title   string `json:"title" form:"title" binding: "required,title"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type registerPostRes struct {
	pid int `json:"pid" form:"pid" binding: "required,pid"`
}

func (h *Handler) RegisterPost(c *gin.Context) {
	var req registerPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	post := &model.Post{
		Writer:  req.writer,
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
		pid: int(post.ID),
	})

	c.IndentedJSON(http.StatusOK, res)
}
