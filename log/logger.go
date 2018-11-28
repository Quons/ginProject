package log

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"time"
	"github.com/pkg/errors"
)

func init() {
	writer, err := rotatelogs.New(
		"ginProject"+".%Y%m%d%H%M",
		rotatelogs.WithLinkName("/Users/apple/project/myGoProject/src/ginProject/ginProject/log.txt"), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(10*time.Hour*24),                                                        // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24),                                                     // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error.%+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
	logrus.AddHook(lfHook)
}
