package handler

import (
	"net/http"
	"strconv"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type editPostReq struct {
	ID      string `json:"id" form:"id" binding: "required,id"`
	Title   string `json:"title" form:"title" binding: "required,title"`
	Content string `json:"content" form:"content" binding: "required,content"`
}

type eidtPostRes struct {
	ID string `json:"id" form:"id"`
}

func (h *Handler) EditPost(c *gin.Context) {
	var req editPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	pid, err := strconv.ParseUint(req.ID, 10, 32)
	if err != nil {
		return
	}

	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
	}
	post.ID = uint(pid)

	err = h.SubjectService.EditPost(c, post)
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(eidtPostRes{
		ID: req.ID,
	})

	c.IndentedJSON(http.StatusOK, res)
}
