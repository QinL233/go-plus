package dao

import "gorm.io/gorm"

//更新单列
func Update[T any](db *gorm.DB, column string, value interface{}, condition interface{}, args ...interface{}) int64 {
	var entity T
	statement := db.Model(&entity).Where(condition, args).Update(column, value)
	if err := statement.Error; err != nil {
		panic(err)
	}
	return statement.RowsAffected
}

func UpdateEntity[T any](db *gorm.DB, column string, value interface{}, entity T) int64 {
	statement := db.Model(&entity).Where(&entity).Update(column, value)
	if err := statement.Error; err != nil {
		panic(err)
	}
	return statement.RowsAffected
}

func UpdateScope[T any](db *gorm.DB, column string, value interface{}, scope func(db *gorm.DB) *gorm.DB) int64 {
	var entity T
	statement := db.Model(&entity).Scopes(scope).Update(column, value)
	if err := statement.Error; err != nil {
		panic(err)
	}
	return statement.RowsAffected
}

//更新多列
func Updates[T any](db *gorm.DB, value map[string]interface{}, condition interface{}, args ...interface{}) int64 {
	var entity T
	statement := db.Model(&entity).Where(condition, args).Updates(value)
	if err := statement.Error; err != nil {
		panic(err)
	}
	return statement.RowsAffected
}

func UpdatesEntity[T any](db *gorm.DB, value map[string]interface{}, entity T) int64 {
	statement := db.Model(&entity).Where(&entity).Updates(value)
	if err := statement.Error; err != nil {
		panic(err)
	}
	return statement.RowsAffected
}

func UpdatesScope[T any](db *gorm.DB, value map[string]interface{}, scope func(db *gorm.DB) *gorm.DB) int64 {
	var entity T
	statement := db.Model(&entity).Scopes(scope).Updates(value)
	if err := statement.Error; err != nil {
		panic(err)
	}
	return statement.RowsAffected
}
