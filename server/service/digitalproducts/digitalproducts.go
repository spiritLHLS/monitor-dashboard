package digitalproducts

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"server/global"
	"server/model/digitalproducts"
	digitalproductsReq "server/model/digitalproducts/request"
	"server/model/products"
	"server/service/ai"
	"strings"
	"sync"
	"time"
)

// ConversionStatus 转换任务状态
type ConversionStatus struct {
	IsRunning     bool      `json:"is_running"`
	StartTime     time.Time `json:"start_time"`
	CurrentStep   string    `json:"current_step"`
	Progress      int       `json:"progress"`       // 已处理的产品数量
	TotalProducts int       `json:"total_products"` // 总产品数量
	SuccessCount  int       `json:"success_count"`
	FailCount     int       `json:"fail_count"`
	LastError     string    `json:"last_error"`
	UserID        uint      `json:"user_id"`
}

type DigitalProductsService struct {
	conversionMutex  sync.Mutex
	conversionStatus *ConversionStatus
}

// CreateDigitalProducts 创建数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) CreateDigitalProducts(dpd *digitalproducts.DigitalProducts) (err error) {
	err = global.GVA_DB.Create(dpd).Error
	return err
}

// DeleteDigitalProducts 删除数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) DeleteDigitalProducts(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&digitalproducts.DigitalProducts{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&digitalproducts.DigitalProducts{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteDigitalProductsByIds 批量删除数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) DeleteDigitalProductsByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&digitalproducts.DigitalProducts{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&digitalproducts.DigitalProducts{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateDigitalProducts 更新数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) UpdateDigitalProducts(dpd digitalproducts.DigitalProducts) (err error) {
	err = global.GVA_DB.Model(&digitalproducts.DigitalProducts{}).Where("id = ?", dpd.ID).Updates(&dpd).Error
	return err
}

// GetDigitalProducts 根据ID获取数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) GetDigitalProducts(ID string) (dpd digitalproducts.DigitalProducts, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dpd).Error
	return
}

// GetDigitalProductsInfoList 分页获取数字商品表记录
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) GetDigitalProductsInfoList(info digitalproductsReq.DigitalProductsSearch) (list []digitalproducts.DigitalProducts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&digitalproducts.DigitalProducts{})
	var dpds []digitalproducts.DigitalProducts
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PriceUnit != nil && *info.PriceUnit != "" {
		db = db.Where("price_unit = ?", *info.PriceUnit)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&dpds).Error
	return dpds, total, err
}

func (dpdService *DigitalProductsService) GetDigitalProductsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// ConvertProductsToDigital 将products表数据转换为数字商品表，并清理无效记录
func (dpdService *DigitalProductsService) ConvertProductsToDigital(userID uint) (successCount int, failCount int, err error) {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	// 检查是否有任务正在运行
	if dpdService.conversionStatus != nil && dpdService.conversionStatus.IsRunning {
		return 0, 0, fmt.Errorf("转换任务正在进行中，当前步骤：%s，进度：%d/%d，成功：%d，失败：%d",
			dpdService.conversionStatus.CurrentStep,
			dpdService.conversionStatus.Progress,
			dpdService.conversionStatus.TotalProducts,
			dpdService.conversionStatus.SuccessCount,
			dpdService.conversionStatus.FailCount)
	}
	// 初始化状态
	dpdService.conversionStatus = &ConversionStatus{
		IsRunning:   true,
		StartTime:   time.Now(),
		CurrentStep: "初始化",
		Progress:    0,
		UserID:      userID,
	}
	// 启动后台任务
	go dpdService.runConversionTask(userID)
	return 0, 0, nil
}

