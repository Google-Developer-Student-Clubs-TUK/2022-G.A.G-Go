package main

import (
	"net/http"

	docs "gag.com/v2/docs"
	model "gag.com/v2/model"
	eclassService "gag.com/v2/service/eclass"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title gag.com API
// @description api for a gag.com
// @version 1.0
// @BasePath /eclass/v1

func main() {
	docs.SwaggerInfo.BasePath = "/eclass/v1"
	docs.SwaggerInfo.Title = "gag.com API"
	router := gin.Default()
	eclassGroup := router.Group("/eclass/v1")
	{
		eclassGroup.POST("/login", doLogin)
		eclassGroup.GET("/timetable", getTimetable)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}

func doLogin(c *gin.Context) {
	var loginBody model.Login
	// login to json
	if err := c.BindJSON(&loginBody); err != nil {
		return
	}

	cookies := c.Request.Cookies()

	// login to eclassLogin
	eclassLoginBody := model.EclassLoginBody{Usr_id: loginBody.Id, Usr_pwd: loginBody.Pwd}
	eclassServcie := eclassService.Instance(cookies)
	loginResponse := eclassServcie.Login(eclassLoginBody)
	c.IndentedJSON(http.StatusCreated, loginResponse)
}

func getTimetableService() model.ApiResponse[model.Timetable] {
	return model.ApiResponse[model.Timetable]{
		Code:   0,
		Msg:    "success",
		Result: model.Timetable{},
	}
}

// @Summary Get timetable
// @Description Get timetable
// @Tags eclass
// @Param  data body model.Login true "login info"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ApiResponse[model.Timetable]
// @Router /eclass/timetable [get]
// @Failure 400
func getTimetable(c *gin.Context) {
	newTimetable := getTimetableService()
	c.IndentedJSON(http.StatusOK, newTimetable)
}
