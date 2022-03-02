package globalInit

import (
	"cpg-blog/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var App = &app{}

type app struct {
	Name    string
	Version string
	Date    time.Time

	// 项目根目录
	RootDir string

	// 启动时间
	LaunchTime time.Time
	Uptime     time.Duration

	Year int

	Domain string
	Desc   map[string]string

	Build struct {
		GitCommitLog string
		BuildTime    string
		GitRelease   string
		GoVersion    string
		GinVersion   string
	}

	Env string

	locker sync.Mutex
}

func init() {
	App.Version = "V1.0"
	App.LaunchTime = time.Now()
	App.Year = time.Now().Year()

	// 默认在项目根目录运行程序
	App.RootDir = "."

	// 用来处理单元测试或当前目录不在根目录，获取项目根目录
	if !viper.InConfig("http.port") {
		App.RootDir = inferRootDir()
	}

	fileInfo, err := os.Stat(os.Args[0])
	if err != nil {
		panic(err)
	}

	App.Date = fileInfo.ModTime()
	App.Build.GoVersion = runtime.Version()
	App.Build.GinVersion = gin.Version
}

// inferRootDir 递归推导项目根目录
func inferRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		if d == "/" {
			panic("请确保在项目根目录或子目录下运行程序，当前在：" + cwd)
		}

		if util.Exist(d + "/config") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(cwd)
}

func (a *app) FillBuildInfo(gitCommitLog, buildTime, gitRelease string) {
	App.Name = viper.GetString("name")
	App.Domain = viper.GetString("domain") + viper.GetString("http.port")
	App.Desc = viper.GetStringMapString("desc")

	a.Build.GitCommitLog = gitCommitLog
	a.Build.BuildTime = buildTime

	pos := strings.Index(gitRelease, "/")
	if pos >= -1 {
		a.Build.GitRelease = gitRelease[pos+1:]
	}

	fmt.Println(*a)
}

// SetFrameMode debug\test\release
func (a *app) SetFrameMode(mode string) {
	gin.SetMode(mode)
	App.Env = gin.Mode()
}
