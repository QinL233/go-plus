package dao

import (
	"gorm.io/gorm"
)

// 分页scope:每页显示数量size,页数page
func Paginate(size int, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//根据分页计算出偏移量
		if size <= 0 {
			size = 10
		}
		if page <= 0 {
			page = 1
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

/*
*
指定参数进行分页查询 - 单表
*/
func Page[T any](db *gorm.DB, size, page int, condition interface{}, args ...interface{}) (int64, []T) {
	return PageSort[T](db, "", size, page, condition, args...)
}

func PageSort[T any](db *gorm.DB, sort string, size, page int, condition interface{}, args ...interface{}) (int64, []T) {
	return PageSortTo[T, T](db, sort, size, page, condition, args...)
}

func PageSortTo[T any, E any](db *gorm.DB, sort string, size, page int, condition interface{}, args ...interface{}) (int64, []E) {
	return PageSortFieldTo[T, E](db, nil, sort, size, page, condition, args...)
}

func PageSortFieldTo[T any, E any](db *gorm.DB, field []string, sort string, size, page int, condition interface{}, args ...interface{}) (int64, []E) {
	var result []E
	count := Count[T](db, condition, args...)
	if count > 0 {
		var entity T
		query := db.Model(&entity).Where(condition, args...).Scopes(Paginate(size, page))
		if len(field) > 0 {
			query.Select(field)
		}
		if sort != "" {
			query.Order(sort)
		}
		if err := query.Find(&result).Error; err != nil {
			panic(err)
		}
	}
	return count, result
}

/*
*
指定对象进行分页查询 - 单表
*/
func PageEntity[T any](db *gorm.DB, size, page int, entity T) (int64, []T) {
	return PageEntitySort[T](db, "", size, page, entity)
}

func PageEntitySort[T any](db *gorm.DB, sort string, size, page int, entity T) (int64, []T) {
	return PageEntitySortTo[T, T](db, sort, size, page, entity)
}

func PageEntitySortTo[T any, E any](db *gorm.DB, sort string, size, page int, entity T) (int64, []E) {
	return PageEntitySortFieldTo[T, E](db, nil, sort, size, page, entity)
}

func PageEntitySortFieldTo[T any, E any](db *gorm.DB, field []string, sort string, size, page int, entity T) (int64, []E) {
	var result []E
	count := CountEntity(db, entity)
	if count > 0 {
		query := db.Model(&entity).Where(entity).Scopes(Paginate(size, page))
		if len(field) > 0 {
			query.Select(field)
		}
		if sort != "" {
			query.Order(sort)
		}
		if err := query.Find(&result).Error; err != nil {
			panic(err)
		}
	}
	return count, result
}

/*
*
自定义scope进行分页查询 - 用于组合连表等复杂查询
*/
func PageScope[T any](db *gorm.DB, size, page int, scope func(db *gorm.DB) *gorm.DB) (int64, []T) {
	return PageScopeFieldTo[T, T](db, nil, size, page, scope)
}

func PageScopeFieldTo[T any, E any](db *gorm.DB, field []string, size, page int, scope func(db *gorm.DB) *gorm.DB) (int64, []E) {
	var result []E
	//注意 scope 查询不得存在select
	count := CountScope[T](db, scope)
	if count > 0 {
		var entity T
		query := db.Model(&entity).Scopes(Paginate(size, page), scope)
		if len(field) > 0 {
			query.Select(field)
		}
		if err := query.Find(&result).Error; err != nil {
			panic(err)
		}
	}
	return count, result
}
