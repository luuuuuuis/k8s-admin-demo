package db

import (
	"fmt"
	"k8s-server/config"
	"time"

	"k8s-server/utils"

	"github.com/jinzhu/gorm"                  //gorm库
	_ "github.com/jinzhu/gorm/dialects/mysql" //gorm对应的mysql驱动
)

var (
	isInit bool
	GORM   *gorm.DB
	err    error
)

// db的初始化函数，与数据库建立连接
func Init() {
	//判断是否已经初始化了
	if isInit {
		return
	}
	//组装连接配置
	//parseTime是查询结果是否自动解析为时间
	//loc是Mysql的时区设置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.GetString("DB.DbUser"),
		config.Config.GetString("DB.DbPwd"),
		config.Config.GetString("DB.DbHost"),
		config.Config.GetInt("DB.DbPort"),
		config.Config.GetString("DB.DbName"),
	)
	//与数据库建立连接，生成一个*gorm.DB类型的对象
	GORM, err = gorm.Open(config.Config.GetString("DB.DbType"), dsn)
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	//打印sql语句
	GORM.LogMode(config.Config.GetBool("DB.LogMode"))

	//开启连接池
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	GORM.DB().SetMaxIdleConns(config.Config.GetInt("DB.MaxIdleConns"))
	// 连接池最大允许的连接数
	GORM.DB().SetMaxOpenConns(config.Config.GetInt("DB.MaxOpenConns"))
	// 设置了连接可复用的最大时间
	GORM.DB().SetConnMaxLifetime(time.Duration(config.Config.GetInt("DB.MaxLifeTime")) * time.Second)

	isInit = true
	utils.Logger.Info().Msg("连接数据库成功!")
}

// db的关闭函数
func Close() error {
	return GORM.Close()
}
