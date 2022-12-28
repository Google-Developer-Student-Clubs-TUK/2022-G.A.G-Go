package handler

import (
	"net/http"

	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type deletePostReq struct {
	pid int `json:"pid" form:"pid" binding: "required,pid"`
}

func (h *Handler) DeletePost(c *gin.Context) {
	var req deletePostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	err := h.SubjectService.DeletePost(c, uint(req.pid))
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(req)

	c.IndentedJSON(http.StatusOK, res)
}
