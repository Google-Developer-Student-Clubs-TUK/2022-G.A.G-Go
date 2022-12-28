package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type getCommentsReq struct {
	pid int `json:"pid" form:"pid"`
}

type getCommentsRes struct {
	pid      int             `json:"pid" form:"pid"`
	comments []model.Comment `json:"comments" form:"comments"`
}

func (h *Handler) GetComments(c *gin.Context) {
	var req getCommentsReq
	if ok := bindData(c, &req); !ok {
		return
	}

	comments, err := h.SubjectService.GetComments(c, uint(req.pid))
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	var commentRes getCommentsRes
	commentRes.pid = req.pid
	commentRes.comments = comments

	res := app.NewSuccess(commentRes)

	c.IndentedJSON(http.StatusOK, res)
}
