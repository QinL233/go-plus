package web

/**
定义service通用接口
1、实现用户继承该接口并实现方法后，controller能直接快速实现param校验和返回封装
2、实现对service的抽象
*/

type Service interface {
	//常规业务函数，返回data
	handler(service Service) (any, error)
	// Exec 需要重载的接口
	Exec() (any, error)
}

/**
基础service【可拓展】
*/

type BaseService struct {
}

func (s *BaseService) handler(service Service) (any, error) {
	return service.Exec()
}

func (s *BaseService) Exec() (any, error) {
	return nil, nil
}
