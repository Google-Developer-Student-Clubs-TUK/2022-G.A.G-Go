package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type registerCommentReq struct {
	pid     int    `json:"pid" form:"pid" binding: "required,pid"`
	writer  int    `json:"writer" form:"writer" binding: "required,writer"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type registerCommentRes struct {
	cid int `json:"cid" form:"cid" binding: "required,cid"`
}

func (h *Handler) RegisterComment(c *gin.Context) {
	var req registerCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	comment := &model.Comment{
		Writer:  req.writer,
		Content: req.content,
		PID:     uint(req.pid),
	}

	err := h.SubjectService.RegisterComment(c, comment)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(registerCommentRes{
		cid: int(comment.ID),
	})

	c.IndentedJSON(http.StatusOK, res)
}

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

type editCommentReq struct {
	cid     int    `json:"cid" form:"cid" binding: "required,cid"`
	content string `json:"content" form:"content" binding: "required,content"`
}

type eidtCommentRes struct {
	cid int `json:"cid" form:"cid"`
}

func (h *Handler) EditComment(c *gin.Context) {
	var req editCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	comment := &model.Comment{
		Content: req.content,
	}
	comment.ID = uint(req.cid)

	err := h.SubjectService.EditComment(c, comment)
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(eidtCommentRes{
		cid: req.cid,
	})

	c.IndentedJSON(http.StatusOK, res)
}

type deleteCommentReq struct {
	cid int `json:"cid" form:"cid" binding: "required,cid"`
}

func (h *Handler) DeleteComment(c *gin.Context) {
	var req deleteCommentReq
	if ok := bindData(c, &req); !ok {
		return
	}

	err := h.SubjectService.DeleteComment(c, uint(req.cid))
	if err != nil {
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(req)

	c.IndentedJSON(http.StatusOK, res)
}
