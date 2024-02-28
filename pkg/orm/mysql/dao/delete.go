package dao

import "gorm.io/gorm"

func RemoveKey[T any](db *gorm.DB, key any) bool {
	var entity T
	//invalid value, should be pointer to struct or slice
	if err := db.Delete(&entity, key).Error; err != nil {
		panic(err)
	}
	return true
}

func Remove[T any](db *gorm.DB, condition interface{}, args ...interface{}) bool {
	var entity T
	if err := db.Where(condition, args...).Delete(&entity).Error; err != nil {
		panic(err)
	}
	return true
}

func RemoveEntity[T any](db *gorm.DB, entity T) bool {
	if err := db.Where(&entity).Delete(&entity).Error; err != nil {
		panic(err)
	}
	return true
}

func RemoveScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) bool {
	var entity T
	if err := db.Scopes(scope).Delete(&entity).Error; err != nil {
		panic(err)
	}
	return true
}
