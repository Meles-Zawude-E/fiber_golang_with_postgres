package database

import (
	"fmt"
	"strconv"

	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/config"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}
var DB DBinstance
func Connect() {
   p:=config.Config("DB_PORT")
   port, err:=strconv.ParseInt(p, 10, 32)
   if err != nil {
	fmt.Print("Not parsed to string")
   }
   dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",config.Config("DB_HOST"),config.Config("DB_USER"),config.Config("DB_PASSWORD"),config.Config("DB_NAME"), port)
   db, err:=gorm.Open(postgres.Open(dsn), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
   })
   if err!=nil {
	fmt.Print("Failed to connect database")
   }
   db.AutoMigrate(&model.Employee{})
   DB=DBinstance{
	Db: db,
   }
}