// runConversionTask 在后台运行转换任务
func (dpdService *DigitalProductsService) runConversionTask(userID uint) {
	defer func() {
		dpdService.conversionMutex.Lock()
		if dpdService.conversionStatus != nil {
			dpdService.conversionStatus.IsRunning = false
		}
		dpdService.conversionMutex.Unlock()
		if r := recover(); r != nil {
			global.GVA_LOG.Error(fmt.Sprintf("转换任务异常退出: %v", r))
		}
	}()
	// 第一步：清理无效的数字商品记录
	dpdService.updateStatus("清理无效记录", 0, 0)
	cleanupResult := global.GVA_DB.Where("origin_id NOT IN (?)",
		global.GVA_DB.Model(&products.Products{}).Select("id")).
		Delete(&digitalproducts.DigitalProducts{})
	if cleanupResult.Error != nil {
		errMsg := fmt.Sprintf("清理无效数字商品记录失败: %v", cleanupResult.Error)
		global.GVA_LOG.Error(errMsg)
		dpdService.updateStatusError(errMsg)
		return
	}
	if cleanupResult.RowsAffected > 0 {
		global.GVA_LOG.Info(fmt.Sprintf("已清理 %d 条无效的数字商品记录", cleanupResult.RowsAffected))
	}
	// 第二步：获取所有未转换的产品
	dpdService.updateStatus("获取待转换产品", 0, 0)
	var productsList []products.Products
	err := global.GVA_DB.Where("id NOT IN (?)",
		global.GVA_DB.Model(&digitalproducts.DigitalProducts{}).Select("origin_id").Where("origin_id IS NOT NULL")).
		Find(&productsList).Error
	if err != nil {
		errMsg := fmt.Sprintf("获取未转换产品信息失败: %v", err)
		global.GVA_LOG.Error(errMsg)
		dpdService.updateStatusError(errMsg)
		return
	}
	if len(productsList) == 0 {
		global.GVA_LOG.Info("没有需要转换的新产品")
		dpdService.updateStatus("完成", 0, len(productsList))
		return
	}
	global.GVA_LOG.Info(fmt.Sprintf("找到 %d 个需要转换的产品", len(productsList)))
	// 更新总数
	dpdService.conversionMutex.Lock()
	if dpdService.conversionStatus != nil {
		dpdService.conversionStatus.TotalProducts = len(productsList)
	}
	dpdService.conversionMutex.Unlock()
	// 第三步：逐个处理产品转换
	for i, product := range productsList {
		dpdService.updateStatus(fmt.Sprintf("转换产品 ID:%d", product.ID), i, len(productsList))
		// 使用AI解析产品信息
		digitalProduct, parseErr := dpdService.parseProductWithAI(product)
		if parseErr != nil {
			global.GVA_LOG.Error(fmt.Sprintf("AI解析产品ID %d 失败: %v", product.ID, parseErr))
			dpdService.incrementFailCount()
			continue
		}
		// 设置创建者和原始ID
		digitalProduct.CreatedBy = userID
		*digitalProduct.OriginId = int(product.ID)
		// 保存到数字商品表
		createErr := global.GVA_DB.Create(&digitalProduct).Error
		if createErr != nil {
			global.GVA_LOG.Error(fmt.Sprintf("保存数字商品失败，产品ID %d: %v", product.ID, createErr))
			dpdService.incrementFailCount()
			continue
		}
		dpdService.incrementSuccessCount()
		global.GVA_LOG.Info(fmt.Sprintf("成功转换产品ID %d", product.ID))
	}
	dpdService.updateStatus("完成", len(productsList), len(productsList))
	global.GVA_LOG.Info(fmt.Sprintf("转换完成: 成功 %d 个，失败 %d 个",
		dpdService.conversionStatus.SuccessCount, dpdService.conversionStatus.FailCount))
}

// updateStatus 更新任务状态
func (dpdService *DigitalProductsService) updateStatus(step string, progress int, total int) {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	if dpdService.conversionStatus != nil {
		dpdService.conversionStatus.CurrentStep = step
		dpdService.conversionStatus.Progress = progress
		if total > 0 {
			dpdService.conversionStatus.TotalProducts = total
		}
	}
}

// updateStatusError 更新错误状态
func (dpdService *DigitalProductsService) updateStatusError(errMsg string) {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	if dpdService.conversionStatus != nil {
		dpdService.conversionStatus.LastError = errMsg
	}
}

// incrementSuccessCount 增加成功计数
func (dpdService *DigitalProductsService) incrementSuccessCount() {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	if dpdService.conversionStatus != nil {
		dpdService.conversionStatus.SuccessCount++
	}
}

// incrementFailCount 增加失败计数
func (dpdService *DigitalProductsService) incrementFailCount() {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	if dpdService.conversionStatus != nil {
		dpdService.conversionStatus.FailCount++
	}
}

// GetConversionStatus 获取转换任务状态
func (dpdService *DigitalProductsService) GetConversionStatus() *ConversionStatus {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	if dpdService.conversionStatus == nil {
		return nil
	}
	status := *dpdService.conversionStatus
	return &status
}

// parseProductWithAI 使用AI解析产品信息
func (dpdService *DigitalProductsService) parseProductWithAI(product products.Products) (digitalproducts.DigitalProducts, error) {
	// 构建AI提示词
	prompt := dpdService.buildPromptForProduct(product)
	// 调用AI服务获取解析结果
	aiResponse := ai.GetResponse(prompt)
	if aiResponse == "" {
		return digitalproducts.DigitalProducts{}, fmt.Errorf("AI返回空响应")
	}
	// 解析AI响应并构建数字商品对象
	digitalProduct, err := dpdService.parseAIResponse(aiResponse, product)
	if err != nil {
		return digitalproducts.DigitalProducts{}, fmt.Errorf("解析AI响应失败: %v", err)
	}
	return digitalProduct, nil
}

