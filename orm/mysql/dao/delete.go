package dao

import "gorm.io/gorm"

func RemoveKey[T any](db *gorm.DB, key any) (bool, error) {
	var entity T
	//invalid value, should be pointer to struct or slice
	if err := db.Delete(&entity, key).Error; err != nil {
		return false, err
	}
	return true, nil
}

func Remove[T any](db *gorm.DB, condition interface{}, args ...interface{}) (bool, error) {
	var entity T
	if err := db.Where(condition, args...).Delete(&entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func RemoveEntity[T any](db *gorm.DB, entity T) (bool, error) {
	if err := db.Where(&entity).Delete(&entity).Error; err != nil {
		return false, err
	}
	return true, nil
}

func RemoveScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) (bool, error) {
	var entity T
	if err := db.Scopes(scope).Delete(&entity).Error; err != nil {
		return false, err
	}
	return true, nil
}
