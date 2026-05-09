package models

import (
	"time"
)

// IPQC_CHECK_TASK 检验任务表
type IPqcCheckTask struct {
	ObjectRrn        int64     `gorm:"column:OBJECT_RRN;primaryKey"`        // 主键
	OrgRrn           int64     `gorm:"column:ORG_RRN"`                      // 区域主键（指向AD_ORG.OBJECT_RRN）
	IsActive         string    `gorm:"column:IS_ACTIVE"`                    // 是否逻辑有效（Y-有效，N-逻辑已删除）
	Created          time.Time `gorm:"column:CREATED"`                      // 创建时间
	CreatedBy        string    `gorm:"column:CREATED_BY"`                   // 创建人工号
	Updated          time.Time `gorm:"column:UPDATED"`                      // 更新时间
	UpdatedBy        string    `gorm:"column:UPDATED_BY"`                   // 更新人工号
	LockVersion      int64     `gorm:"column:LOCK_VERSION"`                 // 锁版本号
	CheckCode        string    `gorm:"column:CHECK_CODE"`                   // 检验单号
	LotId            string    `gorm:"column:LOT_ID"`                       // 批次号
	SampleIds        string    `gorm:"column:SAMPLE_IDS"`                   // 样本ID(PNLID/SETID)（多个样本以逗号隔开连接显示）
	CheckType        int       `gorm:"column:CHECK_TYPE"`                   // 检验类型（10首件检验20过程检验） 10首件检验  20动态检验  30尾件检验 40来料检验
	CheckReason      string    `gorm:"column:CHECK_REASON"`                 // 检验原因
	CheckTimes       int       `gorm:"column:CHECK_TIMES"`                  // 累计次数
	ProductNumberVer string    `gorm:"column:PRODUCT_NUMBER_VER"`           // 本厂编号,产品型号
	TriggerStep      string    `gorm:"column:TRIGGER_STEP"`                 // 触发工步
	CheckName        string    `gorm:"column:CHECK_NAME"`                   // 检验名称
	TakePosition     string    `gorm:"column:TAKE_POSITION"`                // 取板位置
	WeekCode         int       `gorm:"column:WEEK_CODE"`                    // 周码
	Layers           int       `gorm:"column:LAYERS"`                       // 层数
	LayerLevel       string    `gorm:"column:LAYER_LEVEL"`                  // 层别
	CustomerCode     string    `gorm:"column:CUSTOMER_CODE"`                // 客户代号
	State            string    `gorm:"column:STATE"`                        // 检验状态（10-待检验、20-已退回、30-已作废、40-检验中、50-待审核、60-已完成,5.没有同步样本id），5未到达检验站点
	LotNumber        string    `gorm:"column:LOT_NUMBER"`                   // 检验数量＝批次数量
	CheckNumber      int       `gorm:"column:CHECK_NUMBER"`                 // 抽检数量
	AccNumber        int       `gorm:"column:ACC_NUMBER"`                   // 合格数量
	Unit             string    `gorm:"column:UNIT"`                         // 单位代码
	RejNumber        int       `gorm:"column:REJ_NUMBER"`                   // 不合格数量
	DecisionResult   string    `gorm:"column:DECISION_RESULT"`              // 判定结果
	OriDecisionResult string   `gorm:"column:ORI_DECISION_RESULT"`          // 原始判定结果
	DecisionRemark   string    `gorm:"column:DECISION_REMARK"`              // 判定结果说明
	OriCheckCode     string    `gorm:"column:ORI_CHECK_CODE"`               // 原检验单号
	CustomerModel    string    `gorm:"column:CUSTOMER_MODEL"`               // 客户型号
	PostTime         time.Time `gorm:"column:POST_TIME"`                    // 提交时间
	AuditTime        time.Time `gorm:"column:AUDIT_TIME"`                   // 审核时间
	Auditby          string    `gorm:"column:AUDITBY"`                      // 审核人主键（指向AD_USER.OBJECT_RRN）
	Remark           string    `gorm:"column:REMARK"`                       // 备注
	Stage            string    `gorm:"column:STAGE"`                        // 阶段
	NextAuditName    string    `gorm:"column:NEXT_AUDIT_NAME"`              // 二级审批人
	NextAuditTime    time.Time `gorm:"column:NEXT_AUDIT_TIME"`              // 二级审批时间
	IsRecheck        string    `gorm:"column:IS_RECHECK"`                   // 是否复检(Y/N)
	DefectType       string    `gorm:"column:DEFECT_TYPE"`                  // 缺陷类型
	RecheckRemark    string    `gorm:"column:RECHECK_REMARK"`               // 附件备注
	Attachment       string    `gorm:"column:ATTACHMENT"`                   // 附件
	DefectLevel      string    `gorm:"column:DEFECT_LEVEL"`                 // 缺陷等级
	Source           int       `gorm:"column:SOURCE"`                       // 任务来源，1：OIC创建(手动创建)  2：MES创建(pda创建).3:是首批下的子批
	EquipmentId      string    `gorm:"column:EQUIPMENT_ID"`                 // 设备
	CheckNameRrn     int64     `gorm:"column:CHECK_NAME_RRN"`               // 检验方案主键
	SonLotId         string    `gorm:"column:SON_LOT_ID"`                   // 子批号
	Ipqctype         string    `gorm:"column:IPQCTYPE"`                     // IPQC检验类型（自检/半检）
	Checkwhy         int       `gorm:"column:CHECKWHY"`                     // 触发方式，0：自动；1：手动PDA
	Entermark        string    `gorm:"column:ENTERMARK"`                    // 录入标识
	DeleteRemark     string    `gorm:"column:DELETE_REMARK"`                // 删除备注信息
	NcmCode          string    `gorm:"column:NCM_CODE"`                     // 不合格任务单号
	NcnCode          string    `gorm:"column:NCN_CODE"`                     // NCN单号
	Entryby          string    `gorm:"column:ENTRYBY"`                      // 录入人
}

// TableName 指定 Oracle 表名（必须大写）
func (IPqcCheckTask) TableName() string {
	return "IPQC_CHECK_TASK"
}