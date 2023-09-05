package yaml

import "time"

type Yaml struct {
	Gin struct {
		Prefix       string        `json:"prefix"`
		Port         int           `yaml:"port"`
		Mode         string        `yaml:"mode"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
		Log          struct {
			Filters []string `yaml:"filters"`
			Client  struct {
				Logstash struct {
					Url string `yaml:"url"`
				} `yaml:"logstash"`
				Mysql struct {
					Host     string `yaml:"host"`
					Port     int    `yaml:"port"`
					Username string `yaml:"username"`
					Password string `yaml:"password"`
					Database string `yaml:"database"`
				} `yaml:"mysql"`
			} `yaml:"client"`
		} `yaml:"log"`
	} `yaml:"gin"`
	Orm struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`
		Elastic struct {
			Url string `yaml:"url"`
		} `yaml:"elastic"`
	} `yaml:"orm"`
	Cache struct {
		Redis struct {
			Host        string        `yaml:"host"`
			Port        int           `yaml:"port"`
			Password    string        `yaml:"password"`
			Database    int           `yaml:"database"`
			MaxIdle     int           `yaml:"maxIdle"`
			MaxActive   int           `yaml:"maxActive"`
			IdleTimeout time.Duration `yaml:"idleTimeout"`
		} `yaml:"redis"`
	} `yaml:"cache"`
	Oss struct {
		Minio struct {
			Endpoint  string `yaml:"endpoint"`
			Port      int    `yaml:"port"`
			AccessKey string `yaml:"accessKey"`
			SecretKey string `yaml:"secretKey"`
			Secure    bool   `yaml:"secure"`
			Bucket    string `yaml:"bucket"`
			Prefix    string `json:"prefix"`
		} `yaml:"minio"`
	} `yaml:"oss"`
	Mq struct {
		RocketMq struct {
			Url      string `yaml:"url"`
			LogLevel string `yaml:"logLevel"`
		} `yaml:"rocketMq"`
	} `yaml:"mq"`
}
