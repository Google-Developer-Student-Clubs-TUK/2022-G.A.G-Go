package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTodaySubjects(c *gin.Context) {
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

	todays := make([]model.Subject, 0)
	for _, subject := range subjects {
		if subject.IsToday() {
			todays = append(todays, subject)
		}
	}

	res := app.NewSuccess(todays)

	c.IndentedJSON(http.StatusOK, res)
}
