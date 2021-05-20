package dao

import (
	"io/ioutil"
	"log"

	"baselayout/internal/user/model"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open() *gorm.DB {

	conf := Config()

	var db *gorm.DB
	var err error

	dsn := "user=" + conf.Username + " password= " + conf.Password + " dbname= " + conf.Database + " host= " + conf.Hostname + " port=" + conf.Port + " sslmode=disable TimeZone=Asia/Taipei"
	if gin.Mode() == gin.ReleaseMode {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	}

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.User{})

	return db

}

type Postgresql struct {
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

func Config() (conf *Postgresql) {
	conf = new(Postgresql)
	yamlFile, _ := ioutil.ReadFile("internal/postgresql.yaml")

	errUn := yaml.Unmarshal(yamlFile, conf)
	if errUn != nil {
		log.Fatalf("Unmarshal:", errUn)
	}
	return conf
}
