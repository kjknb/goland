package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}
func Initmysql() {
	// 自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)

	// 从配置获取DSN
	dsn := viper.GetString("mysql.dns")
	if dsn == "" {
		log.Fatal("MySQL DSN配置为空，请检查配置文件")
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 获取底层SQL DB实例以配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("获取数据库实例失败:", err)
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接的最大可复用时间

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("数据库连接测试失败:", err)
	}

	fmt.Println("MySQL初始化成功")

	// 可选：自动迁移模型
	// DB.AutoMigrate(&models.UserBasic{})

	// 示例查询
	// user := models.UserBasic{}
	// result := DB.First(&user)
	// if result.Error != nil {
	//     fmt.Println("查询失败:", result.Error)
	// } else {
	//     fmt.Println("查询成功:", user)
	// }
}
