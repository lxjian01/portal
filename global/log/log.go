package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"portal/global/config"
	"sync"
	"time"
)

const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
	levelFatal = "fatal"
)

var (
	o      = &sync.Once{}
	logger *logrus.Logger
	cfg    *config.LogConfig
)

type LogMetaSrv struct{}

func (*LogMetaSrv) Info(args ...interface{}) {
	Info(args...)
}

func (*LogMetaSrv) Warn(args ...interface{}) {
	Warn(args...)
}

func (*LogMetaSrv) Error(args ...interface{}) {
	Error(args...)
}

func (*LogMetaSrv) Debug(args ...interface{}) {
	Debug(args...)
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func isExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// CreateDir 创建文件夹
func createDir(dir string) (bool, error) {
	if err := os.MkdirAll(dir, os.FileMode(os.O_CREATE)); err != nil {
		return false, err
	}

	return true, nil
}

// CreateFile 创建文件
func createFile(fileFullName string) (bool, error) {
	parentDir := filepath.Dir(fileFullName)
	_, err := createDir(parentDir)
	if err != nil {
		return false, err
	}

	_, err = os.Create(fileFullName)
	if err != nil {
		return false, err
	}

	return true, nil
}

func Init(c *config.LogConfig) {
	o.Do(func() {
		cfg = c
		p := path.Join(cfg.Dir, cfg.Name)
		exist := isExist(p)
		if !exist {
			_, err := createFile(p)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("log path: ", p)
		timeFmt := new(logrus.TextFormatter)
		timeFmt.TimestampFormat = "2006-01-02 15:04:05.000"
		timeFmt.FullTimestamp = true

		if cfg.RetainDay < 0 || cfg.RetainDay > 30 {
			cfg.RetainDay = 7
		}

		writer, err := initWriter(p)
		if err != nil {
			panic(err)
		}

		logLevel, err := logrus.ParseLevel(cfg.Level)
		if err != nil {
			panic(err)
		}

		getLogger(writer, logLevel, timeFmt)
	})
}

func initWriter(file string) (*rotatelogs.RotateLogs, error) {
	writer, err := rotatelogs.New(
		//file+".%Y%m%d%H%M%S",
		file+".%Y-%m-%d",
		rotatelogs.WithLinkName(file),
		rotatelogs.WithMaxAge(time.Duration(cfg.RetainDay)*24*time.Hour),
		rotatelogs.WithRotationTime(86400),
	)

	return writer, err
}

func getLogger(writer *rotatelogs.RotateLogs, logLevel logrus.Level, customerFormatter *logrus.TextFormatter) {
	logger = logrus.New()
	logger.Out = os.Stdout
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		},
		customerFormatter,
	))

	logger.SetLevel(logLevel)
}

func SetLevel(level string) error {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logger.SetLevel(logLevel)
	return nil
}

func Debug(args ...interface{}) {
	log(levelDebug, args...)
}
func Debugf(format string, args ...interface{}) {
	logf(levelDebug, format, args...)
}

func Info(args ...interface{}) {
	log(levelInfo, args...)
}
func Infof(format string, args ...interface{}) {
	logf(levelInfo, format, args...)
}

func Warn(args ...interface{}) {
	log(levelWarn, args...)
}
func Warnf(format string, args ...interface{}) {
	logf(levelWarn, format, args...)
}

func Error(args ...interface{}) {
	log(levelError, args...)
}
func Errorf(format string, args ...interface{}) {
	logf(levelError, format, args...)
}

func Fatal(args ...interface{}) {
	log(levelFatal, args...)
}
func Fatalf(format string, args ...interface{}) {
	logf(levelFatal, format, args...)
}

func log(level string, args ...interface{}) {
	if logger == nil {
		fmt.Println(args...)
	} else {
		switch level {
		case levelDebug:
			logger.Debug(args...)
		case levelWarn:
			logger.Warn(args...)
		case levelError:
			logger.Error(args...)
		case levelFatal:
			logger.Fatal(args...)
		default:
			logger.Info(args...)
		}
	}
}
func logf(level string, format string, args ...interface{}) {
	if logger == nil {
		fmt.Printf(format, args...)
	} else {
		switch level {
		case levelDebug:
			logger.Debugf(format, args...)
		case levelWarn:
			logger.Warnf(format, args...)
		case levelError:
			logger.Errorf(format, args...)
		case levelFatal:
			logger.Fatalf(format, args...)
		default:
			logger.Infof(format, args...)
		}
	}
}
