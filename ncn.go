package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Ncn 结构体用于处理NCN相关操作
type Ncn struct {
	ctx    context.Context
	config *Config
}

// NewNcn 创建Ncn实例
func NewNcn() *Ncn {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		fmt.Printf("配置加载错误: %v\n", err)
		config = defaultConfig() // 使用默认配置
	}

	return &Ncn{
		config: config,
	}
}

// startup 初始化Ncn
func (n *Ncn) startup(ctx context.Context) {
	n.ctx = ctx
}

// SetConfig 设置配置
func (n *Ncn) SetConfig(config *Config) {
	n.config = config
}

// GetConfig 获取配置
func (n *Ncn) GetConfig() *Config {
	return n.config
}

// SearchRequest 搜索请求结构体
type SearchRequest struct {
	Request struct {
		Header struct {
			MessageName   string `json:"MESSAGENAME"`
			TransactionID string `json:"TRANSACTIONID"`
			OrgRRN        string `json:"ORGRRN"`
			OrgName       string `json:"ORGNAME"`
			UserName      string `json:"USERNAME"`
			Language      string `json:"LANGUAGE"`
			IP            string `json:"IP"`
		} `json:"Header"`
		Body struct {
			EntityModel   string `json:"ENTITYMODEL"`
			QueryText     string `json:"QUERYTEXT"`
			QueryName     string `json:"QUERYNAME"`
			Max           string `json:"MAX"`
			First         string `json:"FIRST"`
			OrderByClause string `json:"ORDERBYCLAUSE"`
			WhereClause   string `json:"WHERECLAUSE"`
		} `json:"Body"`
	} `json:"Request"`
}

// SearchResponse 搜索响应结构体
type SearchResponse struct {
	Response struct {
		Header struct {
			TransactionID string `json:"TRANSACTIONID"`
			Result        string `json:"RESULT"`
		} `json:"Header"`
		Body struct {
			DataList []struct {
				CAR_RESP_DEPT          string `json:"CAR_RESP_DEPT"`
				OCAP_AUDIT_PERSON      string `json:"OCAP_AUDIT_PERSON"`
				IS_LOCK                string `json:"IS_LOCK"`
				SYSTEM_MODULE          string `json:"SYSTEM_MODULE"`
				IS_ACTIVE              string `json:"IS_ACTIVE"`
				PRC_DIS                string `json:"PRC_DIS"`
				CAR_AUDIT_UUID         string `json:"CAR_AUDIT_UUID"`
				STEPDESC               string `json:"STEPDESC"`
				STATUS                 string `json:"STATUS"`
				SOURCE_CODE            string `json:"SOURCE_CODE"`
				QA_VERIFIER_RESULT     string `json:"QA_VERIFIER_RESULT"`
				FINISH_TIME            string `json:"FINISH_TIME"`
				OCCURRED_TIME          string `json:"OCCURRED_TIME"`
				CAR_AUDIT_TIME         string `json:"CAR_AUDIT_TIME"`
				OBJECT_RRN             string `json:"OBJECT_RRN"`
				CAR_AUDIT_RESULT       string `json:"CAR_AUDIT_RESULT"`
				ORG_RRN                string `json:"ORG_RRN"`
				STEP_NAME              string `json:"STEP_NAME"`
				RISK_LEVEL             string `json:"RISK_LEVEL"`
				CREATED_BY             string `json:"CREATED_BY"`
				LOCK_TYPE              string `json:"LOCK_TYPE"`
				OCAP_NAME              string `json:"OCAP_NAME"`
				STAGE_NAME             string `json:"STAGE_NAME"`
				RESP_PERSON            string `json:"RESP_PERSON"`
				OCAP_FLOW_END_TIME     string `json:"OCAP_FLOW_END_TIME"`
				ABNORMAL_DEVICE_ID     string `json:"ABNORMAL_DEVICE_ID"`
				ABNORMAL_DESC          string `json:"ABNORMAL_DESC"`
				CAR_AUDIT_DEPT         string `json:"CAR_AUDIT_DEPT"`
				QA_RESP_PERSON         string `json:"QA_RESP_PERSON"`
				RECEIVE_TIME           string `json:"RECEIVE_TIME"`
				OCAP_FLOW_START_TIME   string `json:"OCAP_FLOW_START_TIME"`
				IS_OA                  string `json:"IS_OA"`
				MONTH_COUNT            string `json:"MONTH_COUNT"`
				MAIN_EXC_TYPE_CODE     string `json:"MAIN_EXC_TYPE_CODE"`
				ABNORMAL_GOV_TOOL      string `json:"ABNORMAL_GOV_TOOL"`
				MONCOUNT               string `json:"MONCOUNT"`
				UPDATED_BY             string `json:"UPDATED_BY"`
				CAR_RESP_PERSON        string `json:"CAR_RESP_PERSON"`
				SUB_DEFECT_CODE        string `json:"SUB_DEFECT_CODE"`
				ABNORMAL_MATERIAL_LOT  string `json:"ABNORMAL_MATERIAL_LOT"`
				LOCK_VERSION           string `json:"LOCK_VERSION"`
				RISK_DIMENSION         string `json:"RISK_DIMENSION"`
				SUB_EXC_TYPE_CODE      string `json:"SUB_EXC_TYPE_CODE"`
				IS_MEASURE             string `json:"IS_MEASURE"`
				AUDIT_REMARK           string `json:"AUDIT_REMARK"`
				UPDATED                string `json:"UPDATED"`
				OCAP_CODE              string `json:"OCAP_CODE"`
				CREATED                string `json:"CREATED"`
				QA_VERIFIER_PERSON     string `json:"QA_VERIFIER_PERSON"`
				OCAP_FLOW_CODE         string `json:"OCAP_FLOW_CODE"`
				INITIATOR_RESULT       string `json:"INITIATOR_RESULT"`
				OCAP_AUDIT_RESULT      string `json:"OCAP_AUDIT_RESULT"`
				CAR_OA_CODE            string `json:"CAR_OA_CODE"`
				MAIN_DEFECT_CODE       string `json:"MAIN_DEFECT_CODE"`
				APPLICANT              string `json:"APPLICANT"`
				APPLY_DEPT             string `json:"APPLY_DEPT"`
				IS_FLOWED_OUT          string `json:"IS_FLOWED_OUT"`
				INSTRUMENT_CODE        string `json:"INSTRUMENT_CODE"`
				STAGEDESC              string `json:"STAGEDESC"`
				RESP_DEPT              string `json:"RESP_DEPT"`
				ABNORMAL_MATERIAL_NAME string `json:"ABNORMAL_MATERIAL_NAME"`
				TRIGGER_WAY            string `json:"TRIGGER_WAY"`
				OCAP_AUDIT_TIME        string `json:"OCAP_AUDIT_TIME"`
				TASK_CODE              string `json:"TASK_CODE"`
			} `json:"DATALIST"`
		} `json:"Body"`
	} `json:"Response"`
}

