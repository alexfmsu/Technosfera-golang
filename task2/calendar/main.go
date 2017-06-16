package main

import (
	"time"
)

type Calendar struct {
	month int
}

func NewCalendar(t time.Time) Calendar {
	month := t.Month()

	return Calendar{int(month)}
}

func (c Calendar) CurrentQuarter() int {
	return (c.month-1)/3 + 1
}
