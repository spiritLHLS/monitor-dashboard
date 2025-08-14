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
		if digitalProduct.OriginId == nil {
			digitalProduct.OriginId = new(int)
		}
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
- CPU: ` + product.Cpu + `
- 内存: ` + product.Memory + `
- 硬盘: ` + product.Disk + `
- 流量: ` + product.Traffic + `
- 带宽: ` + product.PortSpeed + `
- 位置: ` + product.Location + `
- 价格: ` + product.Price + `
- 其他信息: ` + product.Additional + `
请提取并转换为以下JSON格式，数值类型请转换为浮点数（如果无法确定数值则设为null）：
{
  "cpu": float64 or null,
  "memory": float64 or null,
  "disk": float64 or null,
  "traffic": float64 or null,
  "portSpeed": float64 or null,
  "location": "string",
  "price": float64 or null,
  "priceUnit": "string",
  "additional": "string"
}
提取规则：
1. CPU：提取核心数，如"4核心"→4.0，"2vCPU"→2.0，"1.5核"→1.5，不存在的设为null
2. 内存：提取GB数值，如"8GB"→8.0，"4G"→4.0，"1.5GB"→1.5，如果不存在单位，假设单位为GB直接提取数值，不存在的设为null
3. 硬盘：提取GB数值，如"100GB SSD"→100.0，"1TB"→1024.0，"512MB"→0.512，如果不存在单位，假设单位为GB直接提取数值，不存在的设为null
4. 流量：提取TB数值，如"10TB"→10.0，"不限"→10240.0，"500GB"→0.5，如果不存在单位，假设单位为TB直接提取数值不存在的设为null
5. 带宽：提取Gbps数值，如"1Gbps"→1.0，"100Mbps"→0.1，"50Mbps"→0.05，小于1Mbps或不存在的设为null
6. 价格：提取数值部分（保留小数），如"$10.99/月"→10.99，"$5.5/月"→5.5，"HK$151.80HKD"这种应该识别为港币，如果存在HKD、GBP等内容，应当优先识别英文代表的价格单位；如果非美元计价，则请按照当前汇率直接转换为美元计价的，不存在的设为null，如果存在多个，那么仅保留最小的那个，一般是按月计费的那个，不存在的设为null
7. 价格单位：提取单位，如"/月"、"/年"、"monthly"等，单年付的设为"annually"，月的设为"monthly"，两年付的三年付的依此类推，不存在的设为"monthly"，如果有多个都仅保留月的
8. 位置：必须按照"城市, 省份, 国家"格式提取，使用英文全称城市和省份名称，国家使用英文二字母缩写。如果某部分信息不存在则留空但保留逗号。例如：
   - "New York, New York, US"（完整信息）
   - "Tokyo, , JP"（无省份信息）
   - ", , US"（仅有国家信息）
   - "Los Angeles, California, US"（完整信息）
   如果完全无法确定位置信息，则设为", , "
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
		Tag        string   `json:"tag"`
		Cpu        *float64 `json:"cpu"`
		Memory     *float64 `json:"memory"`
		Disk       *float64 `json:"disk"`
		Traffic    *float64 `json:"traffic"`
		PortSpeed  *float64 `json:"portSpeed"`
		Location   string   `json:"location"`
		Price      *float64 `json:"price"`
		PriceUnit  string   `json:"priceUnit"`
		Additional string   `json:"additional"`
	}
	err := json.Unmarshal([]byte(jsonStr), &aiResult)
	if err != nil {
		return digitalproducts.DigitalProducts{}, fmt.Errorf("JSON解析失败: %v", err)
	}
	// 安全地处理字符串指针
	var tag, location, priceUnit, additional *string
	// 处理tag
	if originalProduct.Tag != "" {
		tag = &originalProduct.Tag
	} else {
		emptyStr := ""
		tag = &emptyStr
	}
	// 处理location
	if aiResult.Location != "" && aiResult.Location != "string" {
		location = &aiResult.Location
	} else if originalProduct.Location != "" {
		location = &originalProduct.Location
	} else {
		emptyStr := ""
		location = &emptyStr
	}
	// 处理priceUnit
	if aiResult.PriceUnit != "" && aiResult.PriceUnit != "string" {
		priceUnit = &aiResult.PriceUnit
	} else {
		defaultUnit := "monthly"
		priceUnit = &defaultUnit
	}
	// 处理additional
	if aiResult.Additional != "" && aiResult.Additional != "string" {
		additional = &aiResult.Additional
	} else if originalProduct.Additional != "" {
		additional = &originalProduct.Additional
	} else {
		emptyStr := ""
		additional = &emptyStr
	}
	// 构建数字商品对象
	digitalProduct := digitalproducts.DigitalProducts{
		Tag:        tag,
		Cpu:        aiResult.Cpu,
		Memory:     aiResult.Memory,
		Disk:       aiResult.Disk,
		Traffic:    aiResult.Traffic,
		PortSpeed:  aiResult.PortSpeed,
		Location:   location,
		Price:      aiResult.Price,
		PriceUnit:  priceUnit,
		Additional: additional,
		Stock:      new(int),
		OriginId:   new(int),
	}
	// 处理Stock字段
	if originalProduct.Stock != nil {
		if *originalProduct.Stock == 1000 {
			*digitalProduct.Stock = 6
		} else {
			*digitalProduct.Stock = *originalProduct.Stock
		}
	} else {
		*digitalProduct.Stock = 0 // 默认值
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

// BatchConvertProductsToDigital 批量转换指定数字商品记录（重新从原表转换并覆写）
func (dpdService *DigitalProductsService) BatchConvertProductsToDigital(digitalProductIds []uint, userID uint) error {
	dpdService.conversionMutex.Lock()
	defer dpdService.conversionMutex.Unlock()
	// 检查是否有任务正在运行
	if dpdService.conversionStatus != nil && dpdService.conversionStatus.IsRunning {
		return fmt.Errorf("转换任务正在进行中，当前步骤：%s，进度：%d/%d，成功：%d，失败：%d",
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
		CurrentStep: "初始化批量重新转换",
		Progress:    0,
		UserID:      userID,
	}
	// 启动后台任务
	go dpdService.runBatchReconversionTask(digitalProductIds, userID)
	return nil
}

// runBatchReconversionTask 在后台运行批量重新转换任务
func (dpdService *DigitalProductsService) runBatchReconversionTask(digitalProductIds []uint, userID uint) {
	defer func() {
		dpdService.conversionMutex.Lock()
		if dpdService.conversionStatus != nil {
			dpdService.conversionStatus.IsRunning = false
		}
		dpdService.conversionMutex.Unlock()
		if r := recover(); r != nil {
			global.GVA_LOG.Error(fmt.Sprintf("批量重新转换任务异常退出: %v", r))
		}
	}()
	// 第一步：获取选中的数字商品记录
	dpdService.updateStatus("获取选中的数字商品记录", 0, len(digitalProductIds))
	var digitalProductsList []digitalproducts.DigitalProducts
	err := global.GVA_DB.Where("id IN ?", digitalProductIds).Find(&digitalProductsList).Error
	if err != nil {
		errMsg := fmt.Sprintf("获取数字商品记录失败: %v", err)
		global.GVA_LOG.Error(errMsg)
		dpdService.updateStatusError(errMsg)
		return
	}
	if len(digitalProductsList) == 0 {
		global.GVA_LOG.Info("没有找到选中的数字商品记录")
		dpdService.updateStatus("完成", 0, 0)
		return
	}
	// 第二步：提取原表ID并获取对应的原表记录
	dpdService.updateStatus("获取对应的原表记录", 0, len(digitalProductsList))
	var originIds []int
	digitalProductMap := make(map[int]digitalproducts.DigitalProducts) // originId -> digitalProduct
	for _, dp := range digitalProductsList {
		if dp.OriginId != nil && *dp.OriginId > 0 {
			originIds = append(originIds, *dp.OriginId)
			digitalProductMap[*dp.OriginId] = dp
		}
	}
	if len(originIds) == 0 {
		global.GVA_LOG.Info("选中的数字商品记录中没有有效的原表ID")
		dpdService.updateStatus("完成", 0, len(digitalProductsList))
		return
	}
	var productsList []products.Products
	err = global.GVA_DB.Where("id IN ?", originIds).Find(&productsList).Error
	if err != nil {
		errMsg := fmt.Sprintf("获取原表记录失败: %v", err)
		global.GVA_LOG.Error(errMsg)
		dpdService.updateStatusError(errMsg)
		return
	}
	global.GVA_LOG.Info(fmt.Sprintf("找到 %d 个原表记录需要重新转换", len(productsList)))
	// 更新总数
	dpdService.conversionMutex.Lock()
	if dpdService.conversionStatus != nil {
		dpdService.conversionStatus.TotalProducts = len(digitalProductsList)
	}
	dpdService.conversionMutex.Unlock()
	// 第三步：处理所有选中的数字商品记录
	processedCount := 0
	for _, digitalProduct := range digitalProductsList {
		processedCount++
		dpdService.updateStatus(fmt.Sprintf("处理数字商品 ID:%d", digitalProduct.ID), processedCount-1, len(digitalProductsList))
		// 检查是否有有效的原表ID
		if digitalProduct.OriginId == nil || *digitalProduct.OriginId <= 0 {
			global.GVA_LOG.Error(fmt.Sprintf("数字商品ID %d 没有有效的原表ID", digitalProduct.ID))
			dpdService.incrementFailCount()
			continue
		}
		originId := *digitalProduct.OriginId
		// 查找对应的原表记录
		var foundProduct *products.Products
		for _, product := range productsList {
			if int(product.ID) == originId {
				foundProduct = &product
				break
			}
		}
		// 如果原表记录不存在（已被删除），删除对应的数字商品记录
		if foundProduct == nil {
			global.GVA_LOG.Info(fmt.Sprintf("原表记录ID %d 已删除，删除对应的数字商品记录ID %d", originId, digitalProduct.ID))
			deleteErr := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
				if err := tx.Model(&digitalproducts.DigitalProducts{}).Where("id = ?", digitalProduct.ID).Update("deleted_by", userID).Error; err != nil {
					return err
				}
				if err := tx.Delete(&digitalproducts.DigitalProducts{}, "id = ?", digitalProduct.ID).Error; err != nil {
					return err
				}
				return nil
			})
			if deleteErr != nil {
				global.GVA_LOG.Error(fmt.Sprintf("删除数字商品记录失败，ID %d: %v", digitalProduct.ID, deleteErr))
				dpdService.incrementFailCount()
			} else {
				global.GVA_LOG.Info(fmt.Sprintf("成功删除数字商品记录ID %d（原表记录已不存在）", digitalProduct.ID))
				dpdService.incrementSuccessCount()
			}
			continue
		}
		// 使用AI重新解析产品信息
		newDigitalProduct, parseErr := dpdService.parseProductWithAI(*foundProduct)
		if parseErr != nil {
			global.GVA_LOG.Error(fmt.Sprintf("AI重新解析产品ID %d 失败: %v", foundProduct.ID, parseErr))
			dpdService.incrementFailCount()
			continue
		}
		// 保持原有的ID、创建时间等不变，只更新解析出的字段
		newDigitalProduct.ID = digitalProduct.ID
		newDigitalProduct.CreatedAt = digitalProduct.CreatedAt
		newDigitalProduct.CreatedBy = digitalProduct.CreatedBy
		newDigitalProduct.UpdatedBy = userID
		if newDigitalProduct.OriginId == nil {
			newDigitalProduct.OriginId = new(int)
		}
		*newDigitalProduct.OriginId = int(foundProduct.ID)
		// 覆写记录
		saveErr := global.GVA_DB.Save(&newDigitalProduct).Error
		if saveErr != nil {
			global.GVA_LOG.Error(fmt.Sprintf("覆写数字商品失败，产品ID %d: %v", foundProduct.ID, saveErr))
			dpdService.incrementFailCount()
			continue
		}
		dpdService.incrementSuccessCount()
		global.GVA_LOG.Info(fmt.Sprintf("成功重新转换并覆写产品ID %d", foundProduct.ID))
	}
	dpdService.updateStatus("批量重新转换完成", len(digitalProductsList), len(digitalProductsList))
	global.GVA_LOG.Info(fmt.Sprintf("批量重新转换完成: 成功 %d 个，失败 %d 个",
		dpdService.conversionStatus.SuccessCount, dpdService.conversionStatus.FailCount))
}
