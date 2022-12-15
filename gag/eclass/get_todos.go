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

func (e *Eclass) GetTodos(ctx context.Context) ([]m.Todo, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://eclass.tukorea.ac.kr/ilos/mp/todo_list.acl", nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range e.cookies {
		req.AddCookie(cookie)
	}
	res, err := client.Do(req)
	todos := []m.Todo{}

	if err != nil {
		return todos, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return todos, errors.New("status code error: " + res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return todos, err
	}
	fmt.Println(doc.Find("body > div").Text())

	doc.Find("body > div").Each(func(i int, s *goquery.Selection) {
		tmpName := s.Find("div.todo_title").Text()
		fmt.Println(tmpName)
		if strings.Contains(tmpName, "과제") || strings.Contains(tmpName, "팀프로젝트") {
			tmpStr := s.AttrOr("onclick", " ")
			fmt.Println("tmpstr" + tmpStr)
			splitStart := strings.Index(tmpStr, "(")
			splitEnd := strings.Index(tmpStr, ",")
			tmpDeadline := s.Find("div.todo_date> span.todo_date").Text()
			deadLineSplit := strings.IndexAny(tmpDeadline, "1234567890")
			deadLine := tmpDeadline[deadLineSplit:]
			if splitStart != -1 && splitEnd != -1 {
				tmpTodo := m.Todo{
					ID:       tmpStr[splitStart+1 : splitEnd],
					Name:     strings.ReplaceAll(strings.Trim(tmpName, (" \n\t")), ("\n                  "), " "),
					DeadLine: strings.ReplaceAll(strings.Trim(deadLine, (" \n\t")), ("\n                  "), " "),
					IsDone:   false,
				}
				todos = append(todos, tmpTodo)
			}

		}
	})
	for _, todo := range todos {
		fmt.Println(todo)
	}
	return todos, err
}
