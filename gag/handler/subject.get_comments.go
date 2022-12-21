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
	pid        string           `json:"pid"`
	pagination model.Pagination `json:"pagination"`
}

type getCommentsRes struct {
	Comments []model.Comment `json:"comments"`
}

func (h *Handler) GetComments(c *gin.Context) {
	var req getCommentsReq
	if ok := bindData(c, &req); !ok {
		return
	}

	if req.pagination.PerPage == 0 {
		req.pagination.PerPage = 10
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

	res := app.NewSuccessPagination(comments, req.pagination)

	c.IndentedJSON(http.StatusOK, res)
}
