package handler

import (
	"gag.com/model"
	"github.com/gin-gonic/gin"
)

// handler layer 내 service 정의
type Handler struct {
	UserService    model.UserService
	SubjectService model.SubjectService
}

// 의존성이 주입되며 handler 레이어 초기설정
type Config struct {
	R              *gin.Engine
	UserService    model.UserService
	SubjectService model.SubjectService
}

func NewHandler(c *Config) {
	// 의존성 주입
	h := &Handler{
		UserService:    c.UserService,
		SubjectService: c.SubjectService,
	}

	v1 := c.R.Group("/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/login", h.Login)
			userGroup.POST("/testlogin", h.TestLogin)
			userGroup.POST("/register", h.DeviceRegister)
			userGroup.POST("/subjects", h.GetSubjects)
			userGroup.POST("/subjects/today", h.GetTodaySubjects)
			userGroup.PUT("/profile/alarm", h.SetAlarm)
			userGroup.PUT("/profile/visibility", h.SetProfileVisibility)
			userGroup.POST("/todos", h.GetTodos)
			userGroup.POST("/profile", h.GetProfile)
			userGroup.POST("/token", h.RefreshToken)
		}
		subjectGroup := v1.Group("/subjects")
		{
			subjectGroup.GET("/posts", h.GetPosts)
			subjectGroup.POST("/posts", h.RegisterPost)
			subjectGroup.PUT("/posts", h.EditPost)
			subjectGroup.DELETE("/posts", h.DeletePost)
			subjectGroup.GET("/comments", h.GetComments)
			subjectGroup.POST("/comments", h.RegisterComment)
			subjectGroup.PUT("/comments", h.EditComment)
			subjectGroup.DELETE("/comments", h.DeleteComment)
		}
	}
}
