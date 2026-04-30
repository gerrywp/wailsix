package models

type AdUser struct {
	ObjectRrn string `gorm:"column:OBJECT_RRN;primaryKey"` // 主键
	UserName  string `gorm:"column:USER_NAME"`             // 用户名
	Password  string `gorm:"column:PASSWORD"`              // 密码
}

// TableName 强制指定 Oracle 表名（必须大写！）
func (AdUser) TableName() string {
	return "AD_USER"
}