// Search 查询NCN数据
func (n *Ncn) Search(searchParams map[string]interface{}) ([]map[string]interface{}, error) {
	// 构造查询SQL
	baseSQL := ` SELECT  DISTINCT T.*, WP.DESCRIPTION AS STEPDESC, PS.DESCRIPTION AS STAGEDESC FROM ( SELECT DISTINCT MAIN.*, COUNT(*) OVER (PARTITION BY MAIN.STEP_NAME ,MAIN.STAGE_NAME,MAIN.MAIN_DEFECT_CODE,MAIN.MAIN_EXC_TYPE_CODE,TO_CHAR(CREATED,'YYYY-MM')  ) AS MONCOUNT FROM NCN_MAIN_TASK MAIN where MAIN.IS_ACTIVE='Y'   ) T  LEFT JOIN NCN_ABNORMAL_LOT nal ON T.TASK_CODE = NAL.NCN_CODE AND NAL.IS_ACTIVE ='Y'  LEFT JOIN PRD_STAGE PS ON PS.NAME = T.STAGE_NAME AND PS.IS_ACTIVE ='Y'  LEFT JOIN WF_PROCESSDEFINITION WP ON WP.NAME = T.STEP_NAME  AND WP.IS_ACTIVE='Y' AND CLASS = 'S'  where 1=1   `

	// 根据搜索参数添加WHERE条件
	if startTime, ok := searchParams["startTime"].(string); ok && startTime != "" {
		baseSQL += fmt.Sprintf(" and to_char(T.CREATED,'yyyy-MM-dd HH:mm:ss') >= '%s 00:00:00'", startTime)
	}
	if endTime, ok := searchParams["endTime"].(string); ok && endTime != "" {
		baseSQL += fmt.Sprintf(" and to_char(T.CREATED,'yyyy-MM-dd HH:mm:ss') <= '%s 23:59:59'", endTime)
	}
	if ncnNumber, ok := searchParams["ncnNumber"].(string); ok && ncnNumber != "" {
		baseSQL += fmt.Sprintf(" and T.TASK_CODE like '%%%s%%'", ncnNumber)
	}
	if systemModule, ok := searchParams["systemModule"].(string); ok && systemModule != "" {
		baseSQL += fmt.Sprintf(" and T.SYSTEM_MODULE = '%s'", systemModule)
	}
	if process, ok := searchParams["process"].(string); ok && process != "" {
		baseSQL += fmt.Sprintf(" and T.STEP_NAME like '%%%s%%'", process)
	}
	if flow, ok := searchParams["flow"].(string); ok && flow != "" {
		baseSQL += fmt.Sprintf(" and T.STAGE_NAME like '%%%s%%'", flow)
	}
	if riskLevel, ok := searchParams["riskLevel"].(string); ok && riskLevel != "" {
		baseSQL += fmt.Sprintf(" and T.RISK_DIMENSION = '%s'", riskLevel)
	}
	if applicant, ok := searchParams["applicant"].(string); ok && applicant != "" {
		baseSQL += fmt.Sprintf(" and T.CREATED_BY like '%%%s%%'", applicant)
	}
	if status, ok := searchParams["status"].(string); ok && status != "" {
		baseSQL += fmt.Sprintf(" and T.STATUS = '%s'", status)
	}
	if mainExceptionCategory, ok := searchParams["mainExceptionCategory"].(string); ok && mainExceptionCategory != "" {
		baseSQL += fmt.Sprintf(" and T.MAIN_EXC_TYPE_CODE = '%s'", mainExceptionCategory)
	}
	if mainDefect, ok := searchParams["mainDefect"].(string); ok && mainDefect != "" {
		baseSQL += fmt.Sprintf(" and T.MAIN_DEFECT_CODE = '%s'", mainDefect)
	}
	if responsibleUnit, ok := searchParams["responsibleUnit"].(string); ok && responsibleUnit != "" {
		baseSQL += fmt.Sprintf(" and T.RESP_DEPT = '%s'", responsibleUnit)
	}
	if ocapProcess, ok := searchParams["ocapProcess"].(string); ok && ocapProcess != "" {
		baseSQL += fmt.Sprintf(" and T.OCAP_FLOW_CODE = '%s'", ocapProcess)
	}
	if exceptionBatch, ok := searchParams["exceptionBatch"].(string); ok && exceptionBatch != "" {
		baseSQL += fmt.Sprintf(" and T.ABNORMAL_MATERIAL_LOT like '%%%s%%'", exceptionBatch)
	}
	if productModel, ok := searchParams["productModel"].(string); ok && productModel != "" {
		baseSQL += fmt.Sprintf(" and T.ABNORMAL_MATERIAL_NAME like '%%%s%%'", productModel)
	}
	if exceptionEquipment, ok := searchParams["exceptionEquipment"].(string); ok && exceptionEquipment != "" {
		baseSQL += fmt.Sprintf(" and T.ABNORMAL_DEVICE_ID like '%%%s%%'", exceptionEquipment)
	}
	if riskGrade, ok := searchParams["riskGrade"].(string); ok && riskGrade != "" {
		baseSQL += fmt.Sprintf(" and T.RISK_LEVEL = '%s'", riskGrade)
	}
	if responsiblePerson, ok := searchParams["responsiblePerson"].(string); ok && responsiblePerson != "" {
		baseSQL += fmt.Sprintf(" and T.RESP_PERSON = '%s'", responsiblePerson)
	}

	// 添加ORDER BY子句
	baseSQL += " order by t.created desc "

	// 构造搜索请求
	reqData := SearchRequest{}
	reqData.Request.Header.MessageName = "PCB.MES.GETENTITYKEYROOTMAPLISTBYQUERY"
	reqData.Request.Header.TransactionID = fmt.Sprintf("%d", time.Now().UnixNano()/1000000) // 生成时间戳作为交易ID
	reqData.Request.Header.OrgRRN = "378341"
	reqData.Request.Header.OrgName = "XCY"
	reqData.Request.Header.UserName = "TST00481"
	reqData.Request.Header.Language = "ZH"
	reqData.Request.Header.IP = "10.1.12.54"

	reqData.Request.Body.EntityModel = ""
	reqData.Request.Body.QueryText = baseSQL
	reqData.Request.Body.QueryName = ""
	reqData.Request.Body.Max = "99999"
	reqData.Request.Body.First = ""
	reqData.Request.Body.OrderByClause = ""
	reqData.Request.Body.WhereClause = ""

	// 转换为JSON字符串
	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		fmt.Printf("JSON序列化错误: %v\n", err)
		return nil, err
	}

	// 构造表单数据
	formData := url.Values{}
	formData.Set("senderId", "PCBMESSender")
	formData.Set("requestMessage", string(reqJSON))

	// 发送POST请求
	client := &http.Client{
		Timeout: time.Duration(n.config.GetESBAPITimeout()) * time.Second,
	}

	resp, err := client.PostForm(n.config.GetESBAPIURL(), formData)
	if err != nil {
		fmt.Printf("请求发送错误: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	var respData SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		fmt.Printf("响应解析错误: %v\n", err)
		return nil, err
	}

	// 检查请求结果
	if respData.Response.Header.Result != "SUCCESS" {
		fmt.Printf("搜索失败: 结果=%s\n", respData.Response.Header.Result)
		return nil, fmt.Errorf("搜索失败")
	}

	// 转换为map格式返回
	var result []map[string]interface{}
	for _, item := range respData.Response.Body.DataList {
		// 使用JSON序列化/反序列化方式将结构体转换为map
		itemJSON, err := json.Marshal(item)
		if err != nil {
			fmt.Printf("序列化数据项错误: %v\n", err)
			continue
		}
		var itemMap map[string]interface{}
		err = json.Unmarshal(itemJSON, &itemMap)
		if err != nil {
			fmt.Printf("反序列化数据项错误: %v\n", err)
			continue
		}
		result = append(result, itemMap)
	}

	return result, nil
}
