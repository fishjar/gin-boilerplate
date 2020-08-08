/*
	db 数据库连接
*/

package db

import (
	"fmt"
	"time"

	"github.com/8treenet/gcache"
	"github.com/8treenet/gcache/option"
	"github.com/fishjar/gin-boilerplate/config"

	"github.com/fishjar/gin-boilerplate/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // 引入mysql驱动
	_ "github.com/jinzhu/gorm/dialects/sqlite" // 引入sqlite驱动
)

// DB 为ORM全局实例
var DB *gorm.DB

func init() {
	// 创建数据库连接
	db, err := gorm.Open(config.Config.DBDriver, config.Config.DBPath)
	if err != nil {
		fmt.Println("打开数据库错误：", err.Error())
		panic("连接数据库失败")
	}
	if config.Config.APPEnv == "development" {
		db.LogMode(true) // 输出SQL日志
	}

	// 测试数据库连接
	if db.DB() == nil { // 如果数据库底层连接的不是一个 *sql.DB，那么该方法会返回 nil
		fmt.Println("获取数据库接口错误")
		panic("连接数据库失败")
	}
	if err := db.DB().Ping(); err != nil {
		fmt.Println("Ping数据库错误：", err.Error())
		panic("连接数据库失败")
	}

	// db设置
	db.DB().SetMaxIdleConns(10)           // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(100)          // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(time.Hour) // 设置连接的最大可复用时间
	db.SetLogger(logger.LogGorm)          // TODO:log设置
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0)) // TODO:log设置

	// 缓存设置
	opt := option.DefaultOption{}
	opt.Expires = 300             //缓存时间，默认120秒。范围 30-43200
	opt.Level = option.LevelModel //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = true         //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存。 ps: affected如果未0，不触发更新。
	opt.PenetrationSafe = true    //开启防穿透, 默认false。 ps:防击穿强制全局开启。

	//缓存中间件附加到gorm.DB
	gcache.AttachDB(db, &opt, &option.RedisOption{
		Addr:     config.Config.Redis.Addr,     // redis 地址
		Password: config.Config.Redis.Password, // redis 密码
		DB:       config.Config.Redis.Name,     // use default DB
	})

	DB = db
}
