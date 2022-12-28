package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type registerPostReq struct {
	writer  string `json:"writer" form:"writer" binding: "required,writer"`
	sid     string `json:"sid" form:"sid" binding: "required,sid"`
	title   string `json:"title" form:"title" binding: "required,title"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type registerPostRes struct {
	pid int `json:"pid" form:"pid" binding: "required,pid"`
}

func (h *Handler) RegisterPost(c *gin.Context) {
	var req registerPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	post := &model.Post{
		Writer:  req.writer,
		SID:     req.sid,
		Title:   req.title,
		Content: req.content,
	}

	err := h.SubjectService.RegisterPost(c, post)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(registerPostRes{
		pid: int(post.ID),
	})

	c.IndentedJSON(http.StatusOK, res)
}

type getPostsReq struct {
	sid        string           `json:"sid" form": "sid"`
	pagination model.Pagination `json:"pagination" form": "pagination"`
}

type getPostsRes struct {
	sid   string       `json:"sid" form": "sid"`
	posts []model.Post `json:"posts" form": "posts"`
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

	var postRes getPostsRes
	postRes.sid = req.sid
	postRes.posts = posts

	res := app.NewSuccessPagination(postRes, req.pagination)

	c.IndentedJSON(http.StatusOK, res)
}

type editPostReq struct {
	pid     int    `json:"pid" form:"pid" binding: "required,pid"`
	title   string `json:"title" form:"title" binding: "required,title"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type eidtPostRes struct {
	pid int `json:"pid" form:"pid"`
}

func (h *Handler) EditPost(c *gin.Context) {
	var req editPostReq
	if ok := bindData(c, &req); !ok {
		return
	}

	post := &model.Post{
		Title:   req.title,
		Content: req.content,
	}
	post.ID = uint(req.pid)

	err := h.SubjectService.EditPost(c, post)
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(eidtPostRes{
		pid: req.pid,
	})

	c.IndentedJSON(http.StatusOK, res)
}

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
