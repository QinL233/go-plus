package yaml

import "time"

type Yaml struct {
	Web struct {
		Prefix       string        `json:"prefix"`
		Port         int           `yaml:"port"`
		Mode         string        `yaml:"mode"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	} `yaml:"web"`
	Logstash struct {
		Tag string `yaml:"tag"`
		Url string `yaml:"url"`
	} `yaml:"logstash"`
	Orm struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
			Charset  string `yaml:"charset"`
		} `yaml:"mysql"`
		Elastic struct {
			Url string `yaml:"url"`
		} `yaml:"elastic"`
		Clickhouse struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"clickhouse"`
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
		} `yaml:"minio"`
	} `yaml:"oss"`
	Mq struct {
		RocketMq struct {
			Url      string `yaml:"url"`
			LogLevel string `yaml:"logLevel"`
		} `yaml:"rocketMq"`
	} `yaml:"mq"`
}
