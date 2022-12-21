package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type registerCommentReq struct {
	uid     string `json:"uid" form:"uid" binding: "required,uid"`
	pid     string `json:"pid" form:"pid" binding: "required,pid"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type registerCommentRes struct {
	cid uint `json:"cid" form:"cid" binding: "required,cid"`
}

func (h *Handler) RegisterComment(c *gin.Context) {
	var req registerCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	comment := &model.Comment{
		Writer:  req.ID,
		Content: req.Content,
	}

	err := h.SubjectService.RegisterComment(c, comment)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(registerCommentRes{
		ID: comment.ID,
	})

	c.IndentedJSON(http.StatusOK, res)
}
