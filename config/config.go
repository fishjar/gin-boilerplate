/*
	配置文件
*/

package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/jinzhu/configor"
)

// Config 项目配置
var Config = struct {
	APPName          string        `default:"app name"`                       // 项目名称
	APPEnv           string        `default:"development" env:"CONFIGOR_ENV"` // 运行环境
	HTTPPort         uint          `default:"4000"`                           // 运行端口
	JWTSignKey       string        `default:"123456"`                         // JWT加密用的密钥
	JWTExpiresMinute time.Duration `default:"60"`                             // JWT过期时间(分钟)
	PasswordSalt     string        `default:"123456"`                         // 加密密钥
	SwaggerName      string        `default:"admin"`                          // swagger 用户名
	SwaggerPassword  string        `default:"123456"`                         // swagger 密码
	FilesPath        string        `default:"tmp"`                            // 临时文件目录
	UploadPath       string        `default:"upload"`                         // 上传文件目录
	LogPath          string        `default:"log"`                            // 日志文件目录

	RootPath       string        // 项目根目录
	FilesFullPath  string        // 临时文件完整目录
	UploadFullPath string        // 上传文件完整目录
	LogFullPath    string        // 日志文件完整目录
	JWTExpiresIn   time.Duration // JWT过期时间
	DBDriver       string        // 数据库驱动
	DBPath         string        // 数据库地址

	MySQL struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"3306"`
		User     string `default:"root"`
		Password string `default:"123456"`
		Name     string `default:"testdb"`
	}

	Redis struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"6379"`
		Password string
		Name     int `default:"0"`
		Addr     string
	}

	SQLite struct {
		FilePath string `default:"db"`
		FileName string `default:"sqlite.db"`
		FullPath string
	}
}{}

func init() {
	Config.RootPath, _ = os.Getwd()
	configor.Load(&Config, path.Join(Config.RootPath, "config", "config.yml"))
	// fmt.Printf("config: %#v", Config)
}

func init() {
	Config.FilesFullPath = path.Join(Config.RootPath, Config.FilesPath)
	Config.UploadFullPath = path.Join(Config.FilesFullPath, Config.UploadPath)
	Config.LogFullPath = path.Join(Config.FilesFullPath, Config.LogPath)
	Config.SQLite.FullPath = path.Join(Config.FilesFullPath, Config.SQLite.FilePath)

	Config.JWTExpiresIn = Config.JWTExpiresMinute * time.Minute
	Config.Redis.Addr = Config.Redis.Host + ":" + strconv.Itoa(int(Config.Redis.Port))

	if Config.APPEnv == "development" {
		Config.DBDriver = "sqlite3"
		Config.DBPath = path.Join(Config.SQLite.FullPath, Config.SQLite.FileName)
	} else {
		Config.DBDriver = "mysql"
		Config.DBPath = Config.MySQL.User + ":" + Config.MySQL.Password + "@tcp(" + Config.MySQL.Host + ":" + strconv.Itoa(int(Config.MySQL.Port)) + ")/" + Config.MySQL.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// fmt.Printf("full config: %#v", Config)
	data, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println("-------------配置信息------------")
	fmt.Printf("%s\n", data)
}

func init() {
	fullPaths := []string{Config.UploadFullPath, Config.LogFullPath, Config.SQLite.FullPath}
	for _, p := range fullPaths {
		if err := os.MkdirAll(p, 0755); err != nil {
			fmt.Println(err)
			panic("目录创建失败")
		}
	}
}
