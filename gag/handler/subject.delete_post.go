package handler

import (
	"net/http"
	"strconv"

	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type deletePostReq struct {
	pid string `json:"pid" form:"pid" binding: "required,pid"`
}

func (h *Handler) DeletePost(c *gin.Context) {
	var req deletePostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	pid, err := strconv.ParseUint(req.pid, 10, 32)
	if err != nil {
		return
	}

	err = h.SubjectService.DeletePost(c, uint(pid))
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(req)

	c.IndentedJSON(http.StatusOK, res)
}
