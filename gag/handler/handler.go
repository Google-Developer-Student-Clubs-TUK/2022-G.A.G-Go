package handler

import (
	"gag.com/model"
	"github.com/gin-gonic/gin"
)

// handler layer 내 service 정의
type Handler struct {
	UserService model.UserService
}

// 의존성이 주입되며 handler 레이어 초기설정
type Config struct {
	R           *gin.Engine
	UserService model.UserService
}

func NewHandler(c *Config) {
	// 의존성 주입
	h := &Handler{
		UserService: c.UserService,
	}

	v1 := c.R.Group("/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/login", h.Login)
			userGroup.POST("/testlogin", h.TestLogin)
			userGroup.POST("/register", h.DeviceRegister)
			userGroup.POST("/subjects", h.GetSubjects)
			userGroup.PUT("/alarm", h.SetAlarm)
			userGroup.PUT("/profile/visibility", h.SetProfileVisibility)
		}
	}
}