// buildPromptForProduct 为产品构建AI提示词
func (dpdService *DigitalProductsService) buildPromptForProduct(product products.Products) string {
	prompt := `请从以下商品信息中提取规范化的硬件配置信息，并以JSON格式返回。请严格按照以下格式返回，不要包含任何其他文字：
商品信息：
- TAG: ` + product.Tag + `
- CPU: ` + product.Cpu + `
- 内存: ` + product.Memory + `
- 硬盘: ` + product.Disk + `
- 流量: ` + product.Traffic + `
- 带宽: ` + product.PortSpeed + `
- 位置: ` + product.Location + `
- 价格: ` + product.Price + `
- 其他信息: ` + product.Additional + `
请提取并转换为以下JSON格式，数值类型请转换为整数（如果无法确定数值则设为null）：
{
  "tag": "string",
  "cpu": int or null,
  "memory": int or null,
  "disk": int or null,
  "traffic": int or null,
  "portSpeed": int or null,
  "location": "string",
  "price": int or null,
  "priceUnit": "string",
  "additional": "string"
}
提取规则：
1. CPU：提取核心数，如"4核心"→4，"2vCPU"→2，不存在的设为null
2. 内存：提取GB数值，如"8GB"→8，"4G"→4，不存在的设为null
3. 硬盘：提取GB数值，如"100GB SSD"→100，"1TB"→1024，不存在的设为null
4. 流量：提取TB数值，如"10TB"→10，"不限"→10240，不存在的设为null
5. 带宽：提取Gbps数值，如"1Gbps"→1，"100Mbps"→0.1（小于1Mbps或不存在的设为null）
6. 价格：只提取数值部分，如"$10/月"→10，如果非美元计价，则请按照当前汇率直接转换为美元计价的，不存在的设为null
7. 价格单位：提取单位，如"/月"、"/年"、"monthly"等，单年付的设为"Annually"，月的设为"monthly"，两年付的三年付的依此类推，不存在的设为null
`
	return prompt
}

// parseAIResponse 解析AI响应并构建数字商品对象
func (dpdService *DigitalProductsService) parseAIResponse(aiResponse string, originalProduct products.Products) (digitalproducts.DigitalProducts, error) {
	jsonStr := dpdService.extractJSON(aiResponse)
	if jsonStr == "" {
		return digitalproducts.DigitalProducts{}, fmt.Errorf("无法从AI响应中提取JSON")
	}
	var aiResult struct {
		Tag        string `json:"tag"`
		Cpu        *int   `json:"cpu"`
		Memory     *int   `json:"memory"`
		Disk       *int   `json:"disk"`
		Traffic    *int   `json:"traffic"`
		PortSpeed  *int   `json:"portSpeed"`
		Location   string `json:"location"`
		Price      *int   `json:"price"`
		PriceUnit  string `json:"priceUnit"`
		Additional string `json:"additional"`
	}
	err := json.Unmarshal([]byte(jsonStr), &aiResult)
	if err != nil {
		return digitalproducts.DigitalProducts{}, fmt.Errorf("JSON解析失败: %v", err)
	}
	// 构建数字商品对象
	digitalProduct := digitalproducts.DigitalProducts{
		Tag:        &aiResult.Tag,
		Cpu:        aiResult.Cpu,
		Memory:     aiResult.Memory,
		Disk:       aiResult.Disk,
		Traffic:    aiResult.Traffic,
		PortSpeed:  aiResult.PortSpeed,
		Location:   &aiResult.Location,
		Price:      aiResult.Price,
		PriceUnit:  &aiResult.PriceUnit,
		Additional: &aiResult.Additional,
		Stock:      new(int),
	}
	// 如果AI解析失败，使用原始数据作为备用
	if digitalProduct.Tag == nil || *digitalProduct.Tag == "" {
		digitalProduct.Tag = &originalProduct.Tag
	}
	if digitalProduct.Location == nil || *digitalProduct.Location == "" {
		digitalProduct.Location = &originalProduct.Location
	}
	if digitalProduct.Additional == nil || *digitalProduct.Additional == "" {
		digitalProduct.Additional = &originalProduct.Additional
	}
	if originalProduct.Stock != nil && *originalProduct.Stock == 1000 {
		*digitalProduct.Stock = 6
	} else if originalProduct.Stock != nil && *originalProduct.Stock != 1000 {
		*digitalProduct.Stock = *originalProduct.Stock
	}
	return digitalProduct, nil
}

// extractJSON 从AI响应中提取JSON字符串
func (dpdService *DigitalProductsService) extractJSON(response string) string {
	// 查找JSON开始和结束位置
	start := strings.Index(response, "{")
	if start == -1 {
		return ""
	}
	// 从后往前查找最后一个 }
	end := strings.LastIndex(response, "}")
	if end == -1 || end <= start {
		return ""
	}
	return response[start : end+1]
}
