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

// 全局变量存储JWT配置
var (
	JwtSecret string
	JwtExpire int
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败:", err)
		// 尝试从环境变量获取配置
		setupFromEnv()
		return
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))

	// 读取JWT配置
	JwtSecret = viper.GetString("jwt.secret")
	JwtExpire = viper.GetInt("jwt.expire")

	// 如果没有配置JWT密钥，尝试从环境变量获取或使用默认值
	if JwtSecret == "" {
		JwtSecret = os.Getenv("GINCHAT_JWT_SECRET")
		if JwtSecret == "" {
			// 如果环境变量也没有设置，使用一个默认值（仅用于开发）
			JwtSecret = "fallback-dev-secret-key-change-in-production"
			fmt.Println("警告: 使用默认JWT密钥，生产环境请设置jwt.secret配置或GINCHAT_JWT_SECRET环境变量")
		}
	}

	if JwtExpire == 0 {
		expireStr := os.Getenv("GINCHAT_JWT_EXPIRE")
		if expireStr == "" {
			JwtExpire = 24 // 默认24小时
		} else {
			fmt.Sscanf(expireStr, "%d", &JwtExpire)
		}
	}

	// 检查JWT密钥长度，建议至少32字符
	if len(JwtSecret) < 32 {
		fmt.Println("警告: JWT密钥长度不足，建议使用至少32字符的密钥")
	}

	fmt.Println("JWT配置加载成功 - 有效期:", JwtExpire, "小时")
}

// 从环境变量获取配置的备用方法
func setupFromEnv() {
	fmt.Println("尝试从环境变量获取配置...")

	// 从环境变量获取MySQL配置
	mysqlDns := os.Getenv("GINCHAT_MYSQL_DNS")
	if mysqlDns == "" {
		mysqlDns = "root:123mdx0.0@tcp(127.0.0.1:3306)/GinChat1?charset=utf8mb4&parseTime=True&loc=Local"
		fmt.Println("使用默认MySQL连接字符串")
	}
	viper.Set("mysql.dns", mysqlDns)

	// 从环境变量获取JWT配置
	JwtSecret = os.Getenv("GINCHAT_JWT_SECRET")
	if JwtSecret == "" {
		JwtSecret = "fallback-dev-secret-key-change-in-production"
		fmt.Println("警告: 使用默认JWT密钥，生产环境请设置GINCHAT_JWT_SECRET环境变量")
	}

	expireStr := os.Getenv("GINCHAT_JWT_EXPIRE")
	if expireStr == "" {
		JwtExpire = 24 // 默认24小时
	} else {
		fmt.Sscanf(expireStr, "%d", &JwtExpire)
	}
}

func Initmysql() {
	// 从配置获取DSN
	dsn := viper.GetString("mysql.dns")
	if dsn == "" {
		log.Fatal("MySQL DSN配置为空，请检查配置文件")
	}

	// 自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)

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
}
