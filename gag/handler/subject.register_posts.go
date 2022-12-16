package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type registerPostReq struct {
	ID        string `json:"id" form:"id" binding: "required,id"`
	SubjectId string `json:"subject_id" form:"subject_id" binding: "required,subject_id"`
	Title     string `json:"title" form:"title" binding: "required,title"`
	Content   string `json:"content" form:"content" binding: "required,content"`
}

type registerPostRes struct {
	ID uint `json:"id" form:"id" binding: "required,id"`
}

func (h *Handler) RegisterPost(c *gin.Context) {
	var req registerPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	post := &model.Post{
		Writer:  req.ID,
		SID:     req.SubjectId,
		Title:   req.Title,
		Content: req.Content,
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
		ID: post.ID,
	})

	c.IndentedJSON(http.StatusOK, res)
}
