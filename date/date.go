package date

import (
	"fmt"
	"time"
)

// Date ...
type Date struct {
	datetime interface{}
}

// DateHandler ...
type DateHandler func(d *Date)

// DateFormat ...
const DateFormat = "2006-01-02"

// DateTimeFormat ...
const DateTimeFormat = "2006-01-02 15:04:05"

// NewDate init Date struct
// if len(opts)=0, the base time is time.Now()
func NewDate(opts ...DateHandler) *Date {
	var d = &Date{}
	for _, o := range opts {
		o(d)
	}
	return d
}

// BindTime bind time.Time type as base time
func BindTime(dateObj time.Time) DateHandler {
	return func(d *Date) {
		d.datetime = dateObj
	}
}

// BindDate bind DateFormat type as base time
func BindDate(dateStr string) DateHandler {
	return func(d *Date) {
		d.datetime, _ = time.Parse(DateFormat, dateStr)
	}
}

// BindDateTime bind DateTimeFormat type as base time
func BindDateTime(datetimeStr string) DateHandler {
	return func(d *Date) {
		d.datetime, _ = time.Parse(DateTimeFormat, datetimeStr)
	}
}

// TodayDate ...
func (d *Date) TodayDate() string {
	return d.now().Format(DateFormat)
}

// TodayDateTime ...
func (d *Date) TodayDateTime() string {
	return d.now().Format(DateTimeFormat)
}

// TodayStartDateTime ...
func (d *Date) TodayStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.TodayDate())
}

// TodayEndDateTime ...
func (d *Date) TodayEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.TodayDate())
}

// YesterdayDate ...
func (d *Date) YesterdayDate() string {
	return d.now().AddDate(0, 0, -1).Format(DateFormat)
}

// YesterdayStartDateTime ...
func (d *Date) YesterdayStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.YesterdayDate())
}

// YesterdayEndDateTime ...
func (d *Date) YesterdayEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.YesterdayDate())
}

// TomorrowDate ...
func (d *Date) TomorrowDate() string {
	return d.now().AddDate(0, 0, 1).Format(DateFormat)
}

// TomorrowStartDateTime ...
func (d *Date) TomorrowStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.TomorrowDate())
}

// TomorrowEndDateTime ...
func (d *Date) TomorrowEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.TomorrowDate())
}

// WeekStartDate ...
func (d *Date) WeekStartDate() string {
	return d.weekStart().Format(DateFormat)
}

// LastWeekStartDate ...
func (d *Date) LastWeekStartDate() string {
	return d.weekStart().AddDate(0, 0, -7).Format(DateFormat)
}

// LastWeekStartDateTime ...
func (d *Date) LastWeekStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.LastWeekStartDate())
}

// LastWeekEndDate ...
func (d *Date) LastWeekEndDate() string {
	return d.weekStart().AddDate(0, 0, -1).Format(DateFormat)
}

// LastWeekEndDateTime ...
func (d *Date) LastWeekEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.LastWeekEndDate())
}

// MonthStartDate ...
func (d *Date) MonthStartDate() string {
	return d.monthStart().Format(DateFormat)
}

// MonthStartDateTime ...
func (d *Date) MonthStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.MonthStartDate())
}

// LastMonthStartDate ...
func (d *Date) LastMonthStartDate() string {
	return d.monthStart().AddDate(0, -1, 0).Format(DateFormat)
}

// LastMonthStartDateTime ...
func (d *Date) LastMonthStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.LastMonthStartDate())
}

// LastMonthEndDate ...
func (d *Date) LastMonthEndDate() string {
	return d.monthStart().AddDate(0, 0, -1).Format(DateFormat)
}

// LastMonthEndDateTime ...
func (d *Date) LastMonthEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.LastMonthEndDate())
}

// YearStartDate ...
func (d *Date) YearStartDate() string {
	y, _, _ := d.date()
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.Local).Format(DateFormat)
}

// YearStartDateTime ...
func (d *Date) YearStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.YearStartDate())
}

// LastYearStartDate ...
func (d *Date) LastYearStartDate() string {
	t := d.now().AddDate(-1, 0, 0).Format(DateFormat)
	return NewDate(BindDate(t)).YearStartDate()
}

// LastYearStartDateTime ...
func (d *Date) LastYearStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.LastYearStartDate())
}

// LastYearEndDate ...
func (d *Date) LastYearEndDate() string {
	t, _ := time.Parse(DateFormat, d.YearStartDate())
	return t.AddDate(0, 0, -1).Format(DateFormat)
}

// LastYearEndDateTime ...
func (d *Date) LastYearEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.LastYearEndDate())
}

func (d *Date) weekStart() time.Time {
	return d.now().AddDate(0, 0, -int(d.week())+1)
}

func (d *Date) monthStart() time.Time {
	y, m, _ := d.date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func (d *Date) now() time.Time {
	if d.datetime != nil {
		return (d.datetime).(time.Time)
	}
	return time.Now()
}

func (d *Date) week() time.Weekday {
	// 周计算
	return d.now().Weekday()
}

func (d *Date) date() (year int, month time.Month, day int) {
	// 年月日
	year, month, day = d.now().Date()
	return
}
