package util

import (
	"time"
)

func GetToday() string {
	now := time.Now()
	expected := now.Weekday()
	expectedKorean := expected.String()
	switch expected {
	case time.Monday:
		expectedKorean = "월"
	case time.Tuesday:
		expectedKorean = "화"
	case time.Wednesday:
		expectedKorean = "수"
	case time.Thursday:
		expectedKorean = "목"
	case time.Friday:
		expectedKorean = "금"
	case time.Saturday:
		expectedKorean = "토"
	case time.Sunday:
		expectedKorean = "일"
	}
	return expectedKorean
}
