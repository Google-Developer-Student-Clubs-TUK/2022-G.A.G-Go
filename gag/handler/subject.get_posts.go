package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type getPostsReq struct {
	sid        string           `json:"sid"`
	pagination model.Pagination `json:"pagination"`
}

type getPostsRes struct {
	Posts []model.Post `json:"posts"`
}

func (h *Handler) GetPosts(c *gin.Context) {
	var req getPostsReq
	if ok := bindData(c, &req); !ok {
		return
	}

	if req.pagination.PerPage == 0 {
		req.pagination.PerPage = 10
	}

	posts, err := h.SubjectService.GetPosts(c, req.sid, req.pagination)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccessPagination(posts, req.pagination)

	c.IndentedJSON(http.StatusOK, res)
}
