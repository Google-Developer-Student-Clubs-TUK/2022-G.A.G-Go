package eclass

import (
	"io"
	"log"
	"net/http"
	"strings"

	model "gag.com/v2/model"
	util "gag.com/v2/util"
)

func (eclass *eclass) Login(body model.LoginBody) model.ApiResponse[bool] {
	// struct to formdata
	ct, formData, err := util.StructToForm(body)
	if err != nil {
		log.Fatal(err.Error())
		return model.ApiResponse[bool]{
			Code:   -1,
			Msg:    "created form error",
			Result: false,
		}
	}

	// request
	res, err := http.Post(eclass.domain+"/ilos/lo/login.acl", ct, formData)
	if err != nil {
		log.Fatal(err.Error())
		return model.ApiResponse[bool]{
			Code:   -100,
			Msg:    "서버와의 접속에 실패하였습니다",
			Result: false,
		}
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return model.ApiResponse[bool]{
			Code:   -100,
			Msg:    "서버와의 접속에 실패하였습니다",
			Result: false,
		}
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
		return model.ApiResponse[bool]{
			Code:   -1,
			Msg:    "ReadAll error",
			Result: false,
		}
	}

	// 성공
	responseString := string(responseBytes)
	if strings.Contains(responseString, `document.location.href="https://eclass.tukorea.ac.kr/ilos/main/main_form.acl"`) {

		// set cookie
		eclass.cookies = res.Cookies()
		return model.ApiResponse[bool]{
			Code:   0,
			Msg:    "로그인 성공",
			Result: true,
		}
	}

	// 실패
	if strings.Contains(responseString, "로그인 정보가 일치하지 않습니다.") {
		return model.ApiResponse[bool]{
			Code:   -102,
			Msg:    "아이디 또는 비밀번호를 확인해주세요",
			Result: false,
		}
	}

	return model.ApiResponse[bool]{
		Code:   -109,
		Msg:    "알 수 없는 이유로 로그인에 실패하였습니다",
		Result: false,
	}
}
