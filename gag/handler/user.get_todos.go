package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTodos(c *gin.Context) {

	var req defaultReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		ID: req.ID,
	}

	todos := make([]model.Todo, 0)

	todos, err := h.UserService.GetTodos(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(todos)

	c.IndentedJSON(http.StatusOK, res)
}
