package globalInit

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	Db  *gorm.DB
	dsn string
)

func DbInit() {
	dsn = genDsn(viper.GetStringMapString("storage"))
	//程序启动打开数据库连接
	if err := initEngine(); err != nil {
		panic(err)
	}
}

//将数据库连接信息连接成字符串
func genDsn(storageConfig map[string]string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		storageConfig["user"],
		storageConfig["password"],
		storageConfig["host"],
		storageConfig["port"],
		storageConfig["dbname"],
		storageConfig["charset"])
}

func initEngine() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(viper.GetInt("gormLog.slowThreshold")) * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.LogLevel(viper.GetInt("gormLog.logLevel")),                  // 日志级别
			IgnoreRecordNotFoundError: true,                                                               // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                                                              // 禁用彩色打印
		},
	)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//默认日志
		// Logger: logger.Default.LogMode(logger.Silent),
		//默认关闭事务
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
		NamingStrategy: schema.NamingStrategy{
			//表前缀
			TablePrefix: viper.GetString("storage.prefix"),
			//表复数禁用
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := Db.DB()
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	//连接池设置
	viper.SetDefault("storage.max_idle", 2)
	viper.SetDefault("storage.max_conn", 10)
	maxIdle := viper.GetInt("storage.max_idle")
	maxConn := viper.GetInt("storage.max_conn")
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxConn)

	return nil
}

func Transaction() (tx *gorm.DB) {
	tx = Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}
