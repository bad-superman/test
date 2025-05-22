package dao

import (
	"github.com/bad-superman/test/sdk/thegraph"
	"gorm.io/gorm/clause"
)

// UpsertIndexer 实现indexer的upsert操作，主键为id
func (d *Dao) UpsertIndexer(indexer *thegraph.Indexer) error {
	return d.myClient.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // 以id为唯一键
		UpdateAll: true,                          // 冲突时更新所有字段
	}).Create(indexer).Error
}
