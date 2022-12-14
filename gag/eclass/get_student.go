package eclass

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"gag.com/eclass/model"
	"github.com/PuerkitoBio/goquery"
)

func (e *Eclass) GetStudent(ctx context.Context) (*model.Student, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://eclass.tukorea.ac.kr/ilos/mp/myinfo_form.acl", nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range e.cookies {
		req.AddCookie(cookie)
	}
	res, err := client.Do(req)
	student := &model.Student{}

	if err != nil {
		return student, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return student, errors.New("status code error: " + res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return student, err
	}

	doc.Find("#uploadForm > div:nth-child(5) > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			student.Name = strings.Split(s.Find("td:nth-child(2)").Text(), "(")[0]
		}
		if i == 1 {
			student.Phone = s.Find("td:nth-child(2) > div").Text()
		}
		if i == 2 {
			student.Email = s.Find("td:nth-child(2) > div").Text()
		}
	})

	imageUrl, exists := doc.Find("#uploadForm > div:nth-child(5) > table > tbody > tr:nth-child(1) > td:nth-child(3) > div:nth-child(1) > img").Attr("src")
	if !exists {
		return student, errors.New("img exists error")
	}

	student.ImageUrl = imageUrl

	return student, err
}
