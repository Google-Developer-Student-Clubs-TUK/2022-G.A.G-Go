package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type defaultReq struct {
	Key string `json:"key" form:"key" binding: "required,key"`
	ID  string `json:"id" form:"id" binding: "required,id"`
}

type subjectRes struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	IsPinned string `json:"isPinned"`
}

func (h *Handler) GetSubjects(c *gin.Context) {
	var req defaultReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		ID: req.ID,
	}

	subects := make([]model.Subject, 0)

	subjects, err := h.UserService.GetSubjects(c, req.Key, user, subects)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(subjects)

	c.IndentedJSON(http.StatusOK, res)
}
