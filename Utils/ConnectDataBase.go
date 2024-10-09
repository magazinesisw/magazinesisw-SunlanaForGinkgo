package Utils

import (
	"Ginko/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() (*gorm.DB, error) {

	//dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	var database config.Database
	DataBaseConfig := config.DatabaseConfig{}
	err := ReadYML("..\\..\\config\\DataBase.yaml", &DataBaseConfig)

	if err != nil {
		panic(err)
	}
	database = DataBaseConfig.Database
	dsn := database.Username + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err1 != nil {

		panic("链接数据库失败" + err1.Error())
	}

	DB = db

	return db, nil

}
