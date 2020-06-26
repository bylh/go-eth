package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Trade struct {
	HuobiKey      string
	HuobiSecret   string
	BinanceKey    string
	BinanceSecret string
	OkexKey       string
	OkexSecret    string
}

var TradeSetting = &Trade{}
var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	fmt.Println("Setup start")
	var err error
	// 读取配置文件
	cfg, err = ini.Load("extra/trade/app.ini")
	// 读取失败处理
	fmt.Println("Setup start read file")
	if err != nil {
		// 此处如果读取失败，则在fatalf中直接退出程序即os.Exit(1)，不需要手动return;
		log.Fatalf("setting.Setup, fail to parse 'extra/trade/app.ini': %v", err)
	}

	fmt.Println("Setup read file success")

	// 文件中[app]字段映射到AppSetting, 其他类同
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("trade", TradeSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
