package gotime

import "testing"

func TestNew(t *testing.T) {
	gt := Now()
	t.Log(gt.FormatDateString())
	t.Log(gt.FormatDateTimeString())
	t.Log(gt.FormatDateTimeStartString())
	t.Log(gt.FormatDateTimeEndString())
	gt = gt.LastWeek()
	t.Log(gt.FormatDateTimeString())
	gt = gt.LastMonth()
	t.Log(gt.FormatDateTimeString())
	gt = gt.LastYear()
	t.Log(gt.FormatDateTimeString())

	gt = NewWithDateString("2011-12-02")
	t.Log(gt.FormatDateString())
	t.Log(gt.FormatDateTimeString())
	t.Log(gt.FormatDateTimeStartString())
	t.Log(gt.FormatDateTimeEndString())
	gt = gt.NextWeek()
	t.Log(gt.FormatDateTimeString())
	gt = gt.NextMonth()
	t.Log(gt.FormatDateTimeString())
	gt = gt.NextYear()
	t.Log(gt.FormatDateTimeString())
}
