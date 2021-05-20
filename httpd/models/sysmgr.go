package models

type Menu struct {
	BaseModel
	Pid         int `gorm:"column:pid;type:int;default=0" json:"pid" form:"pid"`
	Title   string `gorm:"column:title;type:varchar(64)" json:"title" form:"title" binding:"required"`
	PTitle   string `json:"ptitle"`
	Path      string `gorm:"column:path;type:varchar(64)" json:"path" form:"path" binding:"required"`
	Icon    string `gorm:"column:icon;type:varchar(64)" json:"icon" form:"icon" binding:""`
	Sort uint8 `gorm:"column:sort;type:tinyint;default=1" json:"sort" form:"sort" binding:""`
}


type LoginUser struct {
	Uid   string      `json:"uid"`
	Cn string   `json:"cn"`
	PhoneNo    string   `json:"phone_no"`
	Email    string   `json:"email"`
}

type SystemUser struct {
	Uid        string `gorm:"primaryKey;column:uid;type:varchar(50)" json:"uid" form:"uid" binding:""`
	Cn         string `gorm:"column:cn;type:varchar(50)" json:"cn" form:"cn" binding:"required"`
	Position   string `gorm:"column:position;type:varchar(2)" json:"position" form:"position" binding:""`
	Email      string `gorm:"column:email;type:varchar(50)" json:"email" form:"email" binding:"required,email"`
	PhoneNo    string `gorm:"column:phone_no;type:varchar(50)" json:"phone_no" form:"phone_no" binding:"required"`
	CreateUser string `gorm:"column:create_user;type:varchar(50);default='system'" json:"create_user"`
	CreateTime MyTime `gorm:"column:create_time;type:timestamp" json:"create_time"`
	UpdateUser string `gorm:"column:update_user;type:varchar(50);default='system'" json:"update_user"`
	UpdateTime MyTime `gorm:"column:update_time;type:timestamp" json:"update_time"`
}

type SystemRole struct {
	Id         int    `gorm:"primaryKey;column:id;type:int" json:"id"`
	RoleCode   string `gorm:"column:role_code;type:varchar(50)" json:"role_code"`
	RoleName   string `gorm:"column:role_name;type:varchar(50)" json:"role_name"`
	CreateUser string `gorm:"column:createUser;type:varchar(50);default='system'" json:"create_user"`
	CreateTime MyTime `gorm:"column:createTime;type:timestamp" json:"create_time"`
	UpdateUser string `gorm:"column:updateUser;type:varchar(50);default='system'" json:"update_user"`
	UpdateTime MyTime `gorm:"column:updateTime;type:timestamp" json:"update_time"`
}

type SystemUserRole struct {
	Id int `gorm:"primaryKey;column:id;type:int;" json:"id"`
	Uid string `gorm:"column:uid;type:varchar(50)" json:"uid"`
	RoleCode string `gorm:"column:role_code;type:varchar(50)" json:"role_code"`
}
