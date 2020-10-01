package gotime

import (
	"fmt"
	"time"
)

// DateFormat ...
const DateFormat = "2006-01-02"

// DateTimeFormat ...
const DateTimeFormat = "2006-01-02 15:04:05"

// DateTimeStartFormat ...
const DateTimeStartFormat = "2006-01-02 00:00:00"

// DateTimeEndFormat ...
const DateTimeEndFormat = "2006-01-02 23:59:59"

type GoTime struct {
	time.Time
}

func Now() GoTime {
	return GoTime{time.Now()}
}

func New() GoTime {
	return Now()
}

func NewWithTime(t time.Time) GoTime {
	return GoTime{t}
}

// NewWithDateString must be DateFormat(2006-01-02), or will panic
func NewWithDateString(d string) GoTime {
	parse, err := time.Parse(DateFormat, d)
	if err != nil {
		panic(err.Error())
	}
	return GoTime{parse}
}

// NewWithDateString must be DateTimeFormat(2006-01-02 15:04:05), or will panic
func NewWithDateTimeString(dt string) GoTime {
	parse, err := time.Parse(DateTimeFormat, dt)
	if err != nil {
		panic(err.Error())
	}
	return GoTime{parse}
}

// FormatDateString
func (d GoTime) FormatDateString() string {
	return d.Format(DateFormat)
}

// FormatDateTimeString
func (d GoTime) FormatDateTimeString() string {
	return d.Format(DateTimeFormat)
}

// FormatDateTimeStartString
func (d GoTime) FormatDateTimeStartString() string {
	return d.Format(DateTimeStartFormat)
}

// FormatDateTimeEndString
func (d GoTime) FormatDateTimeEndString() string {
	return fmt.Sprintf("%s 23:59:59", d.FormatDateString())
}

// Yesterday ...
func (d GoTime) Yesterday() GoTime {
	d.Time.AddDate(0, 0, -1)
	return d
}

// Tomorrow ...
func (d GoTime) Tomorrow() GoTime {
	d.Time = d.AddDate(0, 0, 1)
	return d
}

// LastWeek ...
func (d GoTime) LastWeek() GoTime {
	d.Time = d.AddDate(0, 0, -7)
	return d
}

// NextWeek ...
func (d GoTime) NextWeek() GoTime {
	d.Time = d.AddDate(0, 0, 7)
	return d
}

// WeekStart is this week of Monday
func (d GoTime) WeekStart() GoTime {
	d.Time = d.AddDate(0, 0, -int(d.Weekday())+1)
	return d
}

// WeekEnd is this week of Sunday
func (d GoTime) WeekEnd() GoTime {
	d.Time = d.WeekStart().AddDate(0, 0, 6)
	return d
}

// LastWeekStart ...
func (d GoTime) LastWeekStart() GoTime {
	d.Time = d.WeekStart().AddDate(0, 0, -7)
	return d
}

// LastWeekEnd ...
func (d GoTime) LastWeekEnd() GoTime {
	d.Time = d.WeekStart().AddDate(0, 0, -1)
	return d
}

// NextWeekStart ...
func (d GoTime) NextWeekStart() GoTime {
	d.Time = d.WeekStart().AddDate(0, 0, 7)
	return d
}

// NextWeekEnd ...
func (d GoTime) NextWeekEnd() GoTime {
	d.Time = d.WeekEnd().AddDate(0, 0, 7)
	return d
}

// LastMonth
func (d GoTime) LastMonth() GoTime {
	d.Time = d.AddDate(0, -1, 0)
	return d
}

// NextMonth
func (d GoTime) NextMonth() GoTime {
	d.Time = d.AddDate(0, 1, 0)
	return d
}

// MonthStart
func (d GoTime) MonthStart() GoTime {
	y, m, _ := d.Date()
	d.Time = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	return d
}

// LastMonthStart
func (d GoTime) LastMonthStart() GoTime {
	y, m, _ := d.AddDate(0, -1, 0).Date()
	d.Time = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	return d
}

// NextMonthStart
func (d GoTime) NextMonthStart() GoTime {
	y, m, _ := d.AddDate(0, 1, 0).Date()
	d.Time = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	return d
}

// NextMonthAfterStart
func (d GoTime) NextMonthAfterStart() GoTime {
	y, m, _ := d.NextMonth().AddDate(0, 1, 0).Date()
	d.Time = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	return d
}

// MonthEnd
func (d GoTime) MonthEnd() GoTime {
	y, m, day := d.NextMonthStart().AddDate(0, 0, -1).Date()
	d.Time = time.Date(y, m, day, 0, 0, 0, 0, time.Local)
	return d
}

// LastMonthEnd
func (d GoTime) LastMonthEnd() GoTime {
	y, m, day := d.MonthStart().AddDate(0, 0, -1).Date()
	d.Time = time.Date(y, m, day, 0, 0, 0, 0, time.Local)
	return d
}

// NextMonthEnd
func (d GoTime) NextMonthEnd() GoTime {
	y, m, day := d.NextMonthAfterStart().AddDate(0, 0, -1).Date()
	d.Time = time.Date(y, m, day, 0, 0, 0, 0, time.Local)
	return d
}

// LastYear
func (d GoTime) LastYear() GoTime {
	d.Time = d.AddDate(-1, 0, 0)
	return d
}

// NextYear
func (d GoTime) NextYear() GoTime {
	d.Time = d.AddDate(1, 0, 0)
	return d
}

// YearStart
func (d GoTime) YearStart() GoTime {
	y, _, _ := d.Date()
	d.Time = time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)
	return d
}

// LastYearStart
func (d GoTime) LastYearStart() GoTime {
	y, _, _ := d.AddDate(-1, 0, 0).Date()
	d.Time = time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)
	return d
}

// NextYearStart
func (d GoTime) NextYearStart() GoTime {
	y, _, _ := d.AddDate(1, 0, 0).Date()
	d.Time = time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)
	return d
}

// NextYearAfterStart
func (d GoTime) NextYearAfterStart() GoTime {
	y, _, _ := d.NextYear().AddDate(1, 0, 0).Date()
	d.Time = time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)
	return d
}

// YearEnd
func (d GoTime) YearEnd() GoTime {
	y, m, day := d.NextYearStart().AddDate(0, 0, -1).Date()
	d.Time = time.Date(y, m, day, 0, 0, 0, 0, time.Local)
	return d
}

// LastYearEnd
func (d GoTime) LastYearEnd() GoTime {
	y, m, day := d.YearStart().AddDate(0, 0, -1).Date()
	d.Time = time.Date(y, m, day, 0, 0, 0, 0, time.Local)
	return d
}

// NextYearEnd
func (d GoTime) NextYearEnd() GoTime {
	y, m, day := d.NextYearAfterStart().AddDate(0, 0, -1).Date()
	d.Time = time.Date(y, m, day, 0, 0, 0, 0, time.Local)
	return d
}
