package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BaseModel struct {
	Id        int `myorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id" binding:""`
	CreateUser string `myorm:"column:create_user;type:varchar(64);default='system'" json:"createUser"`
	CreateTime MyTime `myorm:"column:create_time;type:datetime;autoCreateTime" json:"createTime"`
	UpdateUser string `myorm:"column:update_user;type:varchar(64);default='system'" json:"updateUser"`
	UpdateTime MyTime `myorm:"column:update_time;type:datetime;autoUpdateTime" json:"updateTime"`
}

const myTimeFormat string = "2006-01-02 15:04:05"

//MyTime 自定义时间

type MyTime struct {
	time.Time
}

func (mt MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", mt.Format(myTimeFormat))
	return []byte(formatted), nil
}

func (mt MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if mt.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return mt.Time, nil
}

func (mt *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*mt = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (mt *MyTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+myTimeFormat+`"`, string(b), time.Local)
	*mt = MyTime{Time: now}
	return err
}