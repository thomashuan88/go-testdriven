package date_time

import (
	"fmt"
	"testing"
	"time"
)

func TestYearMonthDay(t *testing.T) {
	year, month, day := time.Now().Date()
	t.Log(year, month, day)
}

func TestMonthWithZero(t *testing.T) {
	tm := time.Now()
	tm = tm.AddDate(0, 2, 0) // add year, month, day
	// date := fmt.Sprintf("%s-%d.txt", tm.Format("2006-01-02"), tm.Hour())
	date := fmt.Sprintf("%d-%d-%d-%d", tm.Year(), tm.Month(), tm.Day(), tm.Hour())

	// date := fmt.Sprintf("%s.txt", tm.Format("2006-01"))
	t.Log(date)
}
