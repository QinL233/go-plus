package dao

import "gorm.io/gorm"

func Create(db *gorm.DB, entity any) int64 {
	r := db.Create(entity)
	if r.Error != nil {
		panic(r.Error)
	}
	return r.RowsAffected
}

func CreateBatch(db *gorm.DB, entity any, size int) int64 {
	r := db.CreateInBatches(entity, size)
	if r.Error != nil {
		panic(r.Error)
	}
	return r.RowsAffected
}
