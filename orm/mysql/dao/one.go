package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/**
通过id查询one - 单表
*/
func OneKeyTry[T any](db *gorm.DB, key any) (T, error) {
	if result, err := OneKey[T](db, key); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	} else {
		//注意泛型返回时必定会返回一个属性为zero的结构，即使用时执行通过属性判断
		return result, nil
	}
}

func OneKey[T any](db *gorm.DB, key any) (T, error) {
	return OneKeyTo[T, T](db, key)
}

func OneKeyTo[T any, E any](db *gorm.DB, key any) (E, error) {
	return OneKeyFieldTo[T, E](db, nil, key)
}

func OneKeyFieldTo[T any, E any](db *gorm.DB, field []string, key any) (E, error) {
	var result E
	var entity T
	query := db.Model(&entity)
	if len(field) > 0 {
		query.Select(field)
	}
	//根据id查询时使用first默认以id排序以提升性能
	if err := query.First(&result, key).Error; err != nil {
		return result, err
	}
	return result, nil
}

/**
通过指定参数查询one - 单表
*/
func OneTry[T any](db *gorm.DB, condition interface{}, args ...interface{}) (T, error) {
	if result, err := One[T](db, condition, args...); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	} else {
		//注意泛型返回时必定会返回一个属性为zero的结构，即使用时执行通过属性判断
		return result, nil
	}
}

func One[T any](db *gorm.DB, condition interface{}, args ...interface{}) (T, error) {
	return OneTo[T, T](db, condition, args...)
}

func OneTo[T any, E any](db *gorm.DB, condition interface{}, args ...interface{}) (E, error) {
	return OneFieldTo[T, E](db, nil, condition, args...)
}

func OneFieldTo[T any, E any](db *gorm.DB, field []string, condition interface{}, args ...interface{}) (E, error) {
	var result E
	var entity T
	query := db.Model(&entity).Where(condition, args...)
	if len(field) > 0 {
		query.Select(field)
	}
	if err := query.Take(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

/**
通过对象查询one - 单表
*/
func OneEntityTry[T any](db *gorm.DB, entity T) (T, error) {
	if result, err := OneEntity[T](db, entity); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	} else {
		//注意泛型返回时必定会返回一个属性为zero的结构，即使用时执行通过属性判断
		return result, nil
	}
}

func OneEntity[T any](db *gorm.DB, entity T) (T, error) {
	return OneEntityTo[T, T](db, entity)
}

func OneEntityTo[T any, E any](db *gorm.DB, entity T) (E, error) {
	return OneEntityFieldTo[T, E](db, nil, entity)
}

func OneEntityFieldTo[T any, E any](db *gorm.DB, field []string, entity T) (E, error) {
	var result E
	query := db.Model(&entity).Where(&entity)
	if len(field) > 0 {
		query.Select(field)
	}
	if err := query.Take(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

/**
自定义模式查询one - 满足多表组合等复杂查询
*/
func OneScopeTry[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) (T, error) {
	if result, err := OneScope[T](db, scope); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	} else {
		//注意泛型返回时必定会返回一个属性为zero的结构，即使用时执行通过属性判断
		return result, nil
	}
}

func OneScope[T any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) (T, error) {
	return OneScopeTo[T, T](db, scope)
}

func OneScopeTo[T any, E any](db *gorm.DB, scope func(db *gorm.DB) *gorm.DB) (E, error) {
	var result E
	var entity T
	query := db.Model(&entity).Scopes(scope)
	if err := query.Take(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
