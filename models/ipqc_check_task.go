package models

import (
	"time"
)

// IPQC_CHECK_TASK 检验任务表
type IPqcCheckTask struct {
	ObjectRrn        int64     `gorm:"column:OBJECT_RRN;primaryKey" json:"object_rrn"`        // 主键
	OrgRrn           int64     `gorm:"column:ORG_RRN" json:"org_rrn"`                          // 区域主键（指向AD_ORG.OBJECT_RRN）
	IsActive         string    `gorm:"column:IS_ACTIVE" json:"is_active"`                      // 是否逻辑有效（Y-有效，N-逻辑已删除）
	Created          time.Time `gorm:"column:CREATED" json:"created"`                          // 创建时间
	CreatedBy        string    `gorm:"column:CREATED_BY" json:"created_by"`                    // 创建人工号
	Updated          time.Time `gorm:"column:UPDATED" json:"updated"`                          // 更新时间
	UpdatedBy        string    `gorm:"column:UPDATED_BY" json:"updated_by"`                    // 更新人工号
	LockVersion      int64     `gorm:"column:LOCK_VERSION" json:"lock_version"`                // 锁版本号
	CheckCode        string    `gorm:"column:CHECK_CODE" json:"check_code"`                    // 检验单号
	LotId            string    `gorm:"column:LOT_ID" json:"lot_id"`                            // 批次号
	SampleIds        string    `gorm:"column:SAMPLE_IDS" json:"sample_ids"`                    // 样本ID(PNLID/SETID)（多个样本以逗号隔开连接显示）
	CheckType        int       `gorm:"column:CHECK_TYPE" json:"check_type"`                    // 检验类型（10首件检验20过程检验） 10首件检验  20动态检验  30尾件检验 40来料检验
	CheckReason      string    `gorm:"column:CHECK_REASON" json:"check_reason"`                // 检验原因
	CheckTimes       int       `gorm:"column:CHECK_TIMES" json:"check_times"`                  // 累计次数
	ProductNumberVer string    `gorm:"column:PRODUCT_NUMBER_VER" json:"product_number_ver"`    // 本厂编号,产品型号
	TriggerStep      string    `gorm:"column:TRIGGER_STEP" json:"trigger_step"`                // 触发工步
	CheckName        string    `gorm:"column:CHECK_NAME" json:"check_name"`                    // 检验名称
	TakePosition     string    `gorm:"column:TAKE_POSITION" json:"take_position"`              // 取板位置
	WeekCode         int       `gorm:"column:WEEK_CODE" json:"week_code"`                      // 周码
	Layers           int       `gorm:"column:LAYERS" json:"layers"`                            // 层数
	LayerLevel       string    `gorm:"column:LAYER_LEVEL" json:"layer_level"`                  // 层别
	CustomerCode     string    `gorm:"column:CUSTOMER_CODE" json:"customer_code"`              // 客户代号
	State            string    `gorm:"column:STATE" json:"state"`                              // 检验状态（10-待检验、20-已退回、30-已作废、40-检验中、50-待审核、60-已完成,5.没有同步样本id），5未到达检验站点
	LotNumber        string    `gorm:"column:LOT_NUMBER" json:"lot_number"`                    // 检验数量＝批次数量
	CheckNumber      int       `gorm:"column:CHECK_NUMBER" json:"check_number"`                // 抽检数量
	AccNumber        int       `gorm:"column:ACC_NUMBER" json:"acc_number"`                    // 合格数量
	Unit             string    `gorm:"column:UNIT" json:"unit"`                                // 单位代码
	RejNumber        int       `gorm:"column:REJ_NUMBER" json:"rej_number"`                    // 不合格数量
	DecisionResult   string    `gorm:"column:DECISION_RESULT" json:"decision_result"`          // 判定结果
	OriDecisionResult string   `gorm:"column:ORI_DECISION_RESULT" json:"ori_decision_result"`  // 原始判定结果
	DecisionRemark   string    `gorm:"column:DECISION_REMARK" json:"decision_remark"`          // 判定结果说明
	OriCheckCode     string    `gorm:"column:ORI_CHECK_CODE" json:"ori_check_code"`            // 原检验单号
	CustomerModel    string    `gorm:"column:CUSTOMER_MODEL" json:"customer_model"`            // 客户型号
	PostTime         time.Time `gorm:"column:POST_TIME" json:"post_time"`                      // 提交时间
	AuditTime        time.Time `gorm:"column:AUDIT_TIME" json:"audit_time"`                    // 审核时间
	Auditby          string    `gorm:"column:AUDITBY" json:"auditby"`                          // 审核人主键（指向AD_USER.OBJECT_RRN）
	Remark           string    `gorm:"column:REMARK" json:"remark"`                            // 备注
	Stage            string    `gorm:"column:STAGE" json:"stage"`                              // 阶段
	NextAuditName    string    `gorm:"column:NEXT_AUDIT_NAME" json:"next_audit_name"`          // 二级审批人
	NextAuditTime    time.Time `gorm:"column:NEXT_AUDIT_TIME" json:"next_audit_time"`          // 二级审批时间
	IsRecheck        string    `gorm:"column:IS_RECHECK" json:"is_recheck"`                    // 是否复检(Y/N)
	DefectType       string    `gorm:"column:DEFECT_TYPE" json:"defect_type"`                  // 缺陷类型
	RecheckRemark    string    `gorm:"column:RECHECK_REMARK" json:"recheck_remark"`            // 附件备注
	Attachment       string    `gorm:"column:ATTACHMENT" json:"attachment"`                    // 附件
	DefectLevel      string    `gorm:"column:DEFECT_LEVEL" json:"defect_level"`                // 缺陷等级
	Source           int       `gorm:"column:SOURCE" json:"source"`                            // 任务来源，1：OIC创建(手动创建)  2：MES创建(pda创建).3:是首批下的子批
	EquipmentId      string    `gorm:"column:EQUIPMENT_ID" json:"equipment_id"`                // 设备
	CheckNameRrn     int64     `gorm:"column:CHECK_NAME_RRN" json:"check_name_rrn"`            // 检验方案主键
	SonLotId         string    `gorm:"column:SON_LOT_ID" json:"son_lot_id"`                    // 子批号
	Ipqctype         string    `gorm:"column:IPQCTYPE" json:"ipqctype"`                        // IPQC检验类型（自检/半检）
	Checkwhy         int       `gorm:"column:CHECKWHY" json:"checkwhy"`                        // 触发方式，0：自动；1：手动PDA
	Entermark        string    `gorm:"column:ENTERMARK" json:"entermark"`                      // 录入标识
	DeleteRemark     string    `gorm:"column:DELETE_REMARK" json:"delete_remark"`              // 删除备注信息
	NcmCode          string    `gorm:"column:NCM_CODE" json:"ncm_code"`                        // 不合格任务单号
	NcnCode          string    `gorm:"column:NCN_CODE" json:"ncn_code"`                        // NCN单号
	Entryby          string    `gorm:"column:ENTRYBY" json:"entryby"`                          // 录入人
}

// TableName 指定 Oracle 表名（必须大写）
func (IPqcCheckTask) TableName() string {
	return "IPQC_CHECK_TASK"
}