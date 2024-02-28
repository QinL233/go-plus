package dao

import "gorm.io/gorm"

func Save(db *gorm.DB, entity any) int64 {
	r := db.Save(entity)
	if r.Error != nil {
		panic(r.Error)
	}
	return r.RowsAffected
}
