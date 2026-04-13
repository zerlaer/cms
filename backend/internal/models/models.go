package models

import (
	"time"
)

type Component struct {
	ID           uint      `gorm:"column:id;primarykey" json:"id"`
	ProductCode  string    `gorm:"column:productCode;size:100;comment:商品编号" json:"productCode"`
	Source       string    `gorm:"column:source;size:50;comment:购买渠道" json:"source"`
	Brand        string    `gorm:"column:brand;size:100;comment:品牌" json:"brand"`
	Model        string    `gorm:"column:model;size:200;comment:厂家型号" json:"model"`
	Package      string    `gorm:"column:package;size:100;comment:封装" json:"package"`
	Name         string    `gorm:"column:name;size:200;comment:商品名称" json:"name"`
	Quantity     int       `gorm:"column:quantity;comment:订购数量" json:"quantity"`
	Price        float64   `gorm:"column:price;type:decimal(10,2);comment:商品金额" json:"price"`
	StockIn      int       `gorm:"column:stockIn;comment:入库数量" json:"stockIn"`
	StockOut     int       `gorm:"column:stockOut;comment:出库数量" json:"stockOut"`
	CurrentStock int       `gorm:"column:currentStock;comment:当前库存" json:"currentStock"`
	CreatedAt    time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type StockRecord struct {
	ID          uint      `gorm:"column:id;primarykey" json:"id"`
	ComponentID uint      `gorm:"column:componentId;comment:元件 ID" json:"componentId"`
	Type        string    `gorm:"column:type;size:20;comment:类型：in/out" json:"type"`
	Quantity    int       `gorm:"column:quantity;comment:数量" json:"quantity"`
	Remark      string    `gorm:"column:remark;size:500;comment:备注" json:"remark"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`
}

func (Component) TableName() string {
	return "components"
}

func (StockRecord) TableName() string {
	return "stockRecords"
}

type User struct {
	ID        uint      `gorm:"column:id;primarykey" json:"id"`
	Username  string    `gorm:"column:username;size:50;uniqueIndex;not null;comment:用户名" json:"username"`
	Password  string    `gorm:"column:password;size:255;not null;comment:密码" json:"-"`
	Email     string    `gorm:"column:email;size:100;comment:邮箱" json:"email"`
	Role      string    `gorm:"column:role;size:20;default:'user';comment:角色：admin/user" json:"role"`
	Status    int       `gorm:"column:status;default:1;comment:状态：1启用，0禁用" json:"status"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
