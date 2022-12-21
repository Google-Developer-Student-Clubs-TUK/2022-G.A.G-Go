package handler

import (
	"net/http"
	"strconv"

	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type deleteCommentReq struct {
	cid string `json:"cid" form:"cid" binding: "required,cid"`
}

func (h *Handler) DeleteComment(c *gin.Context) {
	var req deleteCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	cid, err := strconv.ParseUint(req.cid, 10, 32)
	if err != nil {
		return
	}

	err = h.SubjectService.DeleteComment(c, uint(cid))
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(req)

	c.IndentedJSON(http.StatusOK, res)
}
