package dao

import "gorm.io/gorm"

// InitTables 纯纯的垃圾设计
func InitTables(db *gorm.DB) error {
	// 建表
	return db.AutoMigrate(&User{})
}
