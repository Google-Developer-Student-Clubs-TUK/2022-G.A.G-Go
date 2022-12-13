package eclass

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"gag.com/eclass/model"
	"gag.com/util"
)

func (e *Eclass) Login(ctx context.Context, body *model.LoginBody) error {
	// struct to formdata
	ct, formData, err := util.StructToForm(&body)
	if err != nil {
		return err
	}

	// request
	res, err := http.Post("https://eclass.tukorea.ac.kr/ilos/lo/login.acl", ct, formData)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("status code error: " + res.Status)
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 성공
	responseString := string(responseBytes)
	if strings.Contains(responseString, `document.location.href="https://eclass.tukorea.ac.kr/ilos/main/main_form.acl"`) {

		// set cookie
		e.cookies = res.Cookies()
		return nil
	}

	// 실패
	if strings.Contains(responseString, "로그인 정보가 일치하지 않습니다.") {
		return errors.New("아이디 또는 비밀번호가 잘못되었습니다")
	}

	return errors.New("알 수 없는 에러")
}
