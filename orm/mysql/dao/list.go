package dao

import "gorm.io/gorm"

/**
指定参数进行查询 - 单表
*/
func List[T any](db *gorm.DB, condition interface{}, args ...interface{}) []T {
	return ListSort[T](db, "", condition, args...)
}

func ListSort[T any](db *gorm.DB, sort string, condition interface{}, args ...interface{}) []T {
	return ListSortTo[T, T](db, sort, condition, args...)
}

func ListSortTo[T any, E any](db *gorm.DB, sort string, condition interface{}, args ...interface{}) []E {
	return ListSortLimitTo[T, E](db, sort, 0, condition, args...)
}

func ListSortLimitTo[T any, E any](db *gorm.DB, sort string, limit int, condition interface{}, args ...interface{}) []E {
	return ListSortLimitFieldTo[T, E](db, nil, sort, limit, condition, args...)
}

func ListSortLimitFieldTo[T any, E any](db *gorm.DB, field []string, sort string, limit int, condition interface{}, args ...interface{}) []E {
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
		panic(err)
	}
	return result
}

/**
指定对象进行查询 - 单表
*/
func ListEntity[T any](db *gorm.DB, entity T) []T {
	return ListEntitySort[T](db, "", entity)
}

func ListEntitySort[T any](db *gorm.DB, sort string, entity T) []T {
	return ListEntitySortTo[T, T](db, sort, entity)
}

func ListEntitySortTo[T any, E any](db *gorm.DB, sort string, entity T) []E {
	return ListEntitySortLimitTo[T, E](db, sort, 0, entity)
}

func ListEntitySortLimitTo[T any, E any](db *gorm.DB, sort string, limit int, entity T) []E {
	return ListEntitySortLimitFieldTo[T, E](db, nil, sort, limit, entity)
}

func ListEntitySortLimitFieldTo[T any, E any](db *gorm.DB, field []string, sort string, limit int, entity T) []E {
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
		panic(err)
	}
	return result
}

/**
自定义scope进行查询 - 用于组合连表等复杂查询
*/
func ListScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) []T {
	return ListScopeTo[T, T](db, scope)
}

func ListScopeTo[T any, E any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) []E {
	var result []E
	var entity T
	query := db.Model(&entity).Scopes(scope)
	if err := query.Find(&result).Error; err != nil {
		panic(err)
	}
	return result
}
