package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type getPostsReq struct {
	subjectId string       `json:"subject_id"`
	paging    model.Paging `json:"paging"`
}

type getPostsRes struct {
	Posts []model.Post `json:"posts"`
}

func (h *Handler) GetPosts(c *gin.Context) {
	var req getPostsReq
	if ok := bindData(c, &req); !ok {
		return
	}

	if req.paging.PerPage == 0 {
		req.paging.PerPage = 10
	}

	posts, err := h.SubjectService.GetPosts(c, req.subjectId, req.paging)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccessPaging(posts, req.paging)

	c.IndentedJSON(http.StatusOK, res)
}
