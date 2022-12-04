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

	aesKey, err := rh.DecryptString(login.Key)
	if err != nil {
		apiFailure(c, 1)
	}
	iv := util.PKCS5Padding([]byte(aesKey[0:8]), 16)

	// AES 복호화
	jsonData := util.AESDecrypt([]byte(login.AesData), []byte(aesKey), iv)

	// 로그인
	var loginBody model.LoginBody
	json.Unmarshal(jsonData, &loginBody)
	eclassServcie := eclassService.Instance()
	LoginResponse := eclassServcie.Login(loginBody)
	// 로그인 실패시
	if LoginResponse.Code != 0 {
		apiFailure(c, LoginResponse.Code)
	}

	// 로그인 성공시
	// 성공한 데이터 UserDB에 넣기
	user := model.UserDB{Id: loginBody.Usr_id, PrivateKey: temp.PrivateKey, AesData: login.AesData}
	db.Create(&user)

	// tempDB 삭제
	db.Delete(&model.TempDB{}, login.UUID)

	// 성공 메시지 전송
	apiSuccess(c, model.Result{Success: "true"})
}

// func getTimetable(c *gin.Context) {
// 	newTimetable := getTimetableService()
// 	c.IndentedJSON(http.StatusOK, newTimetable)
// }

func deviceRegister(c *gin.Context) {
	var device model.Device
	c.BindJSON(&device)

	// rha 키 생성
	rh := util.RSAHelper{}
	rh.GenerateKey(1024)

	private_key, _ := rh.PrivateToStringPEM()
	public_key, _ := rh.PublicToStringPEM()

	// db에 rha 키 저장
	temp := model.TempDB{UUID: device.UUID, PrivateKey: private_key, PublicKey: public_key}
	db.Create(&temp)

	apiSuccess(c, model.RSA{PublicKey: public_key})
}

// 상의해보고 사용계획
func apiSuccess[T any](c *gin.Context, result T) {
	resultJson, _ := json.Marshal(result)

	response := model.ApiResponse[string]{
		Code:   0,
		Msg:    "성공",
		Result: string(resultJson),
	}

	c.IndentedJSON(http.StatusOK, response)
}

func apiFailure(c *gin.Context, code int) {
	response := model.ApiResponse[string]{
		Code:   code,
		Msg:    ErrorCodeToMsg(code),
		Result: "false",
	}

	c.IndentedJSON(http.StatusNotFound, response)
}

func ErrorCodeToMsg(code int) string {
	codeMap := map[int]string{
		1: "error1",
		2: "error2",
		3: "error3",
	}

	errorMessage, exists := codeMap[code]
	if !exists {
		return "존재하지 않는 에러코드"
	}

	return errorMessage
}
