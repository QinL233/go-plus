package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Config Yaml

var PathList = []string{
	"app.yaml",
	"app.yml",
	"config/app.yaml",
	"config/app.yml",
	"conf/app.yaml",
	"conf/app.yml",
}

/**
 * 加载配置文件
 * 文件路径按：传参 > app.yaml&yml > config/app.yaml&yml > conf/app.yaml&yml
 */

func Init(conf ...string) {
	if len(conf) > 0 {
		path := conf[0]
		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("load config err %v", err)
		}
		if err = yaml.Unmarshal(fileBytes, &Config); err != nil {
			log.Fatalf("load config err %v", err)
		}
		log.Println(fmt.Sprintf("load config path [%db]%s", len(fileBytes), path))
	} else {
		for _, path := range PathList {
			fileBytes, err := ioutil.ReadFile(path)
			if err != nil {
				continue
			}
			if len(fileBytes) > 0 {
				if err = yaml.Unmarshal(fileBytes, Config); err != nil {
					log.Fatalf("load config err %v", err)
				}
				log.Println(fmt.Sprintf("load config path [%db]%s", len(fileBytes), path))
				break
			}
		}
	}
}
