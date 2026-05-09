package main

import (
	"context"
	"fmt"
	"strings"

	"wailsix/models"
)

type Ipqc struct{
	ctx context.Context
}

func NewIpqc() *Ipqc {
	return &Ipqc{}
}

// SearchIpqcByCheckCode 根据检验单号进行模糊查询IPQC单据
// 参数:
//   - checkCode: 检验单号（支持模糊查询）
// 返回:
//   - []models.IPqcCheckTask: 匹配的IPQC单据列表（最多30条）
//   - error: 错误信息
func (i *Ipqc) SearchIpqcByCheckCode(checkCode string) ([]models.IPqcCheckTask, error) {
	if models.DB == nil {
		return nil, fmt.Errorf("数据库连接未初始化")
	}

	var ipqcTasks []models.IPqcCheckTask

	// 构建查询条件
	query := models.DB.Model(&models.IPqcCheckTask{})
	
	// 如果checkCode不为空，添加模糊查询条件
	if checkCode != "" {
		query = query.Where("CHECK_CODE LIKE ?", "%"+strings.ToUpper(checkCode)+"%")
	}

	// 添加条件：只查询有效的记录（IS_ACTIVE = 'Y'）
	query = query.Where("IS_ACTIVE = ?", "Y")

	// 限制返回最多30条记录，并按创建时间降序排列
	err := query.Limit(30).Order("CREATED DESC").Find(&ipqcTasks).Error
	if err != nil {
		return nil, fmt.Errorf("查询IPQC单据失败: %w", err)
	}

	return ipqcTasks, nil
}