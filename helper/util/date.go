package util

import (
	"time"
)

var (
	LayoutDefault  = "2006-01-02 15:04:05"
	LayoutDateOnly = "2006-01-02"
	LocalTimeZone  = "Asia/Ho_Chi_Minh"
	Loc, _         = time.LoadLocation(LocalTimeZone)
)

func TimeNow() time.Time {
	return time.Now().In(Loc)
}

func TimeParse(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func StrToTime(dateTime string) (time.Time, error) {
	return StrToTimeWithLayout(dateTime, LayoutDefault)
}

func StrToTimeWithLayout(dateTime, layout string) (time.Time, error) {
	DateTime, err := time.ParseInLocation(layout, dateTime, Loc)
	if err != nil {
		return time.Time{}, err
	}
	return DateTime, err
}

func StartOfDate(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, Loc)
}

func EndOfDate(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 23, 59, 59, 0, Loc)
}

func StartOfMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, Loc)
}

func EndOfMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month)+1, 0, 23, 59, 59, 0, Loc)
}

func IsBetween(startTime, endTime, checkTime time.Time) bool {
	return checkTime.After(startTime) && checkTime.Before(endTime)
}

func IsAfter(start, end time.Time) bool {
	return start.After(end)
}

func IsBefore(start, end time.Time) bool {
	return start.Before(end)
}
