package handler

import (
	"net/http"

	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type deleteCommentReq struct {
	cid int `json:"cid" form:"cid" binding: "required,cid"`
}

func (h *Handler) DeleteComment(c *gin.Context) {
	var req deleteCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	err := h.SubjectService.DeleteComment(c, uint(req.cid))
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(req)

	c.IndentedJSON(http.StatusOK, res)
}
