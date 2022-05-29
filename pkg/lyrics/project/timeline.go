package project

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type TimeFlag struct {
	Start *Time
	End   *Time
}

func GenerateTimeFlag(start *Time, end *Time) *TimeFlag {
	return &TimeFlag{
		Start: start,
		End:   end,
	}
}

func ReadLineFromString(line string) (*TimeFlag, error) {
	result := strings.Split(line, TimeFlagSep)
	if 2 != len(result) {
		return nil, errors.New("Wrong line: " + line)
	}
	start, err := ReadTimeFromString(result[0])
	if nil != err {
		return nil, err
	}
	end, err := ReadTimeFromString(result[1])
	return &TimeFlag{
		Start: start,
		End:   end,
	}, err
}

func (that *TimeFlag) ExportString() string {
	return fmt.Sprintf("%s%s%s", that.Start.ExportString(), TimeFlagSep, that.End.ExportString())
}

type Timeline struct {
	Offset int64
	Lines  *[]*TimeFlag
}

func GenerateTimeline(offset int64) *Timeline {
	return &Timeline{
		Offset: offset,
		Lines:  &[]*TimeFlag{},
	}
}

func (that *Timeline) DealTimeOffset(time *Time) *Time {
	return TransferTimestampToTime(time.TransferTimeToTimestamp() + that.Offset)
}

func (that *Timeline) ReadTimelineFromFile(path string) error {
	file, err := os.Open(path)
	defer func(file *os.File) {
		errIn := file.Close()
		if errIn != nil {
			log.Println(err)
			return
		}
	}(file)
	if nil != err {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		timeFlag, errIn := ReadLineFromString(scanner.Text())
		if nil != errIn {
			return errIn
		}
		that.Append(timeFlag)
	}
	return nil
}

func (that *Timeline) Append(flag *TimeFlag) {
	*that.Lines = append(*that.Lines, flag)
}

func (that *Timeline) Insert(position int, flag *TimeFlag) {
	front := (*that.Lines)[0:position]
	front = append(front, flag)
	*that.Lines = append(front, (*that.Lines)[position:]...)
}

func (that *Timeline) ExportString() string {
	result := ""
	for line := range *that.Lines {
		result += (*that.Lines)[line].ExportString() + "\n"
	}
	return result
}

func (that *Timeline) ExportFile(path string, title string) error {
	file, err := os.OpenFile(path+title+"_timeline.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if nil != err {
		return err
	}
	defer func(file *os.File) {
		errIn := file.Close()
		if errIn != nil {
			log.Println(err)
		}
	}(file)
	_, err = file.WriteString(that.ExportString())
	return err
}
