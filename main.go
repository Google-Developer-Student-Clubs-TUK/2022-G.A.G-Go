package main

import (
	"encoding/json"
	"net/http"

	model "gag.com/v2/model"
	eclassService "gag.com/v2/service/eclass"
	util "gag.com/v2/util"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	eclassGroup := router.Group("/eclass")
	{
		eclassGroup.POST("/login", login)
		eclassGroup.POST("/deviceRegister", deviceRegister)
		// eclassGroup.GET("/timetable", getTimetable)
	}
	router.Run("localhost:8080")
}

func login(c *gin.Context) {
	var login model.Login
	c.BindJSON(&login)

	// uuid를 이용해 tempDB 가져오기
	var temp model.TempDB
	db.First(&temp, "uuid = ?", login.UUID)

	// RSA 복호화
	rh := util.RSAHelper{}
	rh.PrivateFromStringPEM(temp.PrivateKey)

	aesData, err := rh.DecryptString(temp.UUID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, model.ApiResponse[bool]{
			Code:   -1,
			Msg:    "rsa 복호화 오류",
			Result: false,
		})
	}

	// AES 복호화
	jsonData := util.AESDecrypt([]byte(aesData), []byte(login.AES), []byte(login.IV))

	// 로그인
	var loginBody model.LoginBody
	json.Unmarshal(jsonData, &loginBody)
	eclassServcie := eclassService.Instance()
	LoginResponse := eclassServcie.Login(loginBody)
	// 로그인 실패시
	if LoginResponse.Code != 0 {
		c.IndentedJSON(http.StatusNotFound, LoginResponse)
	}

	// 로그인 성공시
	// 성공한 데이터 UserDB에 넣기
	user := model.UserDB{StudentId: loginBody.Usr_id, PrivateKey: temp.PrivateKey, AES: login.AES}
	db.Create(&user)

	// tempDB 삭제
	db.Delete(&model.TempDB{}, login.UUID)

	// 성공 메시지 전송
	c.IndentedJSON(http.StatusCreated, model.ApiResponse[bool]{
		Code:   0,
		Msg:    "성공",
		Result: true,
	})
}

// func getTimetable(c *gin.Context) {
// 	newTimetable := getTimetableService()
// 	c.IndentedJSON(http.StatusOK, newTimetable)
// }

func deviceRegister(c *gin.Context) {
	uuid := c.PostForm("uuid")

	// rha 키 생성
	rh := util.RSAHelper{}
	rh.GenerateKey(1024)

	private_key, _ := rh.PrivateToStringPEM()
	public_key, _ := rh.PublicToStringPEM()

	// db에 rha 키 저장
	temp := model.TempDB{UUID: uuid, PrivateKey: private_key, PublicKey: public_key}
	db.Create(&temp)

	response := model.ApiResponse[string]{
		Code:   0,
		Msg:    "디바이스 등록 성공",
		Result: public_key,
	}

	c.IndentedJSON(http.StatusCreated, response)
}

// // 상의해보고 사용계획
// func apiSuccess[T any](c *gin.Context, result T) {
// 	c.IndentedJSON(http.StatusCreated, model.ApiResponse[T]{
// 		Code:   0,
// 		Msg:    "성공",
// 		Result: result,
// 	})
// }

// func apiFailure(c *gin.Context, code int, msg string) {
// 	c.IndentedJSON(http.StatusCreated, model.ApiResponse[bool]{
// 		Code:   code,
// 		Msg:    msg,
// 		Result: false,
// 	})
// }
// failure result
// 1. 백엔드에서 타입에 맞게 임의값 넣기
// 2. 프론트에서 알아서 처리
