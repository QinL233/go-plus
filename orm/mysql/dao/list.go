package dao

import "gorm.io/gorm"

/**
指定参数进行查询 - 单表
*/
func List[T any](db *gorm.DB, condition interface{}, args ...interface{}) ([]T, error) {
	return ListSort[T](db, "", condition, args...)
}

func ListSort[T any](db *gorm.DB, sort string, condition interface{}, args ...interface{}) ([]T, error) {
	return ListSortTo[T, T](db, sort, condition, args...)
}

func ListSortTo[T any, E any](db *gorm.DB, sort string, condition interface{}, args ...interface{}) ([]E, error) {
	return ListSortLimitTo[T, E](db, sort, 0, condition, args...)
}

func ListSortLimitTo[T any, E any](db *gorm.DB, sort string, limit int, condition interface{}, args ...interface{}) ([]E, error) {
	return ListSortLimitFieldTo[T, E](db, nil, sort, limit, condition, args...)
}

func ListSortLimitFieldTo[T any, E any](db *gorm.DB, field []string, sort string, limit int, condition interface{}, args ...interface{}) ([]E, error) {
	var result []E
	var entity T
	query := db.Model(&entity).Where(condition, args...)
	if len(field) > 0 {
		query.Select(field)
	}
	if sort != "" {
		query.Order(sort)
	}
	if limit > 0 {
		query.Limit(limit)
	}
	if err := query.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

/**
指定对象进行查询 - 单表
*/
func ListEntity[T any](db *gorm.DB, entity T) ([]T, error) {
	return ListEntitySort[T](db, "", entity)
}

func ListEntitySort[T any](db *gorm.DB, sort string, entity T) ([]T, error) {
	return ListEntitySortTo[T, T](db, sort, entity)
}

func ListEntitySortTo[T any, E any](db *gorm.DB, sort string, entity T) ([]E, error) {
	return ListEntitySortLimitTo[T, E](db, sort, 0, entity)
}

func ListEntitySortLimitTo[T any, E any](db *gorm.DB, sort string, limit int, entity T) ([]E, error) {
	return ListEntitySortLimitFieldTo[T, E](db, nil, sort, limit, entity)
}

func ListEntitySortLimitFieldTo[T any, E any](db *gorm.DB, field []string, sort string, limit int, entity T) ([]E, error) {
	var result []E
	query := db.Model(&entity).Where(&entity)
	if len(field) > 0 {
		query.Select(field)
	}
	if sort != "" {
		query.Order(sort)
	}
	if limit > 0 {
		query.Limit(limit)
	}
	if err := query.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

/**
自定义scope进行查询 - 用于组合连表等复杂查询
*/
func ListScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) ([]T, error) {
	return ListScopeTo[T, T](db, scope)
}

func ListScopeTo[T any, E any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) ([]E, error) {
	var result []E
	var entity T
	query := db.Model(&entity).Scopes(scope)
	if err := query.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
