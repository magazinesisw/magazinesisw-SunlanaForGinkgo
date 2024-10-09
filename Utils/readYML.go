package Utils

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func ReadYML(filePath string, target interface{}) error {
	//打开yaml
	file, err := os.Open(filePath)
	if err != nil {

		log.Println("打开文件失败：", err)
		return err

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("关闭文件失败：", err)
		}
	}(file)

	//解析yaml
	out := yaml.NewDecoder(file).Decode(target)

	return out

}
