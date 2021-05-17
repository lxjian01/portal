package utils

import (
	"time"
)

const myTimeFormat string = "2006-01-02 15:04:05"

//MyTime 自定义时间

type MyTime time.Time

func (mt MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(myTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(mt).AppendFormat(b, myTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (mt *MyTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+myTimeFormat+`"`, string(b), time.Local)
	*mt = MyTime(now)
	return err
}

func (mt MyTime) String() string {
	return time.Time(mt).Format(myTimeFormat)
}

func (mt MyTime)Now() MyTime {
	return MyTime(time.Now())
}

func (mt MyTime)ParseTime(t time.Time) MyTime {
	return MyTime(t)
}

func (mt MyTime) format() string {
	return time.Time(mt).Format(myTimeFormat)
}