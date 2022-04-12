package vipers

import (
	"fmt"
	"github.com/spf13/viper"
)

//读取config配置

var (
	AppDebug         bool //是否是debug模式
	AllowCrossDomain bool //是否允许跨域
	HttpPort         string
	HttpIp           string
	JwtKey           string

	DbName     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DBCharset  string

	Filename      string
	MaxSize       int
	MaxBackups    int
	MaxAge        int
	Compress      bool
	TextLogFormat string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	viper.SetConfigName("config")  //配置文件名
	viper.SetConfigType("yaml")    // 配置文件类型，可以是yaml、json、xml。。。
	viper.AddConfigPath("config/") //配置文件路径
	err := viper.ReadInConfig()    //读取配置文件信息
	viper.WatchConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))

	}

	LoadDb()
	LoadRouter()
	LoadLog()
	//LoadQiniu()

}
func LoadDb() {
	DbHost = viper.GetString("Gorm.Host")
	DbName = viper.GetString("Gorm.DbName")
	DbPort = viper.GetString("Gorm.Port")
	DbPassWord = viper.GetString("Gorm.Password")
	DBCharset = viper.GetString("Gorm.Charset")
	DbUser = viper.GetString("Gorm.User")
}
func LoadRouter() {
	AppDebug = viper.GetBool("HttpServer.AppDebug")
	AllowCrossDomain = viper.GetBool("HttpServer.AllowCrossDomain")
	HttpPort = viper.GetString("HttpServer.HttpPort")
	HttpIp = viper.GetString("HttpServer.HttpIp")
	JwtKey = viper.GetString("Token.JwtTokenSignKey")
}
func LoadLog() {
	Filename = viper.GetString("Logs.myuLogFilename")
	MaxSize = viper.GetInt("Logs.MaxSize")
	MaxBackups = viper.GetInt("Logs.MaxBackups")
	MaxAge = viper.GetInt("Logs.MaxAge")
	Compress = viper.GetBool("Logs.Compress")
	TextLogFormat = viper.GetString("Logs.TextLogFormat")
}
func LoadQiniu() {
	AccessKey = viper.GetString("Qiniu.AccessKey")
	SecretKey = viper.GetString("Qiniu.SecretKey")
	Bucket = viper.GetString("Qiniu.Bucket")
	QiniuSever = viper.GetString("Qiniu.QiniuSever")
}
