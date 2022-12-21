package handler

import (
	"log"
	"net/http"
	"strings"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

type (
	TodayRes struct {
		ID        string `json:"id" gorm:"primaryKey"`
		Name      string `json:"name"`
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
		IsPinned  bool   `json:"isPinned"`
		Room      string `json:"room"`
	}
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
	tmpList := make([]TodayRes, 0)
	for _, subject := range subjects {
		if subject.IsToday() {
			parseIndex := strings.LastIndex(subject.StartTime, "~")
			startTime := subject.StartTime[parseIndex-5 : parseIndex]
			endTime := subject.StartTime[parseIndex+1 : parseIndex+6]
			tmp := &TodayRes{
				ID:        subject.ID,
				Name:      subject.Name,
				StartTime: startTime,
				EndTime:   endTime,
				IsPinned:  subject.IsPinned,
				Room:      subject.Room,
			}

			tmpList = append(tmpList, *tmp)
		}
	}

	res := app.NewSuccess(tmpList)

	c.IndentedJSON(http.StatusOK, res)
}
