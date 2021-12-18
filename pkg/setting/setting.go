package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Wechat struct {
	AppID     string
	AppSecret string
	Token     string

	GzhhAppId     string
	GzhhAppSecret string
	GzhhToken     string

	MerchantIndexSigned	int //用哪个商户号收报名费
	MerchantIndexSMS	int //用哪个商户号收短信费
	MerchantIndexTransfer	int //用哪个商户号提现

	MerchantIds    []string
	PayKeys        []string
	RootCas        []string

	NotifyURL      string
	SpbillCreateIP string
	DrawLimitOnce int //每笔限额，单位分
	DrawLimitDay int //每日限额，单位分
	DrawDefaultFeerate int //默认提现费率 ，千分之几

	TemplateIdEventStart     string
	TemplateIdSignedCheck    string
	TemplateIdSignedProgress string
}

var WechatSetting = &Wechat{}

type App struct {
	JwtUserSecret string
	JwtAdminSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	UploadPath     string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	//HttpsCertCrtPath string
	//HttpsCertKeyPath string
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
	Host          string
	Password      string
	MaxIdle       int
	MaxActive     int
	IdleTimeout   time.Duration
	DefaultQueue  string
	Broker        string
	ResultBackend string
}

var RedisSetting = &Redis{}

type Aliyun struct {
	AliyunAccessKeyId      string
	AliyunAccessKeySecret  string
	AliyunOssEndpoint      string
	AliyunOssDefaultBucket string
	AliyunOssDomain        string
}

var AliyunSetting = &Aliyun{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("wechat", WechatSetting)
	mapTo("aliyun", AliyunSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
