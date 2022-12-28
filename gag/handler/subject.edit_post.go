package handler

import (
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type editPostReq struct {
	pid     int    `json:"pid" form:"pid" binding: "required,pid"`
	title   string `json:"title" form:"title" binding: "required,title"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type eidtPostRes struct {
	pid int `json:"pid" form:"pid"`
}

func (h *Handler) EditPost(c *gin.Context) {
	var req editPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	post := &model.Post{
		Title:   req.title,
		Content: req.content,
	}
	post.ID = uint(req.pid)

	err := h.SubjectService.EditPost(c, post)
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(eidtPostRes{
		pid: req.pid,
	})

	c.IndentedJSON(http.StatusOK, res)
}
