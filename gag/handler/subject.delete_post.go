package handler

import (
	"net/http"
	"strconv"

	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type deletePostsReq struct {
	ID string `json:"id" form:"id" binding: "required,id"`
}

func (h *Handler) DeletePost(c *gin.Context) {
	var req deletePostsReq
	if ok := bindData(c, &req); !ok {
		return
	}

	pid, err := strconv.ParseUint(req.ID, 10, 32)
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
