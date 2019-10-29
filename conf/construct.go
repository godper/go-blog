package conf

import "time"

//App 配置
type App struct {
	JwtSecret string
	PageSize  int
	PrefixURL string

	RuntimeRootPath string

	ImagePrefixURL string
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

//Server 配置
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
