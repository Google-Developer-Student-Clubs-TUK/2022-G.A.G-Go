package handler

import (
	"log"
	"net/http"
	"strconv"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type getCommentsReq struct {
	pid string `json:"pid" form:"pid"`
}

type getCommentsRes struct {
	pid      string          `json:"pid" form:"pid"`
	Comments []model.Comment `json:"comments" form:"comments"`
}

func (h *Handler) GetComments(c *gin.Context) {
	var req getCommentsReq
	if ok := bindData(c, &req); !ok {
		return
	}

	pid, err := strconv.ParseUint(req.pid, 10, 32)
	if err != nil {
		return
	}

	comments, err := h.SubjectService.GetComments(c, uint(pid))
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	var commentRes getCommentsRes
	commentRes.pid = req.pid
	commentRes.Comments = comments

	res := app.NewSuccess(commentRes)

	c.IndentedJSON(http.StatusOK, res)
}
