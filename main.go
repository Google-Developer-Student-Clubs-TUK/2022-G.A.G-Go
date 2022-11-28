package main

import (
	"net/http"
	model "runner/model"
	service "runner/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	eclassGroup := router.Group("/eclass")
	{
		eclassGroup.POST("/login", doLogin)
		// eclassGroup.GET("/timetable", getTimetable)
	}
	router.Run("localhost:8080")
}

func doLogin(c *gin.Context) {
	var loginBody model.Login

	// login to json
	if err := c.BindJSON(&loginBody); err != nil {
		return
	}

	// login to eclassLogin
	eclassLoginBody := model.EclassLoginBody{Usr_id: loginBody.Id, Usr_pwd: loginBody.Pwd}

	eclassServcie := service.Eclass{}
	resp := eclassServcie.Login(eclassLoginBody)
	c.IndentedJSON(http.StatusCreated, resp)
}

// func getTimetable(c *gin.Context) {
// 	newTimetable := getTimetableService()
// 	c.IndentedJSON(http.StatusOK, newTimetable)
// }
