package handler

import (
	"log"
	"net/http"

	"gag.com/model"
	"gag.com/model/app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TestLogin(c *gin.Context) {
	var req loginReq
	if ok := bindData(c, &req); !ok {
		return
	}

	user := &model.User{
		UUID:        req.UUID,
		ID:          req.ID,
		AesPassword: req.Password,
	}

	err := h.UserService.TestLogin(c, req.Key, user)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(app.Status(err), gin.H{
			"error": err,
		})
		return
	}

	res := app.NewSuccess(loginRes{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		ImageURL: user.ImageURL,
		Major:    "컴퓨터공학과",
	})

	c.IndentedJSON(http.StatusOK, res)
}
