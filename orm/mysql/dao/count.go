package dao

import "gorm.io/gorm"

func Count[T any](db *gorm.DB, condition interface{}, args ...interface{}) (int64, error) {
	var count int64
	var entity T
	if err := db.Model(&entity).Where(condition, args...).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func CountEntity[T any](db *gorm.DB, entity T) (int64, error) {
	var count int64
	if err := db.Model(&entity).Where(&entity).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func CountScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	var entity T
	//注意此处scope不允许写入select field...
	//Scan error on column index 0, name "username": converting driver.Value type []uint8 ("root") to a int64: invalid syntax
	if err := db.Model(&entity).Scopes(scope).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
