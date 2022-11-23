package repository

import (
	"github.com/thebluefowl/zion/model"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) model.ChannelRepository {
	return &ChannelRepository{db: db}
}

func (r *ChannelRepository) Create(channel *model.Channel) error {
	return r.db.Create(channel).Error
}

func (r *ChannelRepository) Get(tenantID, id string) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.First(&channel, model.Channel{TenantID: tenantID, ID: id}).Error
	return &channel, err
}

func (r *ChannelRepository) Filter(tenantID string, name string) ([]model.Channel, error) {
	channels := []model.Channel{}
	filter := model.Channel{TenantID: tenantID}
	if name != "" {
		filter.Name = name
	}
	err := r.db.Find(&channels, model.Channel{TenantID: tenantID, Name: name}).Error
	return channels, err
}

func (r *ChannelRepository) Delete(tenantID, id string) error {
	return r.db.Where("tenant_id = ? AND id = ?", tenantID, id).Delete(&model.Channel{}).Error
}
