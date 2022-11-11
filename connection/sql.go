package connection

import (
	"github.com/jinzhu/gorm"
)

type StoreConnection struct {
	db *gorm.DB
}

func NewStoreConnection(db *gorm.DB) *StoreConnection {
	return &StoreConnection{db: db}
}

func (sc *StoreConnection) Get(id, tenantID string) *Connection {
	var c Connection
	sc.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&c)
	return &c
}

func (sc *StoreConnection) Create(c *Connection) error {
	return sc.db.Create(c).Error
}
