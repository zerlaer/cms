package service

import (
	"cms/internal/models"
	"errors"

	"gorm.io/gorm"
)

type ComponentService struct {
	db *gorm.DB
}

func NewComponentService(db *gorm.DB) *ComponentService {
	return &ComponentService{db: db}
}

func (s *ComponentService) GetAll(page, pageSize int) ([]models.Component, int64, error) {
	var components []models.Component
	var total int64

	offset := (page - 1) * pageSize

	if err := s.db.Model(&models.Component{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := s.db.Offset(offset).Limit(pageSize).Order("id DESC").Find(&components).Error
	return components, total, err
}

func (s *ComponentService) GetByID(id uint) (*models.Component, error) {
	var component models.Component
	err := s.db.First(&component, id).Error
	if err != nil {
		return nil, err
	}
	return &component, nil
}

func (s *ComponentService) Create(component *models.Component) error {
	component.CurrentStock = component.Quantity
	component.StockIn = component.Quantity
	component.StockOut = 0
	return s.db.Create(component).Error
}

func (s *ComponentService) Update(id uint, component *models.Component) error {
	return s.db.Model(&models.Component{}).Where("id = ?", id).Updates(component).Error
}

func (s *ComponentService) Delete(id uint) error {
	return s.db.Delete(&models.Component{}, id).Error
}

func (s *ComponentService) StockIn(id uint, quantity int, remark string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var component models.Component
		if err := tx.First(&component, id).Error; err != nil {
			return err
		}

		component.Quantity += quantity
		component.StockIn += quantity
		component.CurrentStock = component.Quantity - component.StockOut

		if err := tx.Save(&component).Error; err != nil {
			return err
		}

		record := &models.StockRecord{
			ComponentID: id,
			Type:        "in",
			Quantity:    quantity,
			Remark:      remark,
		}
		return tx.Create(record).Error
	})
}

func (s *ComponentService) StockOut(id uint, quantity int, remark string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var component models.Component
		if err := tx.First(&component, id).Error; err != nil {
			return err
		}

		if component.CurrentStock < quantity {
			return errors.New("库存不足")
		}

		component.StockOut += quantity
		component.CurrentStock = component.Quantity - component.StockOut

		if err := tx.Save(&component).Error; err != nil {
			return err
		}

		record := &models.StockRecord{
			ComponentID: id,
			Type:        "out",
			Quantity:    quantity,
			Remark:      remark,
		}
		return tx.Create(record).Error
	})
}

func (s *ComponentService) GetStockRecords(componentID uint) ([]models.StockRecord, error) {
	var records []models.StockRecord
	err := s.db.Where("componentId = ?", componentID).Order("createdAt DESC").Find(&records).Error
	return records, err
}

func (s *ComponentService) GetAllStockRecords(page, pageSize int, recordType, componentName, startDate, endDate string) ([]map[string]interface{}, int64, error) {
	type RecordWithComponent struct {
		models.StockRecord
		ComponentName  string `gorm:"column:componentName"`
		ComponentModel string `gorm:"column:componentModel"`
	}

	var records []RecordWithComponent
	var total int64

	offset := (page - 1) * pageSize

	// 构建查询
	query := s.db.Model(&models.StockRecord{}).Joins("JOIN components ON stockRecords.componentId = components.id")

	// 应用筛选条件
	if recordType != "" {
		query = query.Where("stockRecords.type = ?", recordType)
	}

	if componentName != "" {
		query = query.Where("components.name LIKE ?", "%"+componentName+"%")
	}

	if startDate != "" {
		query = query.Where("stockRecords.createdAt >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("stockRecords.createdAt <= ?", endDate)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 执行查询
	if err := query.Select("stockRecords.*, components.name as componentName, components.model as componentModel").Offset(offset).Limit(pageSize).Order("stockRecords.createdAt DESC").Find(&records).Error; err != nil {
		return nil, 0, err
	}

	// 转换为带元件信息的记录
	result := make([]map[string]interface{}, len(records))
	for i, record := range records {
		result[i] = map[string]interface{}{
			"id":             record.ID,
			"componentId":    record.ComponentID,
			"componentName":  record.ComponentName,
			"componentModel": record.ComponentModel,
			"type":           record.Type,
			"quantity":       record.Quantity,
			"remark":         record.Remark,
			"createdAt":      record.CreatedAt,
		}
	}

	return result, total, nil
}
