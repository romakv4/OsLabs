package main

import (
	"fmt"
	"testing"
	"time"
)

type Calendar struct {
	parsed         time.Time
	CurrentQuarter func() int
}

func NewCalendar(parsed time.Time) *Calendar {
	clndr := Calendar{parsed: parsed}
	clndr.CurrentQuarter = func() int {
		var quarter = 0
		switch clndr.parsed.Month().String() {
		case "January":
			quarter = 1
		case "February":
			quarter = 1
		case "March":
			quarter = 1
		case "April":
			quarter = 2
		case "May":
			quarter = 2
		case "June":
			quarter = 2
		case "July":
			quarter = 3
		case "August":
			quarter = 3
		case "September":
			quarter = 3
		case "October":
			quarter = 4
		case "November":
			quarter = 4
		case "December":
			quarter = 4
		}
		return quarter
	}
	return &clndr
}

func TestCurrentQuarter(t *testing.T) {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1},
		{month: "02", quarter: 1},
		{month: "03", quarter: 1},
		{month: "04", quarter: 2},
		{month: "05", quarter: 2},
		{month: "06", quarter: 2},
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}

	//TODO Реализовать Календарь

	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		calendar := NewCalendar(parsed)
		actual := calendar.CurrentQuarter()
		if actual != test.quarter {
			t.Error("Month:", test.month,
				"Expected Quarter:", test.quarter,
				"Actual Quarter:", actual)
		}
	}
}
