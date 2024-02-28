package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/*
*
通过id查询one - 单表
*/
func TryOneKey[T any](db *gorm.DB, key any) T {
	defer func() {
		if err := recover(); err != nil && !errors.Is(err.(error), gorm.ErrRecordNotFound) {
			panic(err)
		}
	}()
	return OneKey[T](db, key)
}

func OneKey[T any](db *gorm.DB, key any) T {
	return OneKeyTo[T, T](db, key)
}

func OneKeyTo[T any, E any](db *gorm.DB, key any) E {
	return OneKeyFieldTo[T, E](db, nil, key)
}

func OneKeyFieldTo[T any, E any](db *gorm.DB, field []string, key any) E {
	var result E
	var entity T
	query := db.Model(&entity)
	if len(field) > 0 {
		query.Select(field)
	}
	//根据id查询时使用first默认以id排序以提升性能
	if err := query.First(&result, key).Error; err != nil {
		panic(err)
	}
	return result
}

/*
*
通过指定参数查询one - 单表
*/
func TryOne[T any](db *gorm.DB, condition interface{}, args ...interface{}) T {
	defer func() {
		if err := recover(); err != nil && !errors.Is(err.(error), gorm.ErrRecordNotFound) {
			panic(err)
		}
	}()
	return One[T](db, condition, args...)
}

func One[T any](db *gorm.DB, condition interface{}, args ...interface{}) T {
	return OneTo[T, T](db, condition, args...)
}

func OneTo[T any, E any](db *gorm.DB, condition interface{}, args ...interface{}) E {
	return OneFieldTo[T, E](db, nil, condition, args...)
}

func OneFieldTo[T any, E any](db *gorm.DB, field []string, condition interface{}, args ...interface{}) E {
	var result E
	var entity T
	query := db.Model(&entity).Where(condition, args...)
	if len(field) > 0 {
		query.Select(field)
	}
	if err := query.Take(&result).Error; err != nil {
		panic(err)
	}
	return result
}

/*
*
通过对象查询one - 单表
*/
func TryOneEntity[T any](db *gorm.DB, entity T) T {
	defer func() {
		if err := recover(); err != nil && !errors.Is(err.(error), gorm.ErrRecordNotFound) {
			panic(err)
		}
	}()
	return OneEntity[T](db, entity)
}

func OneEntity[T any](db *gorm.DB, entity T) T {
	return OneEntityTo[T, T](db, entity)
}

func OneEntityTo[T any, E any](db *gorm.DB, entity T) E {
	return OneEntityFieldTo[T, E](db, nil, entity)
}

func OneEntityFieldTo[T any, E any](db *gorm.DB, field []string, entity T) E {
	var result E
	query := db.Model(&entity).Where(&entity)
	if len(field) > 0 {
		query.Select(field)
	}
	if err := query.Take(&result).Error; err != nil {
		panic(err)
	}
	return result
}

/*
*
自定义模式查询one - 满足多表组合等复杂查询
*/
func TryOneScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) T {
	defer func() {
		if err := recover(); err != nil && !errors.Is(err.(error), gorm.ErrRecordNotFound) {
			panic(err)
		}
	}()
	return OneScope[T](db, scope)
}

func OneScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) T {
	return OneScopeTo[T, T](db, scope)
}

func OneScopeTo[T any, E any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) E {
	var result E
	var entity T
	query := db.Model(&entity).Scopes(scope)
	if err := query.Take(&result).Error; err != nil {
		panic(err)
	}
	return result
}
