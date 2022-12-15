package eclass

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	m "gag.com/model"
	"github.com/PuerkitoBio/goquery"
)

func (e *Eclass) GetSubjects(ctx context.Context) ([]m.Subject, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://eclass.tukorea.ac.kr/ilos/main/main_form.acl", nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range e.cookies {
		req.AddCookie(cookie)
	}
	res, err := client.Do(req)
	subjects := []m.Subject{}

	if err != nil {
		return subjects, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return subjects, errors.New("status code error: " + res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return subjects, err
	}
	fmt.Println(doc.Find("#contentsIndex > div.index-leftarea02 > div:nth-child(2) > ol > li").Text())

	doc.Find("#contentsIndex > div.index-leftarea02 > div:nth-child(2) > ol > li").Each(func(i int, s *goquery.Selection) {

		startAndRoom := s.Find("span").Text()
		splitPoint := strings.Index(startAndRoom, "(")
		endPoint := strings.Index(startAndRoom, ")")
		if i > 1 && splitPoint != -1 {

			tmpSubject := m.Subject{
				ID:        s.Find("em").AttrOr("kj", " "),
				Name:      strings.ReplaceAll(strings.Trim(s.Find("em").Text(), (" \n\t")), ("\n                  "), " "),
				StartTime: strings.Trim(startAndRoom[:splitPoint], (" \n\t")),
				Room:      startAndRoom[splitPoint+1 : endPoint],
			}
			subjects = append(subjects, tmpSubject)
		}
	})

	return subjects, err
}
