package utils

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

//MyTime 自定义时间

type MyTime time.Time

func (mt *MyTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*mt = MyTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*mt = MyTime(now)
	return
}

func (mt MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(mt).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (mt MyTime) Value() (driver.Value, error) {
	if mt.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(mt).Format(TimeFormat)), nil
}

func (mt *MyTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*mt = MyTime(tTime)
	return nil
}

func (mt MyTime) String() string {
	return time.Time(mt).Format(TimeFormat)
}

func (mt MyTime) Now() MyTime {
	return MyTime(time.Now())
}

func (mt MyTime) ParseTime(t time.Time) MyTime {
	return MyTime(t)
}

func (mt MyTime) format() string {
	return time.Time(mt).Format(TimeFormat)
}