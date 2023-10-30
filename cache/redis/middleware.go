package redis

import "fmt"

// CacheHandler 缓存处理函数
func CacheHandler[R any](prefix string, key string, second int, handler func() R) R {
	cache := Client{Prefix: prefix}
	if cache.Exists(key) {
		var r R
		err := cache.GetObject(key, &r)
		if err != nil {
			fmt.Println(err)
		} else {
			return r
		}
	}
	r := handler()
	cache.Set(key, r, second)
	return r
}
