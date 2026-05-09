package models

import "time"

// IPqcCheckFlow 检验流程
type IPqcCheckFlow struct {
	ObjectRrn   int64      `gorm:"column:OBJECT_RRN;primaryKey;type:number(19,0)"` // 主键
	OrgRrn      int64      `gorm:"column:ORG_RRN;type:number(19,0)"`               // 区域主键（指向AD_ORG.OBJECT_RRN）
	IsActive    string     `gorm:"column:IS_ACTIVE;type:varchar2(1);default:Y"`     // 是否逻辑有效（Y-有效，N-逻辑已删除）
	Created     *time.Time `gorm:"column:CREATED;type:date;autoCreateTime"`         // 创建时间
	CreatedBy   string     `gorm:"column:CREATED_BY;type:varchar2(32)"`             // 创建人工号
	Updated     *time.Time `gorm:"column:UPDATED;type:date;autoUpdateTime"`         // 更新时间
	UpdatedBy   string     `gorm:"column:UPDATED_BY;type:varchar2(32)"`             // 更新人工号
	LockVersion int64      `gorm:"column:LOCK_VERSION;type:number(19,0)"`           // 锁版本号
	Code        string     `gorm:"column:CODE;type:varchar2(32);uniqueIndex:IPQC_CHECK_FLOW_IDX1"` // 检验流程代码
	Name        string     `gorm:"column:NAME;type:varchar2(64)"`                   // 检验流程名称
	State       string     `gorm:"column:STATE;type:varchar2(10);default:Active"`   // 状态（Active-生效中，Invalid-已失效）
	FlowType    int8       `gorm:"column:FLOW_TYPE;type:number(3,0)"`               // 流程类别（10生产参数类20外观检验类30测量类）
}

// TableName 指定表名，防止 GORM 默认加 s 变成 ipqc_check_flows
func (IPqcCheckFlow) TableName() string {
	return "IPQC_CHECK_FLOW"
}
