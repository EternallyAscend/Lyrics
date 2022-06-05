package project

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour        int64
	Minute      int64
	Second      int64
	MicroSecond int64
}

func GenerateTime(hour int64, minute int64, second int64, micro int64) *Time {
	return &Time{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		MicroSecond: micro,
	}
}

func ReadTimeFromString(time string) (*Time, error) {
	result := strings.Split(time, TimeSep)
	if 4 != len(result) {
		return nil, errors.New("Wrong time: " + time)
	}
	hour, err := strconv.ParseInt(result[0], 10, 64)
	if nil != err {
		return nil, err
	}
	minute, err := strconv.ParseInt(result[1], 10, 64)
	if nil != err {
		return nil, err
	}
	second, err := strconv.ParseInt(result[2], 10, 64)
	if nil != err {
		return nil, err
	}
	micro, err := strconv.ParseInt(result[3], 10, 64)
	if nil != err {
		return nil, err
	}
	return &Time{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		MicroSecond: micro,
	}, nil
}

func (that *Time) ExportString() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		that.HourString(),
		TimeSep,
		that.MinuteString(),
		TimeSep,
		that.SecondString(),
		TimeSep,
		strconv.FormatInt(that.MicroSecond, 10))
}

func (that *Time) DisplayString() string {
	return fmt.Sprintf("%s:%s:%s.%s",
		that.HourString(),
		that.MinuteString(),
		that.SecondString(),
		that.MicroHundredString())
}

func (that *Time) TransferTimeToTimestamp() int64 {
	return ((that.Hour*60+that.Minute)*60+that.Second)*1000 + that.MicroSecond
}

func (that *Time) HourString() string {
	base := ""
	if that.Hour < 10 {
		base = "0"
	}
	base += strconv.FormatInt(that.Hour, 10)
	return base
}

func (that *Time) MinuteWithHourString() string {
	base := ""
	if that.Hour*60+that.Minute < 100 {
		base = "0"
	}
	base += strconv.FormatInt(that.Hour*60+that.Minute, 10)
	return base
}

func (that *Time) MinuteString() string {
	base := ""
	if that.Minute < 10 {
		base = "0"
	}
	base += strconv.FormatInt(that.Minute, 10)
	return base
}

func (that *Time) SecondString() string {
	base := ""
	if that.Second < 10 {
		base = "0"
	}
	base += strconv.FormatInt(that.Second, 10)
	return base
}

func (that *Time) MicroHundredString() string {
	base := ""
	if that.MicroSecond < 100 {
		base = "0"
		if that.MicroSecond < 10 {
			base = "00"
		}
	}
	base += strconv.FormatInt(that.MicroSecond, 10)
	return base
}

func (that *Time) MicroTwoDecimalString() string {
	base := ""
	if that.MicroSecond < 100 {
		base = "0"
	}
	base += strconv.FormatInt(that.MicroSecond/10, 10)
	return base
}

func (that *Time) ExportSRT() string {
	return that.HourString() + ":" + that.MinuteString() + ":" + that.SecondString() + "," + that.MicroHundredString()
}

func (that *Time) ExportLRC() string {
	return "[" + that.MinuteWithHourString() + ":" + that.SecondString() + "." + that.MicroTwoDecimalString() + "]"
}

func TransferTimestampToTime(timestamp int64) *Time {
	hour := timestamp / 1000 / 60 / 60
	minute := timestamp/1000/60 - hour*60
	second := (timestamp/1000 - minute*60) - hour*60*60
	micro := timestamp % 1000
	return &Time{
		Hour:        hour,
		Minute:      minute,
		Second:      second,
		MicroSecond: micro,
	}
}

func TransferMillSecondToString(ms int64) string {
	time := TransferTimestampToTime(ms)
	return time.DisplayString()
}
