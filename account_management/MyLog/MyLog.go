package mylog

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

const (
	output_dir = "./logs/"
	out_path   = "app.MyLog"
	err_path   = "app.myerr"
)

func init() {
	_, err := os.Stat(output_dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(output_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("creating the dir failed![%v]\n", err)
			}
		}
	}

	//*set log foramt
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(d time.Duration, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendInt64(int64(d) / 1000000)
		},
	})

	//*判断两个日志等级
	infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return true
	})

	warnLeve := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.WarnLevel
	})

	//?获取info、warn日志文件的io.Writer抽象
	infoHook_1 := os.Stdout
	infoHook_2 := getWriter(out_path)
	errHook := getWriter(err_path)

	//^创建具体Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_1), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_2), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errHook), warnLeve),
	)

	//*传入zap.AddCaller() 打印日志文件名和行数
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	Log = logger.Sugar()
	defer logger.Sync()
}

func getWriter(filename string) io.Writer {
	//* 生成rotatelogs的Logger实际文件名 app.MyLog.YYmmddHH
	//* app.log指向最新日志文件链接
	//* 保存7天内日志，每1小时(整点)分隔一次
	hook, err := rotatelogs.New(
		output_dir+filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}

	return hook
}
