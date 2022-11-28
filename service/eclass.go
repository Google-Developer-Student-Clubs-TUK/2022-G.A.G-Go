package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	model "runner/model"
	"strings"
)

type Eclass struct {
}

func (eclass *Eclass) Login(body model.EclassLoginBody) model.ApiResponse[bool] {
	// struct to json
	pbytes, _ := json.Marshal(body)
	buff := bytes.NewBuffer(pbytes)

	// post request
	res, err := http.Post("https://eclass.tukorea.ac.kr/ilos/lo/login.acl", "application/json", buff)

	// err
	if err != nil {
		log.Fatal(err)
		return model.ApiResponse[bool]{
			ResultCode: 400,
			Msg:        err.Error(),
			Result:     false,
		}
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return model.ApiResponse[bool]{
			ResultCode: 400,
			Msg:        err.Error(),
			Result:     false,
		}
	}
	responseString := string(responseBytes)

	// result err
	if !strings.Contains(responseString, `document.location.href="https://eclass.tukorea.ac.kr/ilos/main/member/login_form.acl"`) {
		log.Fatal("잘못된 접근")
		return model.ApiResponse[bool]{
			ResultCode: 400,
			Msg:        err.Error(),
			Result:     false,
		}
	}
	if strings.Contains(responseString, "로그인 정보가 일치하지 않습니다.") {
		log.Fatal("로그인 정보가 일치하지 않습니다.")
		return model.ApiResponse[bool]{
			ResultCode: 400,
			Msg:        err.Error(),
			Result:     false,
		}
	}

	return model.ApiResponse[bool]{
		ResultCode: 200,
		Msg:        "로그인 성공",
		Result:     true,
	}
}
