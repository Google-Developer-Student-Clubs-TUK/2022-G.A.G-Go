package handler

import (
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type editCommentReq struct {
	cid     int    `json:"cid" form:"cid" binding: "required,cid"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type eidtCommentRes struct {
	cid int `json:"cid" form:"cid"`
}

func (h *Handler) EditComment(c *gin.Context) {
	var req editCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	comment := &model.Comment{
		Content: req.content,
	}
	comment.ID = uint(req.cid)

	err := h.SubjectService.EditComment(c, comment)
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(eidtCommentRes{
		cid: req.cid,
	})

	c.IndentedJSON(http.StatusOK, res)
}
