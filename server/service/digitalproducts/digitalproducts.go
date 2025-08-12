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
)

type DigitalProductsService struct{}

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

// Service函数 - 添加到 DigitalProductsService 结构体中
// ConvertProductsToDigital 将products表数据转换为数字商品表
// Author [yourname](https://github.com/yourname)
func (dpdService *DigitalProductsService) ConvertProductsToDigital(productIds []uint, userID uint) (successCount int, failCount int, err error) {
	// 获取所有需要转换的产品
	var products []products.Products
	err = global.GVA_DB.Where("id IN ?", productIds).Find(&products).Error
	if err != nil {
		return 0, 0, fmt.Errorf("获取产品信息失败: %v", err)
	}
	if len(products) == 0 {
		return 0, 0, fmt.Errorf("未找到指定的产品")
	}
	successCount = 0
	failCount = 0
	// 逐个处理产品转换
	for _, product := range products {
		// 检查是否已经转换过
		var existingCount int64
		global.GVA_DB.Model(&digitalproducts.DigitalProducts{}).Where("origin_id = ?", product.ID).Count(&existingCount)
		if existingCount > 0 {
			global.GVA_LOG.Info(fmt.Sprintf("产品ID %d 已经转换过，跳过", product.ID))
			failCount++
			continue
		}
		// 使用AI解析产品信息
		digitalProduct, parseErr := dpdService.parseProductWithAI(product)
		if parseErr != nil {
			global.GVA_LOG.Error(fmt.Sprintf("AI解析产品ID %d 失败: %v", product.ID, parseErr))
			failCount++
			continue
		}
		// 设置创建者和原始ID
		digitalProduct.CreatedBy = userID
		*digitalProduct.OriginId = int(product.ID)
		// 保存到数字商品表
		createErr := global.GVA_DB.Create(&digitalProduct).Error
		if createErr != nil {
			global.GVA_LOG.Error(fmt.Sprintf("保存数字商品失败，产品ID %d: %v", product.ID, createErr))
			failCount++
			continue
		}
		successCount++
		global.GVA_LOG.Info(fmt.Sprintf("成功转换产品ID %d", product.ID))
	}
	return successCount, failCount, nil
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
  "additional": "string",
  "stock": int or null
}
提取规则：
1. CPU：提取核心数，如"4核心"→4，"2vCPU"→2，不存在的设为null
2. 内存：提取GB数值，如"8GB"→8，"4G"→4，不存在的设为null
3. 硬盘：提取GB数值，如"100GB SSD"→100，"1TB"→1024，不存在的设为null
4. 流量：提取TB数值，如"10TB"→10，"不限"→10240，不存在的设为null
5. 带宽：提取Gbps数值，如"1Gbps"→1，"100Mbps"→0.1（小于1Mbps或不存在的设为null）
6. 价格：只提取数值部分，如"$10/月"→10，如果非美元计价，则请按照当前汇率直接转换为美元计价的，不存在的设为null
7. 价格单位：提取单位，如"/月"、"/年"、"monthly"等，单年付的设为"Annually"，月的设为"monthly"，两年付的三年付的依此类推，不存在的设为null
8. 库存：如果有库存信息则提取数值，否则设为10`
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
		Stock      *int   `json:"stock"`
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
		Stock:      aiResult.Stock,
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
