package util

import (
	"fmt"
	"time"
)

func GetCalendarWeek() (string, error) {
	t := time.Now()
	yr, cw := t.ISOWeek()

	fmt.Println(yr, cw)

	return "", nil
}
