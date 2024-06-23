package DDBOT

import (
	_ "embed" // embed the default config file
	"fmt"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/rock8526652/DDBOT/lsp"
	"github.com/rock8526652/DDBOT/warn"
	"github.com/rock8526652/MiraiGo-Template/bot"
	"github.com/rock8526652/MiraiGo-Template/config"
	"github.com/sirupsen/logrus"

	_ "github.com/rock8526652/DDBOT/logging"
	_ "github.com/rock8526652/DDBOT/lsp/acfun"
	_ "github.com/rock8526652/DDBOT/lsp/douyu"
	_ "github.com/rock8526652/DDBOT/lsp/huya"
	_ "github.com/rock8526652/DDBOT/lsp/twitcasting"
	_ "github.com/rock8526652/DDBOT/lsp/weibo"
	_ "github.com/rock8526652/DDBOT/lsp/youtube"
	_ "github.com/rock8526652/DDBOT/msg-marker"
)

// defaultConfig 默认配置文件
//
//go:embed default_application.yaml
var defaultConfig string

// SetUpLog 使用默认的日志格式配置，会写入到logs文件夹内，日志会保留七天
func SetUpLog() {
	writer, err := rotatelogs.New(
		path.Join("logs", "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.WithError(err).Error("unable to write logs")
		return
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
	})
	logrus.AddHook(lfshook.NewHook(writer, &logrus.TextFormatter{
		FullTimestamp:    true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
		ForceQuote:       true,
	}))
}

// Run 启动bot，这个函数会阻塞直到收到退出信号
func Run() {
	if fi, err := os.Stat("device.json"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("警告：没有检测到device.json，正在生成，如果是第一次运行，可忽略")
			bot.GenRandomDevice()
		} else {
			warn.Warn(fmt.Sprintf("检查device.json文件失败 - %v", err))
			os.Exit(1)
		}
	} else {
		if fi.IsDir() {
			warn.Warn("检测到device.json，但目标是一个文件夹！请手动确认并删除该文件夹！")
			os.Exit(1)
		} else {
			fmt.Println("检测到device.json，使用存在的device.json")
		}
	}

	if fi, err := os.Stat("application.yaml"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("警告：没有检测到配置文件application.yaml，正在生成，如果是第一次运行，可忽略")
			sb := strings.Builder{}
			sb.WriteString(defaultConfig)
			if err := os.WriteFile("application.yaml", []byte(sb.String()), 0755); err != nil {
				warn.Warn(fmt.Sprintf("application.yaml生成失败 - %v", err))
				os.Exit(1)
			} else {
				fmt.Println("最小配置application.yaml已生成，请按需修改，如需高级配置请查看帮助文档")
			}
		} else {
			warn.Warn(fmt.Sprintf("检查application.yaml文件失败 - %v", err))
			os.Exit(1)
		}
	} else {
		if fi.IsDir() {
			warn.Warn("检测到application.yaml，但目标是一个文件夹！请手动确认并删除该文件夹！")
			os.Exit(1)
		} else {
			fmt.Println("检测到application.yaml，使用存在的application.yaml")
		}
	}

	config.GlobalConfig.SetConfigName("application")
	config.GlobalConfig.SetConfigType("yaml")
	config.GlobalConfig.AddConfigPath(".")
	config.GlobalConfig.AddConfigPath("./config")

	err := config.GlobalConfig.ReadInConfig()
	if err != nil {
		warn.Warn(fmt.Sprintf("读取配置文件失败！请检查配置文件格式是否正确 - %v", err))
		os.Exit(1)
	}
	config.GlobalConfig.WatchConfig()
	config.Base()

	// 快速初始化
	bot.Init()

	// 初始化 Modules
	bot.StartService()

	// 登录
	bot.Login()

	// 刷新好友列表，群列表
	bot.RefreshList()

	lsp.Instance.PostStart(bot.Instance)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	bot.Stop()
}
